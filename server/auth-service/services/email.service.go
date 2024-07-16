package services

import (
	"github.com/efaraz27/go-auth/server/auth-service/core"
	"github.com/efaraz27/go-auth/server/auth-service/dtos/protobufs"
	"github.com/gofiber/fiber/v2"
	amqp "github.com/rabbitmq/amqp091-go"
	"google.golang.org/protobuf/proto"
)

// EmailService is a struct that defines the email service
type EmailService struct {
	rabbitmq  *core.RabbitMQ
	queueName string
}

// NewEmailService is a function that returns a new email service
func NewEmailService(rabbitmq *core.RabbitMQ, queueName string) *EmailService {
	return &EmailService{rabbitmq, queueName}
}

// SendVerificationEmailRequest is a method that publishes an email event to the queue
func (s *EmailService) SendVerificationEmailRequest(c *fiber.Ctx, to string, token string) error {

	payload := &protobufs.SendVerificationEmailRequest{
		To:    to,
		Token: token,
	}

	data, err := proto.Marshal(payload)
	if err != nil {
		return err
	}

	err = s.rabbitmq.Ch.PublishWithContext(c.Context(), "", s.queueName, false, false, amqp.Publishing{
		ContentType: "application/protobuf",
		Body:        data,
	})

	return nil
}
