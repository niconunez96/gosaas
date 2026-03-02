package email

import (
	sharedDomain "_compiled/modules/shared/domain"
	"context"
	"os"

	"github.com/resend/resend-go/v2"
)

type ResendEmailSender struct {
	client *resend.Client
}

func NewResendEmailSender() *ResendEmailSender {
	client := resend.NewClient(os.Getenv("RESEND_API_KEY"))
	return &ResendEmailSender{
		client: client,
	}
}

func (s *ResendEmailSender) SendEmail(
	ctx context.Context,
	email sharedDomain.Email,
	recipient sharedDomain.Recipient,
) error {
	return nil
}

func (s *ResendEmailSender) SendEmails(
	ctx context.Context,
	email sharedDomain.Email,
	recipients ...sharedDomain.Recipient,
) error {
	return nil
}
