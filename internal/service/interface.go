package service

import "context"

type MailService interface {
	SendConfirmationEmail(ctx context.Context, email, code string) error
	SendPasswordResetEmail(ctx context.Context, email, code string) error
}
