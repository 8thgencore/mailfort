package mail

import "log/slog"


type MailService struct {
	log *slog.Logger
}

func NewMailService(log *slog.Logger) *MailService {
	return &MailService{
		log,
	}
}