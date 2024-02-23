package service

import "context"

type MailService interface {
	SendOtpCode(ctx context.Context, email, code string) error
}
