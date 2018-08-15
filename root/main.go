package main // import "github.com/keks/testmod/root"

import (
	"fmt"
	
	"github.com/keks/testmod/child"
)

func main() {
	fmt.Println("child is child:", child.IsChild)
	fmt.Println("child is leaf:", child.IsLeaf)
}
