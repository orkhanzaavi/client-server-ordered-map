package logger

import (
	"context"
	"github.com/lmittmann/tint"
	"github.com/sethvargo/go-envconfig"
	"log/slog"
	"os"
)

type LoggerConfig struct {
	Level int `env:"LOG_LEVEL"`
}

func InitDefaultSlogLogger() {
	cfg := LoggerConfig{}
	err := envconfig.Process(context.Background(), &cfg)
	if err != nil {
		slog.Error("failed to process logger config", slog.String("err", err.Error()))
		return
	}
	w := os.Stdout

	slog.SetDefault(
		slog.New(
			tint.NewHandler(
				w, &tint.Options{
					Level: slog.Level(cfg.Level),
				},
			),
		),
	)
}
