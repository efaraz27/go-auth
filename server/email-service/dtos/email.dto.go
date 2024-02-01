package dtos

// EmailType is a type that defines the email type
type EmailType string

const (
	VerifyEmail    string = "verify-email"
	ForgotPassword string = "forgot-password"
)

// SendEmailDTO is a struct that defines the send email DTO
type SendEmailDTO struct {
	To    string `json:"to" validate:"required,email"`
	Type  string `json:"type" validate:"required"`
	Token string `json:"token" validate:"required"`
}
