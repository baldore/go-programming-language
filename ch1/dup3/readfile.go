// This reads a file by throwing everything into memory.
package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	filename := os.Args[1]
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	fmt.Print(string(data))
}
