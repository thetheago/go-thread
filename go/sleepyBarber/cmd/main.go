package main

import (
	"sleepybarber/cmd/barber"
	"sleepybarber/cmd/utils"
)

func main() {
	channel := make(barber.Chairs, 5)

	var barberMan barber.Barber = barber.Barber{Sleeping: false}
	var shop barber.Shop = barber.Shop{Chairs: channel}

	go startBarber(&shop, &barberMan)
	go startCustomers(&shop)

	for {
	}
}

func startBarber(shop *barber.Shop, barberMan *barber.Barber) {
	go barberMan.Start(shop)
}

func startCustomers(shop *barber.Shop) {
	var customerId uint = 0

	for {
		utils.SleepRandomMedium()

		var client barber.Customer = barber.Customer{Id: customerId}
		go client.JoinBarberShop(shop)

		customerId++
	}
}

/*

   2. Problema do barbeiro dorminhoco: em uma barbearia há um único barbeiro
   (representado por um thread).


   O barbeiro corta o cabelo de clientes que aguardam para ser atendidos (cada cliente é re-
   presentado por uma thread).


   Se não há nenhum cliente, o barbeiro dorme. Um cliente aguarda pelo
   barbeiro se houver uma cadeira disponível na sala de espera; caso contrário, o cliente imediatamente
   deixa a barbearia. Se houver uma cadeira disponível, o cliente a ocupa (há um número limitado
   de cadeiras na sala de espera). Se o barbeiro estiver dormindo, o cliente o acorda. Assuma que
   há n cadeiras na barbearia. Escreva uma classe BarbeiroDorminhoco para sincronizar as threads
   do problema utilizando semáforos. Inclua também em sua solução classes para as demais threads e
   uma classe Principal.
*/
