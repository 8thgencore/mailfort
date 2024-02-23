package http

import (
	"github.com/8thgencore/mailfort/internal/delivery/http/response"
	"github.com/8thgencore/mailfort/internal/service"
	"github.com/gin-gonic/gin"
)

// MailHandler handles HTTP requests related to sending emails.
type MailHandler struct {
	svc service.MailService
}

// NewMail Handler creates a new MailHandler instance with the passed MailService.
func NewMailHandler(svc service.MailService) *MailHandler {
	return &MailHandler{
		svc,
	}
}

// sendEmailWithOTPCodeRequest represents the request body for initiating authentication
type sendEmailWithOTPCodeRequest struct {
	Email   string `json:"email"    binding:"required" example:"user@example.com"`
	OTPCode string `json:"otp_code" binding:"required" example:"123456"`
}

// SendConfirmationEmail processes a request to generate and send an email with a confirmation code.
//
//	@Summary		Send Confirmation Email
//	@Description	Processes a request to generate and send an email with a confirmation code.
//	@Tags			Mail
//	@Accept			json
//	@Produce		json
//	@Param			request	body		sendEmailWithOTPCodeRequest	true	"Request body for sending confirmation email"
//	@Success		200		{object}	response.Response			"Success response"
//	@Failure		400		{object}	response.ErrorResponse		"Validation error"
//	@Failure		500		{object}	response.ErrorResponse		"Internal server error"
//	@Router			/email-confirmation [post]
func (h *MailHandler) SendConfirmationEmail(ctx *gin.Context) {
	var req sendEmailWithOTPCodeRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.ValidationError(ctx, err)
		return
	}

	err := h.svc.SendConfirmationEmail(ctx, req.Email, req.OTPCode)
	if err != nil {
		response.HandleError(ctx, err)
		return
	}

	response.HandleSuccess(ctx, nil)
}

// SendPasswordReset processes a request to generate and send an email with a password reset code.
//
//	@Summary		Send Password Reset Email
//	@Description	Processes a request to generate and send an email with a password reset code.
//	@Tags			Mail
//	@Accept			json
//	@Produce		json
//	@Param			request	body		sendEmailWithOTPCodeRequest	true	"Request body for sending password reset email"
//	@Success		200		{object}	response.Response			"Success response"
//	@Failure		400		{object}	response.ErrorResponse		"Validation error"
//	@Failure		500		{object}	response.ErrorResponse		"Internal server error"
//	@Router			/password-reset [post]
func (h *MailHandler) SendPasswordReset(ctx *gin.Context) {
	var req sendEmailWithOTPCodeRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.ValidationError(ctx, err)
		return
	}

	err := h.svc.SendPasswordResetEmail(ctx, req.Email, req.OTPCode)
	if err != nil {
		response.HandleError(ctx, err)
		return
	}

	response.HandleSuccess(ctx, nil)
}
