package barber

import (
	"fmt"
)

type Customer struct {
	Id uint
}

func (customer Customer) JoinBarberShop(shop *Shop) {
	select {
	case shop.Chairs <- customer.Id:
		fmt.Printf("[  Cliente %d  ] 👨 ⏱️ entrou na barbearia, esperando barbeiro..\n", customer.Id)
	default:
		fmt.Printf("[  Cliente %d  ] 👨 🚫 tentou entrar na babearia mas estava cheia, indo embora..\n", customer.Id)
		return
	}
}
