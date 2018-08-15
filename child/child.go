package child // import "github.com/keks/testmod/child"

import (
	"github.com/keks/testmod/leaf"
)

var IsChild = true

func IsLeaf() bool {
	return !leaf.IsLeaf()
}
