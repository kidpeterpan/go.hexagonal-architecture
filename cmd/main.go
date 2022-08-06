package main

import (
	"fmt"
	"log"

	"go.hexagonal-architecture/internal/adapters/core/arithmetic"
	"go.hexagonal-architecture/internal/ports"
)

func main() {
	var arith ports.Arithmetic
	arith = arithmetic.NewAdapter()
	res, err := arith.Additional(3, 2)
	if err != nil {
		log.Fatalf("Additional error: %v", err)
	}
	fmt.Println(res)
}
