package utils

import (
	"fmt"
	"time"
	
)

func Display(s string) {
	for i := 0; i < 4; i++ {
		fmt.Println(s)
	}
	time.Sleep(500*time.Millisecond)
	go func(s string) {
		for i := 0; i < 5; i++ {
			fmt.Println(s)
		}
	}("this is anonymouse goroutine")
}
