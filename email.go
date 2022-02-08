package main

import (
	"fmt"
	"net/smtp"
)

type EmailConfig struct {
	SMTPHost     string
	FromEmail    string
	SMTPUsername string
	SMTPPassword string
}

func sendEmail(recipient, subject, body string, conf EmailConfig) error {
	msg := fmt.Sprintf(
		"From: %s\nTo: %s\nSubject: %s\n\n%s",
		conf.FromEmail,
		recipient,
		subject,
		body,
	)

	auth := smtp.CRAMMD5Auth(conf.SMTPUsername, conf.SMTPPassword)
	return smtp.SendMail(conf.SMTPHost, auth, conf.FromEmail, []string{recipient}, []byte(msg))
}
