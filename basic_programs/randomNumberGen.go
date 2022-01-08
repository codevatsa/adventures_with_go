// Random number generation
// The way math/rand (https://pkg.go.dev/math/rand) package generates random number is by using a seed which is fed to a generator
// and then gives random number as output .The problem is go has the same seed everytime code is run so the same sequence is generated when code is run
// For different numbers, seed with a different value, such as time.Now().UnixNano(), which yields a constantly-changing number.

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// create a unique seed by using current time to ensure seed is unique
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	rangeval := 10 // range upto 10
	for i := 0; i < 3; i++ {
		fmt.Println("Random number generated at", time.Now(), "====", r.Intn(rangeval))
	}
}
