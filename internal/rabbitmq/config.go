package rabbitmq

import (
	"context"
	"fmt"
	"github.com/sethvargo/go-envconfig"
)

type AmqpConfig struct {
	Port     int    `env:"AMQP_PORT"`
	Host     string `env:"AMQP_HOST"`
	User     string `env:"AMQP_USER"`
	Pwd      string `env:"AMQP_PASSWORD"`
	Exchange string `env:"AMQP_EXCHANGE"`
}

func NewAmqpConfig() (*AmqpConfig, error) {
	cfg := AmqpConfig{}
	return &cfg, envconfig.Process(context.Background(), &cfg)
}

func Must(cfg *AmqpConfig, err error) *AmqpConfig {
	if err != nil {
		panic(err)
	}
	return cfg
}

func (c *AmqpConfig) GetConnectionString() string {
	return fmt.Sprintf("amqp://%s:%s@%s:%d/", c.User, c.Pwd, c.Host, c.Port)
}
