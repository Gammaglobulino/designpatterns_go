package builder_functional

import (
	"../builder_functional"
	"fmt"
	"testing"
)

func TestPersonBuilder_Build(t *testing.T) {
	b := builder_functional.PersonBuilder{}
	p := b.Called("Andrea").
		WorksAsA("dev").
		ContactHimUsing("+393440560238", "mazzantia@hotmail.com").
		Build()
	fmt.Println(*p)
}
