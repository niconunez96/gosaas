package infra

import (
	sharedDomain "_compiled/modules/shared/domain"
	"_compiled/modules/shared/infra/email"
	"os"
)

type Container struct {
	Env string

	EmailSender sharedDomain.EmailSender
}

func (c *Container) GetEmailSender() sharedDomain.EmailSender {
	if c.EmailSender == nil {
		c.EmailSender = email.NewResendEmailSender()
	}
	return c.EmailSender
}

var DIContainer = &Container{Env: os.Getenv("ENV")}
