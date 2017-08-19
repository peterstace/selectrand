package main

import (
	"fmt"
	"math/rand"

	"github.com/peterstace/selectrand"
)

func main() {
	rnd := rand.New(selectrand.SelectSource{})
	for i := 0; i < 10; i++ {
		fmt.Printf("%063b\n", rnd.Int63())
	}
}
