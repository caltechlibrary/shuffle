package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	// Caltech Library package
	"github.com/caltechlibrary/shuffle"
)

func main() {
	// Create our random number engine.
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	a := []int{1, 2, 3, 4, 5, 6, 7}
	shuffle.Ints(a, r)
	fmt.Printf("Shuffled ints: %+v\n", a)
	b := []string{"one", "two", "three", "four", "five", "six", "seven"}
	shuffle.Strings(b, r)
	fmt.Printf("Shuffled strings: %s\n", strings.Join(b, ", "))
}
