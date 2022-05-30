package main

import "fmt"

func for_loop_example() {

	strings := []string{"hello", "world"}
	for i, s := range strings {
		fmt.Println(i, s)
	}

	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	for {
		fmt.Println("loop")
		break
	}
}
