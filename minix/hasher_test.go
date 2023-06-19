package minix

import (
	"fmt"
	"testing"
)

func TestHasher(t *testing.T) {
	p1, p2 := Hasher("region", "b3c659a887f00950")
	fmt.Println(p1, p2)

}
