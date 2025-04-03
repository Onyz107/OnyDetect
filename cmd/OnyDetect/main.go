package main

import (
	"fmt"
	"time"

	"github.com/Onyz107/OnyDetect/detectVM"
)

func main() {
	s := time.Now()
	defer func() {
		fmt.Printf("Finished execution in %.2f seconds.\n", time.Since(s).Seconds())
	}()

	for i, v := range detectVM.Run() {
		if v {
			fmt.Printf("VM or sandbox detected! Index: %d\n", i)
			return
		}
	}
	fmt.Println("No VM or sandbox detected!")
}
