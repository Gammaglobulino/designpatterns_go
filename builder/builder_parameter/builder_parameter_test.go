package builder_parameter

import (
	"../builder_parameter"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSendEmail(t *testing.T) {
	msg := builder_parameter.SendEmail(func(b *builder_parameter.EmailBuilder) {
		b.From("foo@bar.com").
			To("bar@baz.com").
			Subject("Meeting").
			Body("Hello, do you want to meet?")
	})
	assert.EqualValues(t, "Email to foo@bar.com from bar@baz.com - sent\n", msg)
}
