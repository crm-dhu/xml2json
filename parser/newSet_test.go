package parser

import (
	"testing"
)

func TestAdd(t *testing.T) {
	set := NewSet("A", "B")
	if set.Size() != 2 {
		t.Errorf("Set size is not %d", set.Size())
	}
}
