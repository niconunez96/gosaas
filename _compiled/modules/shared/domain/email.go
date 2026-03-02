package sharedDomain

import "context"

type TemplateName string

const TemplateNameWelcome TemplateName = "welcome"

var TemplatePaths = map[TemplateName]string{
	TemplateNameWelcome: "templates/welcome.html",
}

type Attachment struct {
	Content     []byte
	Filename    string
	ContentType string
}

type Recipient struct {
	Email        string
	Replacements map[string]any
	Attachments  []Attachment
}

type Email struct {
	From               string
	Subject            string
	TemplateName       TemplateName
	GeneralAttachments []Attachment
}

type EmailSender interface {
	SendEmail(ctx context.Context, email Email, recipient Recipient) error
	SendEmails(ctx context.Context, email Email, recipients ...Recipient) error
}
