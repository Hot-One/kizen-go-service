package email

import (
	"fmt"
	"net/smtp"

	"emperror.dev/errors"
	"github.com/Hot-One/kizen-go-service/config"
	"github.com/Hot-One/kizen-go-service/pkg/logger"
)

type EmailI interface {
	Send(to string, message string) error
}

type email struct {
	cfg *config.Config
	log logger.Logger
}

func NewEmail(cfg *config.Config, log logger.Logger) EmailI {
	return &email{
		cfg: cfg,
		log: log,
	}
}

func (e *email) Send(to string, message string) error {
	auth := smtp.PlainAuth("", e.cfg.EmailUser, e.cfg.EmailPassword, e.cfg.EmailHost)

	msg := "To: \"" + to + "\" <" + to + ">\n" +
		"From: \"" + e.cfg.EmailUser + "\" <" + e.cfg.EmailUser + ">\n" +
		"Subject: " + "Your verification code" + "\n" +
		message + "\n"

	fmt.Println(e.cfg.EmailHost, e.cfg.EmailPort, e.cfg.EmailUser, e.cfg.EmailPassword)

	if err := smtp.SendMail(fmt.Sprintf("%s:%d", e.cfg.EmailHost, e.cfg.EmailPort), auth, e.cfg.EmailUser, []string{to}, []byte(msg)); err != nil {
		return errors.Wrap(err, "error while sending message to email")
	}

	return nil
}
