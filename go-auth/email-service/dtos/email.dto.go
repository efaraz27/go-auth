package dtos

// EmailType is a type that defines the email type
type EmailType string

const (
	VerifyEmail    EmailType = "verify-email"
	ForgotPassword EmailType = "forgot-password"
)

// SendEmailDTO is a struct that defines the send email DTO
type SendEmailDTO struct {
	To    string    `json:"to" validate:"required,email"`
	Type  EmailType `json:"type" validate:"required"`
	Token string    `json:"token" validate:"required"`
}
