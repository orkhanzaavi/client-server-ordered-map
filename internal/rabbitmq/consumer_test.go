package rabbitmq_test

import (
	"github.com/rabbitmq/amqp091-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"testwork/internal/rabbitmq"
	"testwork/mocks"
)

func TestConsumer_Consume(t *testing.T) {
	t.Run(
		"should consume a message", func(t *testing.T) {
			queue := &amqp091.Queue{
				Name: "test",
			}
			amqpChannel := mocks.NewAmqpConsumerChannel(t)
			messages := make(chan amqp091.Delivery, 1)
			messages <- amqp091.Delivery{
				Body: []byte("test:arg1,arg2"),
			}
			rOnlyMsgs := func() <-chan amqp091.Delivery { return messages }()
			amqpChannel.On("Consume", queue.Name, "", true, false, false, false, amqp091.Table(nil)).
				Return(rOnlyMsgs, nil)
			consumer := rabbitmq.NewConsumer(amqpChannel, queue)

			cmdChannel, err := consumer.Consume()

			t.Log("Given one message in the queue")
			t.Log("When we consume the message")
			require.NoError(t, err)
			cmd := <-cmdChannel
			close(messages)
			t.Log("Then we should receive the message")
			assert.Equal(t, "test", cmd.Name)
			require.Len(t, cmd.Arguments, 2)
			assert.Equal(t, "arg1", cmd.Arguments[0])
			assert.Equal(t, "arg2", cmd.Arguments[1])
		},
	)
}
