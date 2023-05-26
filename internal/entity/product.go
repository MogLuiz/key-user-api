package entity

import (
	"errors"
	"time"

	"github.com/MogLuiz/key-user-api/pkg/entity"
)

var (
	ErrIDIsRequired    = errors.New("id is required")
	ErrInvalidID       = errors.New("invalid id")
	ErrNameIsRequired  = errors.New("name is required")
	ErrPriceIsRequired = errors.New("price is required")
	ErrInvalidPrice    = errors.New("invalid price")
)

type Product struct {
	ID        entity.ID `json:"id"`
	Name      string    `json:"name"`
	Pricing   int       `json:"price"`
	CreatedAt time.Time `json:"created_at"`
}

func NewProduct(name string, pricing int) (*Product, error) {
	product := &Product{
		ID:        entity.NewID(),
		Name:      name,
		Pricing:   pricing,
		CreatedAt: time.Now(),
	}
	if err := product.Validate(); err != nil {
		return nil, err
	}
	return product, nil
}

func (p *Product) Validate() error {
	if p.ID.String() == "" {
		return ErrIDIsRequired
	}
	if _, err := entity.ParseID(p.ID.String()); err != nil {
		return ErrInvalidID
	}
	if p.Name == "" {
		return ErrNameIsRequired
	}
	if p.Pricing == 0 {
		return ErrPriceIsRequired
	}
	if p.Pricing < 0 {
		return ErrInvalidPrice
	}
	return nil
}
