package rabbitmq

import (
	"context"
	"github.com/rabbitmq/amqp091-go"
	"testwork/internal/command"
	"time"
)

type AmqpProducerChannel interface {
	PublishWithContext(
		ctx context.Context,
		exchange, key string,
		mandatory, immediate bool,
		msg amqp091.Publishing,
	) error
}

type Producer struct {
	channel      AmqpProducerChannel
	exchangeName string
}

func NewProducerFromConfig(config AmqpConfig) (*Producer, error) {
	factory := newAmqpFactory(&config)
	channel, err := factory.CreateChannel()
	if err != nil {
		return nil, err
	}
	err = factory.CreateExchange(config.Exchange, channel)
	if err != nil {
		return nil, err
	}

	return NewProducer(channel, config.Exchange), nil
}

func NewProducer(
	channel AmqpProducerChannel,
	exchangeName string,
) *Producer {
	return &Producer{
		channel:      channel,
		exchangeName: exchangeName,
	}
}

func (p *Producer) Publish(
	ctx context.Context,
	cmd command.Command,
) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	return p.channel.PublishWithContext(
		ctx,
		p.exchangeName,
		"",
		false,
		false,
		amqp091.Publishing{
			ContentType: "text/plain",
			Body:        []byte(cmd.String()),
		},
	)
}
