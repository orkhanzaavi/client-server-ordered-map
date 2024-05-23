package rabbitmq_test

import (
	"context"
	"github.com/rabbitmq/amqp091-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"testwork/internal/command"
	"testwork/internal/rabbitmq"
	"testwork/mocks"
)

func TestProducer_Publish(t *testing.T) {
	t.Run(
		"should publish a message", func(t *testing.T) {
			channel := mocks.NewAmqpProducerChannel(t)

			channel.On(
				"PublishWithContext",
				mock.Anything,
				"test",
				"",
				false,
				false,
				mock.MatchedBy(
					func(msg amqp091.Publishing) bool {
						msg.Body = []byte("test:arg1,arg2")
						return true
					},
				),
			).Return(nil)
			producer := rabbitmq.NewProducer(channel, "test")

			err := producer.Publish(
				context.Background(),
				command.Command{
					Name:      "test",
					Arguments: []string{"arg1", "arg2"},
				},
			)

			t.Log("Given a producer")
			t.Log("When we publish a message")
			t.Log("Then the message should be published")
			assert.NoError(t, err)
			channel.AssertExpectations(t)
		},
	)
}
