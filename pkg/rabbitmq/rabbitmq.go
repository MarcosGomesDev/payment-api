package rabbitmq

import (
	"context"

	amqp "github.com/rabbitmq/amqp091-go"
)

func OpenChannel() (*amqp.Channel, error) {
	cnn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")

	if err != nil {
		panic(err)
	}

	ch, err := cnn.Channel()

	if err != nil {
		panic(err)
	}

	return ch, nil
}

func Consume(ch *amqp.Channel, out chan amqp.Delivery, queue string) error {
	msgs, err := ch.Consume(
		queue,
		"go-payment",
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return err
	}

	for msg := range msgs {
		out <- msg
	}

	return nil
}

func Publish(ctx context.Context, ch *amqp.Channel, body, exName string) error {
	err := ch.PublishWithContext(
		ctx,
		exName,
		"",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/json",
			Body:        []byte(body),
		},
	)

	if err != nil {
		return err
	}

	return nil
}