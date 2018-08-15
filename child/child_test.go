package child

import (
	"testing"
)

func TestIsLeaf(t *testing.T) {
	if IsLeaf() {
		t.Error("child is not a leaf")
	}
}
