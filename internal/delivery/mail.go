package handler

import (
	"context"

	"github.com/8thgencore/mailfort/internal/service"
	mailpb "github.com/8thgencore/mailfort/protos/gen/go/mail/v1"
)

// MailServiceServer implements the gRPC server interface.
type MailServiceServer struct {
	mailpb.UnimplementedMailServiceServer
	svc service.MailService
}

// NewMailServiceServer creates a new MailServiceServer instance with the passed MailService.
func NewMailServiceServer(svc service.MailService) *MailServiceServer {
	return &MailServiceServer{svc: svc}
}

// SendConfirmationEmailOTPCode handles sending a confirmation email.
func (s *MailServiceServer) SendConfirmationEmailOTPCode(ctx context.Context, req *mailpb.SendEmailWithOTPCodeRequest) (*mailpb.Response, error) {
	err := s.svc.SendConfirmationEmail(ctx, req.Email, req.OtpCode)
	if err != nil {
		return &mailpb.Response{Message: err.Error(), Success: false}, err
	}
	return &mailpb.Response{Message: "Success", Success: true}, nil
}

// SendPasswordResetOTPCode handles sending a password reset email.
func (s *MailServiceServer) SendPasswordResetOTPCode(ctx context.Context, req *mailpb.SendEmailWithOTPCodeRequest) (*mailpb.Response, error) {
	err := s.svc.SendPasswordReset(ctx, req.Email, req.OtpCode)
	if err != nil {
		return &mailpb.Response{Message: err.Error(), Success: false}, err
	}
	return &mailpb.Response{Message: "Success", Success: true}, nil
}
