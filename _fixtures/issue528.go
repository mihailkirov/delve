package main

import (
	"fmt"
	"github.com/mkirov/liner"
)

func main() {
	line := liner.NewLiner()
	line.Close()
	fmt.Printf("test\n")
}
