package mail

import (
	"context"
	"crypto/tls"
	"fmt"
	"log/slog"
	"time"

	"github.com/8thgencore/mailfort/internal/config"
	"github.com/8thgencore/mailfort/internal/domain"
	mail "github.com/xhit/go-simple-mail/v2"
)

type messageType string

var (
	registrationType  messageType = "registration"
	resetPasswordType messageType = "reset_password"
)

// SendConfirmationEmail реализует метод интерфейса MailService для отправки письма с кодом подтверждения.
func (s *MailService) SendConfirmationEmail(ctx context.Context, emailTo, code string) error {

	s.log.Info(fmt.Sprintf("Sending confirmation email to %s with code %s", emailTo, code))

	if err := sendMail(*s.log, *s.cfg, emailTo, code, registrationType); err != nil {
		return domain.ErrConflictingData
	}

	s.log.Info("Email Sent")

	return nil
}

// SendPasswordReset реализует метод интерфейса MailService для отправки письма с кодом подтверждения.
func (s *MailService) SendPasswordReset(ctx context.Context, emailTo, code string) error {

	s.log.Info(fmt.Sprintf("Sending confirmation email to %s with code %s", emailTo, code))

	if err := sendMail(*s.log, *s.cfg, emailTo, code, resetPasswordType); err != nil {
		return domain.ErrConflictingData
	}

	s.log.Info("Email Sent")

	return nil
}

func sendMail(log slog.Logger, cfg config.Mail, emailTo, code string, messageType messageType) error {
	server := mail.NewSMTPClient()

	// SMTP Server
	server.Host = cfg.Host
	server.Port = cfg.Port
	server.Username = cfg.Username
	server.Password = cfg.Password
	server.Encryption = mail.EncryptionSTARTTLS

	// Variable to keep alive connection
	server.KeepAlive = false

	// Timeout for connect to SMTP Server
	server.ConnectTimeout = 10 * time.Second

	// Timeout for send the data and wait respond
	server.SendTimeout = 10 * time.Second

	// Set TLSConfig to provide custom TLS configuration. For example,
	// to skip TLS verification (useful for testing):
	server.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// SMTP client
	smtpClient, err := server.Connect()

	if err != nil {
		log.Error("Error to connect SMTP server", "error", err.Error())
		return err
	}

	// New email simple html with inline and CC
	message := mail.NewMSG()
	message.SetFrom(cfg.Username).
		AddTo(emailTo).
		SetSubject("Confirm OTP code")

	switch messageType {
	case registrationType:
		message.SetBody(mail.TextHTML, fmt.Sprintf(htmlBodyForRegistration, code))
	case resetPasswordType:
		message.SetBody(mail.TextHTML, fmt.Sprintf(htmlBodyForPasswordReset, code))
	}

	// always check error after send
	if message.Error != nil {
		log.Error("Message have error", "error", message.Error)
		return err
	}

	// Call Send and pass the client
	err = message.Send(smtpClient)
	if err != nil {
		log.Error("Error sending message", "error", err)
		return err
	}

	return nil
}

const htmlBodyForRegistration = `<html>
	<head>
		<meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
		<title>Registration OTP Code</title>
	</head>
	<body>
		<h1>Registration OTP Code</h1>
		<p>Thank you for registering with us!</p>
		<p>Your OTP code for registration is: <b>%s</b>.</p>
		<p>Please use this code to complete your registration process.</p>
		<p>Have a great day!</p>
	</body>
</html>`

const htmlBodyForPasswordReset = `<html>
	<head>
		<meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
		<title>Password Reset OTP Code</title>
	</head>
	<body>
		<h1>Password Reset OTP Code</h1>
		<p>You have requested to reset your password.</p>
		<p>Your OTP code for password reset is: <b>%s</b>.</p>
		<p>Please use this code to reset your password.</p>
		<p>If you did not request this, please ignore this email.</p>
		<p>Thank you!</p>
	</body>
</html>`
