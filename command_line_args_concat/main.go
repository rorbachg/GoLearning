// Compares time between naive implementation of command line arguments
// concatenation and using "strings" package

package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	s, sep := "", ""

	start := time.Now()
	for i := 0; i < 10000; i++ {
		for _, str := range os.Args[1:] {
			s += sep + str
			sep = " "
		}
	}
	end := time.Since(start)
	fmt.Println(end)

	start = time.Now()
	for i := 0; i < 10000; i++ {
		strings.Join(os.Args[1:], " ")
	}
	end = time.Since(start)
	fmt.Println(end)

}
