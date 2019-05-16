package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println(runtime.NumCPU() + 1)
	if 5 > 1 {
		fmt.Println("bigger")
	}
}
