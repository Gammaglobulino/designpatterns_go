package prototype

import (
	"errors"
	"fmt"
)

//create a shirt object cloner for a shirt factory

type ShirtCloner interface {
	GetClone(s int) (ItemInfoGetter, error)
}

const (
	White = iota
	Black
	Blue
)

func GetShirtsCloner() ShirtCloner {
	return &ShirtsCache{}
}

type ShirtsCache struct{}

func (s *ShirtsCache) GetClone(m int) (ItemInfoGetter, error) {
	switch m {
	case White:
		newItem := *WhiteProto
		return &newItem, nil
	default:
		return nil, errors.New("Shirt model not recognized")
	}

}

type ItemInfoGetter interface {
	GetInfo() string
}

type ShirtColor byte

type Shirt struct {
	Price float32
	SKU   string
	color ShirtColor
}

func (s *Shirt) GetInfo() string {
	return fmt.Sprintf("Price %0.2f , SKU %s , color %d", s.Price, s.SKU, s.color)
}

var WhiteProto *Shirt = &Shirt{
	Price: 15.00,
	SKU:   "empty",
	color: White,
}

func (i *Shirt) GetPrice() float32 {
	return i.Price
}
