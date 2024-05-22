package main

import (
	"fmt"
	"github.com/mihailkirov/liner"
)

func main() {
	line := liner.NewLiner()
	line.Close()
	fmt.Printf("test\n")
}
