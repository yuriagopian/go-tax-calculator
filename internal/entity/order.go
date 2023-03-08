package entity

import "errors"

type Order struct {
	ID         string
	Price      float64
	Tax        float64
	FinalPrice float64
}

func (o *Order) Validate() error {
	if o.ID == "" {
		return errors.New("id is required")
	}

	if o.Price <= 0 {
		return errors.New("invalid Price value")
	}

	if o.Tax <= 0 {
		return errors.New("invalid Tax value")
	}

	return nil // erro nulo
}
