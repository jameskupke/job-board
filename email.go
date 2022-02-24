package main

import (
	"fmt"
	"net/smtp"
	"strings"
)

func sendEmail(recipient, subject, body string, conf EmailConfig) error {
	msg := fmt.Sprintf(
		"From: devICT Job Board <%s>\nTo: %s\nSubject: %s\nContent-Type: text/html; charset=UTF-8\n\n%s",
		conf.FromEmail,
		recipient,
		subject,
		body,
	)

	host := strings.Split(conf.SMTPHost, ":")[0]
	auth := smtp.PlainAuth("", conf.SMTPUsername, conf.SMTPPassword, host)
	return smtp.SendMail(conf.SMTPHost, auth, conf.FromEmail, []string{recipient}, []byte(msg))
}
