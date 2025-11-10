package utils

import (
	"fmt"
)

func Loops() {
	/* for initialization;condition;incr/decr{
	   }
	*/
	for i := 0; i < 10; i++ {
		fmt.Println("hello")
	}
	// or modern range for loop 
	for i:=range(3){
		fmt.Println(i)
	}

	// for loop work as while loop
	i:=0
	for i<4 {
		i+=3
	}
	fmt.Println(i)
}
