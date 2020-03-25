package observer

import (
	"../observer"
	"container/list"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewPerson(t *testing.T) {
	p := observer.NewPerson("Andrea")
	assert.EqualValues(t, "Andrea", p.Name)
}

func TestDoctorService_Notify(t *testing.T) {
	p := observer.NewPerson("Andrea")
	d := &observer.DoctorService{&list.List{}}
	p.Subscribe(d)
	p.CatchACold()
	assert.EqualValues(t, "Andrea", d.Msgs.Front().Value.(string))
}
