package barber

import "fmt"

type Shop struct {
	ChairNumber int
	Barber      *Barber
}

type Barber struct{}

func (barber *Barber) CutHair() {
	fmt.Println("Cortando")
}
