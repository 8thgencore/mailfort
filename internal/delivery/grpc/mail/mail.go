package mail

import (
	"context"

	mailv1 "github.com/8thgencore/mailfort/gen/go/mail"
	"github.com/8thgencore/mailfort/internal/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// MailHandler представляет интерфейс для обработки операций, связанных с электронной почтой.
type MailHandler interface {
	SendConfirmationEmail(ctx context.Context, email, otpCode string) (bool, error)
	SendPasswordReset(ctx context.Context, email, otpCode string) (bool, error)
}

// MailHandlerImpl реализует интерфейс MailHandler для обработки операций с электронной почтой.
type MailHandlerImpl struct {
	mailv1.UnimplementedMailServer
	svc service.MailService
}

func Register(gRPC *grpc.Server, svc service.MailService) {
	mailv1.RegisterMailServer(gRPC, &MailHandlerImpl{svc: svc})
}

// NewMailHandler создает новый экземпляр MailHandlerImpl.
func NewMailHandler() *MailHandlerImpl {
	return &MailHandlerImpl{}
}

// SendConfirmationEmail реализует метод интерфейса MailHandler для отправки письма с кодом подтверждения.
func (h *MailHandlerImpl) SendConfirmationEmail(ctx context.Context, req *mailv1.SendRequest) (*mailv1.SendResponse, error) {
	if err := validateRequest(req); err != nil {
		return nil, err
	}

	err := h.svc.SendConfirmationEmail(ctx, req.Email, req.OtpCode)
	if err != nil {
		return &mailv1.SendResponse{
			IsSuccess: false,
		}, nil
	}

	return &mailv1.SendResponse{
		IsSuccess: true,
	}, nil
}

// SendPasswordReset реализует метод интерфейса MailHandler для отправки письма с кодом для сброса пароля.
func (h *MailHandlerImpl) SendPasswordReset(ctx context.Context, req *mailv1.SendRequest) (*mailv1.SendResponse, error) {
	if err := validateRequest(req); err != nil {
		return nil, err
	}

	err := h.svc.SendPasswordReset(ctx, req.Email, req.OtpCode)
	if err != nil {
		return &mailv1.SendResponse{
			IsSuccess: false,
		}, nil
	}

	return &mailv1.SendResponse{
		IsSuccess: true,
	}, nil
}

func validateRequest(req *mailv1.SendRequest) error {
	if req.GetEmail() == "" {
		return status.Error(codes.InvalidArgument, "email is required")
	}

	if req.GetOtpCode() == "" {
		return status.Error(codes.InvalidArgument, "otp_code is required")
	}

	return nil
}
