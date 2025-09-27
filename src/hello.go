package main

import (
	"fmt"
	"os"
)

func main() {
	cmd := os.Args

	for _, arg := range cmd {
		fmt.Println(arg)
	}

}
