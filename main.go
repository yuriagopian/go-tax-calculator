package main

import "fmt"

type Order struct {
	ID       string
	Price    float64
	Quantity int
}

func (o Order) getTotal() float64 {
	return o.Price * float64(o.Quantity)
}

func main() {
	name := "Yuri"
	name = "Yuri Agopian"
	print("Hello, world! ", name)
	fmt.Println(name)
}
