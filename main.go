package main

import (
	"fmt"

	"github.com/j3fflan3/go-see/randy"
)

func main() {
	fmt.Println(randy.Hello())
	randy.Seed(5)
	fmt.Println(randy.Random())
}
