package service

import (
	"github.com/Vladislav557/auth/internal/models/entity"
	"go.uber.org/zap"
	"net/smtp"
	"os"
)

const (
	from = "auth@gmail.com"
)

type Notifier struct{}

func (n *Notifier) AcceptRegistration(u *entity.User) error {
	to := []string{u.Email}
	link := "http://172.1.10.1:8080/confirm-email?user=" + u.UUID
	msg := []byte("To: " + u.Email + "\r\n" +
		"Subject: New registration\r\n\r\n" +
		"For complete registration click on link " + link + "\r\n")

	err := smtp.SendMail(os.Getenv("SMTP"), nil, from, to, msg)
	if err != nil {
		zap.L().Error("Error sending email", zap.Error(err))
		return nil
	}

	zap.L().Info("Email sent successfully!")
	return nil
}
