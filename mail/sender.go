package mail

import (
	"fmt"
	"net/smtp"

	"github.com/jordan-wright/email"
)

const (
	host          = "smtp.gmail.com"
	serverAddress = "smtp.gmail.com:587"
)

type EmailSender interface {
	SendEmail(
		subject string,
		content string,
		to []string,
		cc []string,
		bcc []string,
		attachFiles []string,
	) error
}

type GmailSender struct {
	name              string
	fromEmailAdress   string
	fromEmailPassword string
}

func NewGmailSender(name, fromEmailAdderss, fromEmailPassword string) EmailSender {
	return &GmailSender{
		name:              name,
		fromEmailAdress:   fromEmailAdderss,
		fromEmailPassword: fromEmailPassword,
	}
}

func (sender *GmailSender) SendEmail(subject string, content string, to []string, cc []string, bcc []string, attachFiles []string) error {
	e := email.NewEmail()
	e.From = fmt.Sprintf("%s <%s>", sender.name, sender.fromEmailAdress)
	e.Subject = subject
	e.HTML = []byte(content)
	e.To = to
	e.Cc = cc
	e.Bcc = bcc

	for _, f := range attachFiles {
		_, err := e.AttachFile(f)
		if err != nil {
			return fmt.Errorf("failed to attach file: %w", err)
		}
	}

	auth := smtp.PlainAuth("", sender.fromEmailAdress, sender.fromEmailPassword, host)
	return e.Send(serverAddress, auth)
}
