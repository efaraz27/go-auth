package services

import (
	"gopkg.in/gomail.v2"
)

// Email template paths
const VERIFY_EMAIL_TEMPLATE_PATH = "templates/verify-email.html"

// EmailService is a struct that defines the email service
type EmailService struct {
	Dialer *gomail.Dialer
}

// NewEmailService is a function that returns a new email service
func NewEmailService(host string, port int, user string, password string) *EmailService {
	dialer := gomail.NewDialer(host, port, user, password)

	return &EmailService{
		Dialer: dialer,
	}
}

// SendVerificationEmail is a function that sends a verification email
func (s *EmailService) SendVerificationEmail(from string, to string, token string) error {
	m := gomail.NewMessage()

	m.SetHeader("From", from)
	m.SetHeader("To", to)
	m.SetHeader("Subject", "Verify your email")
	m.SetBody("text/html", VERIFY_EMAIL_TEMPLATE_PATH)

	return s.Dialer.DialAndSend(m)
}
