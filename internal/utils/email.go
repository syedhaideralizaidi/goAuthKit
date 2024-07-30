package utils

import (
	"fmt"
	"net/smtp"
	"os"
)

var (
	smtpHost     = "smtp.gmail.com"
	smtpPort     = "587"
	smtpUser     = os.Getenv("SMTP_USER")
	smtpPassword = os.Getenv("SMTP_PASSWORD")
)

func sendEmail(to, subject, body string) error {
	from := smtpUser

	// Create the email message
	message := []byte(fmt.Sprintf("From: %s\r\nTo: %s\r\nSubject: %s\r\nContent-Type: text/html; charset=\"UTF-8\"\r\n\r\n%s", from, to, subject, body))

	// Connect to the SMTP server
	auth := smtp.PlainAuth("", smtpUser, smtpPassword, smtpHost)
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, message)
	return err
}

func SendVerificationEmail(email, token string) error {
	subject := "Email Verification"
	body := fmt.Sprintf("<p>Please verify your email using this token: <strong>%s</strong></p>", token)
	return sendEmail(email, subject, body)
}

func SendResetPasswordEmail(email, token string) error {
	subject := "Reset Password Email"
	body := fmt.Sprintf("Please reset your password by clicking the following link: http://localhost:8080/reset-password?token=%s", token)
	return sendEmail(email, subject, body)
}
