package main

import (
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
	slog.Info("Starting the test commands server...")
	config.LoadDefaultEnv()

	cfg := rabbitmq.Must(rabbitmq.NewAmqpConfig())
	consumer, err := rabbitmq.NewConsumerFromConfig(*cfg)
	if err != nil {
		slog.Error("failed to create a consumer", slog.String("err", err.Error()))
		return
	}

	msgs, err := consumer.Consume()
	if err != nil {
		slog.Error("failed to consume messages", slog.String("err", err.Error()))
		return
	}

	processor := command.NewProcessorWithDefaultStorage()

	go func() {
		for msg := range msgs {
			slog.Debug("received a message", slog.String("msg", fmt.Sprintf("%v", msg)))
			err := processor.Process(*msg)
			if err != nil {
				slog.Error("failed to process the message", slog.String("err", err.Error()))
			} else {
				slog.Debug("message processed successfully", slog.String("msg", fmt.Sprintf("%v", msg)))
			}
		}
	}()

	slog.Info(" [*] Waiting for commands from clients. To exit press CTRL+C")

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	<-c
	// Run Cleanup
	slog.Info("Shutting down the test commands server...")
	os.Exit(0)

}
