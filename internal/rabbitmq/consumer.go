package rabbitmq

import (
	"github.com/rabbitmq/amqp091-go"
	"log/slog"
	"testwork/internal/command"
)

type AmqpConsumerChannel interface {
	Consume(
		queue, consumer string,
		autoAck, exclusive, noLocal, noWait bool,
		args amqp091.Table,
	) (<-chan amqp091.Delivery, error)
}

type Consumer struct {
	channel AmqpConsumerChannel
	queue   *amqp091.Queue
}

func NewConsumerFromConfig(config AmqpConfig) (*Consumer, error) {
	factory := newAmqpFactory(&config)
	channel, err := factory.CreateChannel()
	if err != nil {
		return nil, err
	}
	err = factory.CreateExchange(config.Exchange, channel)
	if err != nil {
		return nil, err
	}

	queue, err := factory.CreateQueueForExchange(config.Exchange, channel)

	if err != nil {
		return nil, err
	}

	return NewConsumer(channel, queue), nil
}

func NewConsumer(
	channel AmqpConsumerChannel,
	queue *amqp091.Queue,
) *Consumer {
	return &Consumer{
		channel: channel,
		queue:   queue,
	}
}

func (c *Consumer) Consume() (<-chan *command.Command, error) {
	msgsChannel, err := c.channel.Consume(
		c.queue.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return nil, err
	}

	msgs := make(chan *command.Command, 100)
	go func() {
		c.readMessages(msgsChannel, msgs)
	}()

	return msgs, nil
}

func (c *Consumer) readMessages(msgsChannel <-chan amqp091.Delivery, msgs chan *command.Command) {
	for msg := range msgsChannel {
		cmd, err := command.FromString(string(msg.Body))
		if err != nil {
			slog.Error(
				"failed to convert string to command",
				slog.String("err", err.Error()),
				slog.String("body", string(msg.Body)),
			)
			continue
		}
		msgs <- cmd
	}
}
