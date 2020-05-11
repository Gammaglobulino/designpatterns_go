package builder

import (
	"../builder"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHtmlBuilder_AddChild(t *testing.T) {
	b := builder.NewHtmlBuilder("ul", "/ul")
	b.AddChild("li", "/li", "hello")
	b.AddChild("li", "/li", "world")
	b.AddChild("li", "/li", "I'm here")
	assert.EqualValues(t, "<ul>\n  <li>\n    hello\n  </li>\n  <li>\n    world\n  </li>\n  <li>\n    I'm here\n  </li>\n</ul>\n", b.String())
}

func TestHtmlBuilder_AddChildFluent(t *testing.T) {
	b := builder.NewHtmlBuilder("ul", "/ul")
	b.AddChildFluent("li", "/li", "hello").
		AddChildFluent("li", "/li", "world")
	assert.EqualValues(t, "<ul>\n  <li>\n    hello\n  </li>\n  <li>\n    world\n  </li>\n</ul>\n", b.String())

}
