package zorro

import (
	"encoding/json"
	"fmt"
	"log/slog"

	amqp "github.com/rabbitmq/amqp091-go"
)

const defaultQueuePrefix = "zorrobpm.jobs."

type ConsumerOption func(*Consumer)

func WithQueuePrefix(prefix string) ConsumerOption {
	return func(c *Consumer) {
		c.queuePrefix = prefix
	}
}

type Consumer struct {
	ch          *amqp.Channel
	proxy       Proxy
	handler     Handler
	queuePrefix string
	queue       string
}

func NewConsumer(ch *amqp.Channel, proxy Proxy, h Handler, opts ...ConsumerOption) *Consumer {
	c := &Consumer{
		handler:     h,
		ch:          ch,
		proxy:       proxy,
		queuePrefix: defaultQueuePrefix,
		queue:       defaultQueuePrefix + h.JobName(),
	}
	for _, opt := range opts {
		opt(c)
	}
	return c
}

func (c *Consumer) Start() error {
	if err := c.ch.Qos(1, 0, false); err != nil {
		return fmt.Errorf("set qos: %w", err)
	}

	msgs, err := c.ch.Consume(c.queue, "", false, false, false, false, nil)
	if err != nil {
		return fmt.Errorf("consume queue %s: %w", c.queue, err)
	}

	slog.Info("consumer started", "queue", c.queue)

	go func() {
		for msg := range msgs {
			var input InputDto
			if err := json.Unmarshal(msg.Body, &input); err != nil {
				slog.Error("unmarshal message", "error", err, "queue", c.queue)
				_ = msg.Ack(false)
				continue
			}

			res, err := c.handler.Handle(input)
			if err != nil {
				if err := c.proxy.FailTask(input.ServiceTaskId, err.Error()); err != nil {
					slog.Error("fail task", "error", err)
				}
			} else {
				if err := c.proxy.CompleteTask(res); err != nil {
					slog.Error("complete task", "error", err)
				}
			}

			if err := msg.Ack(false); err != nil {
				slog.Error("ack message", "error", err)
			}
		}
	}()

	return nil
}
