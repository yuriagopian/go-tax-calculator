package main

import "fmt"

type Order struct {
	ID       string
	Price    float64
	Quantity int
}

func (o Order) SetPrice(price float64) {
	o.Price = price
	fmt.Println("Price dentro do setPrice: ", o.Price)
}

func (o Order) getTotal() float64 {
	return o.Price * float64(o.Quantity)
}

func main() {
	name := "Yuri"
	name = "Yuri Agopian"
	fmt.Println("Hello,", name)

	order := Order{
		ID:       "123",
		Price:    10.0,
		Quantity: 5,
	}

	order.SetPrice(20.0)
	fmt.Println(order.getTotal())
}
