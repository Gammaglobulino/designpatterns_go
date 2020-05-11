package builder

import (
	"fmt"
	"strings"
)

const (
	IndentSize = 2
)

type HtmlElement struct {
	NameTagStart string //<p> or <ul>
	NameTagEnd   string
	Text         string // what is contained between tags
	Elements     []HtmlElement
}

func (e *HtmlElement) String() string {
	return e.string(0)
}
func (e *HtmlElement) string(indent int) string {
	sb := strings.Builder{}
	i := strings.Repeat(" ", IndentSize*indent)
	sb.WriteString(fmt.Sprintf("%s<%s>\n",
		i, e.NameTagStart))
	if len(e.Text) > 0 {
		sb.WriteString(strings.Repeat(" ",
			IndentSize*(indent+1)))
		sb.WriteString(e.Text)
		sb.WriteString("\n")
	}

	for _, el := range e.Elements {
		sb.WriteString(el.string(indent + 1))
	}
	sb.WriteString(fmt.Sprintf("%s<%s>\n",
		i, e.NameTagEnd))
	return sb.String()
}

func NewHtmlBuilder(nameTagStart, nameTagEnd string) *HtmlBuilder {
	return &HtmlBuilder{
		NameTagStart: nameTagStart,
		NameTagEnd:   nameTagEnd,
		Root:         HtmlElement{nameTagStart, nameTagEnd, "", []HtmlElement{}},
	}
}

type HtmlBuilder struct {
	NameTagStart string
	NameTagEnd   string
	Root         HtmlElement
}

func (b *HtmlBuilder) String() string {
	return b.Root.String()
}

func (b *HtmlBuilder) AddChild(nameTagStart string, nameTagEnd string, childText string) {
	e := HtmlElement{
		NameTagStart: nameTagStart,
		NameTagEnd:   nameTagEnd,
		Text:         childText,
		Elements:     []HtmlElement{},
	}
	b.Root.Elements = append(b.Root.Elements, e)
}

func (b *HtmlBuilder) AddChildFluent(nameTagStart string, nameTagEnd string, childText string) *HtmlBuilder {
	e := HtmlElement{
		NameTagStart: nameTagStart,
		NameTagEnd:   nameTagEnd,
		Text:         childText,
		Elements:     []HtmlElement{},
	}
	b.Root.Elements = append(b.Root.Elements, e)
	return b
}
