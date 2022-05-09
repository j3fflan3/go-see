package main

import (
	"fmt"

	"github.com/j3fflan3/go-see/randy"
)

func main() {
	randy.Seed(5)
	fmt.Println(randy.Random())
}
