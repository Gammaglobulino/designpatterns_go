package builder_parameter

import (
	"../builder_parameter"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSendEmail(t *testing.T) {
	ret := builder_parameter.SendEmail(func(b *builder_parameter.EmailBuilder) {
		b.From("andrea@and.com").
			To("stefando@stef.com").
			Body("test email").
			Subject("test email")
	})
	assert.EqualValues(t, "andrea@and.com-stefando@stef.com-test email-test email", ret)
}
