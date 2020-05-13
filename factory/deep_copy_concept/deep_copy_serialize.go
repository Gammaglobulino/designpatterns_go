package deep_copy_concept

import (
	"bytes"
	"encoding/gob"
)

type Address struct {
	StreetName   string
	StreetNumber int
}
type Person struct {
	Name   string
	Street *Address
}

func (p *Person) DeepCopy() *Person {
	copiedBuffer := &bytes.Buffer{}
	newEncoder := gob.NewEncoder(copiedBuffer)
	newEncoder.Encode(p)
	copiedPerson := &Person{}
	newDecoder := gob.NewDecoder(copiedBuffer)
	_ = newDecoder.Decode(copiedPerson)
	return copiedPerson
}
