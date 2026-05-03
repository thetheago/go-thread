package main

import (
	"fmt"
	"threadABC/cmd/threads"
	"time"
)

func main() {
	channel := make(chan struct{}, 100)

	threadA := threads.WriteThread{Channel: channel}
	threadB := threads.WriteThread{Channel: channel}
	threadC := threads.ReadThread{Channel: channel}

	go threadA.Run("A", 10)
	go threadB.Run("B", 10)
	go threadC.Run("C")

	time.Sleep(time.Second * 10)

	fmt.Println("")
	fmt.Println("Count A: ", threadA.Count)
	fmt.Println("Count B: ", threadB.Count)
	fmt.Println("Count C: ", threadC.Count)
}
