package service

import "context"

type MailService interface {
	SendConfirmationEmail(ctx context.Context, emailTo, code string) error
	SendPasswordReset(ctx context.Context, emailTo, code string) error
}
