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

	for _, v := range detectVM.Run() {
		if v {
			fmt.Println("VM or sandbox detected!")
			return
		}
	}
	fmt.Println("No VM or sandbox detected!")
}
