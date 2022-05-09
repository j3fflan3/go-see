package randy

import (
	"fmt"
	"testing"
)

func Test_Random(t *testing.T) {
	Seed(5)
	fmt.Println(Random())
}
