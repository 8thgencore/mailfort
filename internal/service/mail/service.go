package mail

import (
	"log/slog"

	"github.com/8thgencore/mailfort/internal/config"
)

type MailService struct {
	log *slog.Logger
	cfg *config.Mail
}

func NewMailService(log *slog.Logger, cfg *config.Mail) *MailService {
	return &MailService{
		log,
		cfg,
	}
}
