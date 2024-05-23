package rabbitmq

import (
	"github.com/rabbitmq/amqp091-go"
)

type amqpFactory struct {
	config *AmqpConfig
}

func newAmqpFactory(config *AmqpConfig) *amqpFactory {
	return &amqpFactory{
		config: config,
	}
}

func (f *amqpFactory) CreateChannel() (*amqp091.Channel, error) {
	conn, err := amqp091.Dial(f.config.GetConnectionString())
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	return ch, nil
}

func (f *amqpFactory) CreateExchange(name string, ch *amqp091.Channel) error {
	err := ch.ExchangeDeclare(
		name,
		"fanout",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	return nil
}

func (f *amqpFactory) CreateQueueForExchange(exchangeName string, ch *amqp091.Channel) (*amqp091.Queue, error) {
	q, err := ch.QueueDeclare(
		"",
		false,
		false,
		true,
		false,
		nil,
	)
	if err != nil {
		return &q, nil
	}

	err = ch.QueueBind(
		q.Name,
		"#",
		exchangeName,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}

	return &q, nil
}
