package randy

/*
#include <stdlib.h>
*/
import "C"

func Random() int {
	return int(C.random())
}

func Seed(i int) {
	C.srandom(C.uint(i))
}

func Hello() string {
	return "Hello, I'm Randy, a Go program that imports \"C\"!"
}
