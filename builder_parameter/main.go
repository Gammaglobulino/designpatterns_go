package builder_parameter

import (
	"fmt"
	"strings"
)

type email struct {
	subject, from, to, body string
}

type EmailBuilder struct {
	email email
}

func (b *EmailBuilder) Subject(text string) *EmailBuilder {
	b.email.subject = text
	return b
}

func (b *EmailBuilder) From(text string) *EmailBuilder {
	if !strings.Contains(text, "@") {
		panic("not a vaild email")
	}
	b.email.from = text
	return b
}

func (b *EmailBuilder) To(text string) *EmailBuilder {
	if !strings.Contains(text, "@") {
		panic("not a vaild email")
	}
	b.email.to = text
	return b
}

func (e *EmailBuilder) Body(text string) *EmailBuilder {
	e.email.body = text
	return e
}

func sendEmailImpl(email *email) string {
	return fmt.Sprintf("%s-%s-%s-%s", email.from, email.to, email.subject, email.body)
}

type build func(*EmailBuilder)

func SendEmail(action build) string {
	builder := EmailBuilder{}
	action(&builder)
	return sendEmailImpl(&builder.email)
}
