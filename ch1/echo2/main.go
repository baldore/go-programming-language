// Echo2 prints its command-line arguments
package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func printArgs1() {
	var s, sep string
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

func printArgs2() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}

func main() {
	start := time.Now()

	// printArgs1()
	printArgs2()

	fmt.Println(time.Since(start))
}
