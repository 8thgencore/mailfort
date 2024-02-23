package mail

import (
	"context"
	"crypto/tls"
	"time"

	mail "github.com/xhit/go-simple-mail/v2"
)

const htmlBody = `<html>
	<head>
		<meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
		<title>Hello Gophers!</title>
	</head>
	<body>
		<p>This is the <b>Go gopher</b>.</p>
		<p><img src="cid:Gopher.png" alt="Go gopher" /></p>
		<p>Image created by Renee French</p>
	</body>
</html>`

// SendOtpCode реализует метод интерфейса MailService для отправки письма с кодом подтверждения.
func (s *MailService) SendOtpCode(ctx context.Context, email, code string) error {

	s.log.Info("Sending confirmation email to %s with code %s", email, code)

	server := mail.NewSMTPClient()

	// SMTP Server
	server.Host = s.cfg.Host
	server.Port = s.cfg.Port
	server.Username = s.cfg.Username
	server.Password = s.cfg.Password
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
		s.log.Error("Error to connect SMTP server", "error", err.Error())
		return err
	}

	// New email simple html with inline and CC
	message := mail.NewMSG()
	message.SetFrom("From Example <nube@example.com>").
		AddTo(s.cfg.Username).
		AddCc(email).
		SetSubject("New Go message").
		SetListUnsubscribe("<mailto:unsubscribe@example.com?subject=https://example.com/unsubscribe>")

	message.SetBody(mail.TextHTML, htmlBody)

	// always check error after send
	if message.Error != nil {
		s.log.Error("Message have error", "error", message.Error)
		return err

	}

	// Call Send and pass the client
	err = message.Send(smtpClient)
	if err != nil {
		s.log.Error("Error sending message", "error", err)
		return err
	}

	s.log.Info("Email Sent")

	return nil
}
