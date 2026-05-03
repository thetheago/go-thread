package threads

import (
	"fmt"
	"time"
)

type WriteThread struct {
	Channel chan struct{}
	Count   uint
}

type ReadThread struct {
	Channel chan struct{}
	Count   uint
}

func (this *WriteThread) Run(letter string, repeatCount int) {
	for i := 0; i < repeatCount; i++ {
		fmt.Print(letter)
		this.Channel <- struct{}{}
		this.Count++
		time.Sleep(time.Millisecond * 10)
	}
}

func (this *ReadThread) Run(letter string) {
	for {
		<-this.Channel
		fmt.Print(letter)
		this.Count++
		time.Sleep(time.Millisecond * 10)
	}
}
