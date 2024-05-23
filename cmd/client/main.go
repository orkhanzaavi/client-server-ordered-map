package main

import (
	"context"
	"flag"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"testwork/internal/command"
	"testwork/internal/config"
	"testwork/internal/logger"
	"testwork/internal/rabbitmq"
)

func init() {
	config.LoadDefaultEnv()
	logger.InitDefaultSlogLogger()
}

func main() {
	fileName := flag.String("file", "", "File name with commands. It is the mandatory flag")
	flag.Parse()
	if *fileName == "" {
		slog.Error("File name parameter '--file=name.txt' is mandatory")
		os.Exit(1)
	}

	slog.Debug("Starting the client...")
	slog.Info("Reading the file with commands...")

	cfg := rabbitmq.Must(rabbitmq.NewAmqpConfig())
	producer, err := rabbitmq.NewProducerFromConfig(*cfg)
	if err != nil {
		slog.Error("failed to create a producer", slog.String("err", err.Error()))
		return
	}

	commandChan, err := command.ReadCommandsFile(*fileName)
	if err != nil {
		slog.Error("failed to read commands from the file", slog.String("err", err.Error()))
		return
	}
	ctx, cancel := context.WithCancel(context.Background())
	endSuccess := make(chan bool)
	endFail := make(chan bool)
	go func() {
		i := 0
		slog.Info(" [*] Sending commands to the server. Press CTRL+C to exit")
		for cmd := range commandChan {
			if cmd == nil {
				continue
			}
			err = producer.Publish(ctx, *cmd)
			if err != nil {
				slog.Error("failed to publish a command", slog.String("err", err.Error()))
				endFail <- true
				return
			}
			i++
			if (i < 10) ||
				(i < 100 && i%10 == 0) ||
				(i >= 100 && i < 1000 && i%100 == 0) ||
				(i >= 1000 && i%1000 == 0) {
				slog.Info(fmt.Sprintf(" [*] Sent %d commands", i))
			}
		}
		endSuccess <- true
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	select {
	case <-endSuccess:
		slog.Info("All commands are sent")
	case <-endFail:
		slog.Error("Script execution is failed")
		os.Exit(1)
	case <-c:
		slog.Info("Shutting down the client...")
		cancel()
	}
}
