package main

import (
	"fmt"
	"math/rand"

	"github.com/peterstace/selectrand"
)

func main() {
	rnd := rand.New(selectrand.Source{})

	fmt.Println("\nInt63:")
	for i := 0; i < 10; i++ {
		fmt.Printf("%064b\n", rnd.Int63())
	}
	fmt.Println("\nUint64:")
	for i := 0; i < 10; i++ {
		fmt.Printf("%064b\n", rnd.Uint64())
	}
	fmt.Println("\nFloat64:")
	for i := 0; i < 10; i++ {
		fmt.Printf("%.70f\n", rnd.Float64())
	}
}
