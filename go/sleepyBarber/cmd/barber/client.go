package barber

import (
	"fmt"
)

type Client struct {
	Id uint
}

func (client Client) JoinBarberShop(shop *Shop) {
	fmt.Println("Cliente entrando na barbearia..")

	if shop.ChairNumber == 0 {
		fmt.Println("Barbearia está cheia, cliente saíndo..")
		return
	}

	shop.ChairNumber--
	fmt.Println("Cliente entrou na barbearia e ocupou um acento..")
	fmt.Println(shop.ChairNumber)
}
