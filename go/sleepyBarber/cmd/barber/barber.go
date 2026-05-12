package barber

import (
	"fmt"
	"sleepybarber/cmd/utils"
)

type Barber struct {
	Sleeping bool
}

func (barber Barber) Start(shop *Shop) {
	fmt.Println("[  Barbeiro  ] 🥸    está pronto para cortar cabelo!")

	for {
		select {
		case customerId := <-shop.Chairs:

			if barber.Sleeping {
				fmt.Println("[  Barbeiro  ] 🥸 ⚠️  ACORDOU!!")
				barber.Sleeping = false
			}

			fmt.Printf("[  Barbeiro  ] 🥸 ✂️  cortando cabelo do cliente %d..\n", customerId)

			utils.SleepBarberCuttingHair()

			fmt.Printf("[  Barbeiro  ] 🥸 ✂️ ✅   terminou de cortar o cabelo do %d..\n", customerId)
		default:
			if !barber.Sleeping {
				fmt.Println("[  Barbeiro  ] 🥸 💤  nenhum cliente na barbaria, tirando uma soneca....")
				barber.Sleeping = true
			}
		}

	}
}
