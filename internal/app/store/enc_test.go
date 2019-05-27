package store

import (
	"testing"
)

func TestBtoi(t *testing.T) {
	inp := []int{21, -1, 53}
	//oup := []string { "21", "-1", "53" }

	for idx, x := range inp {
		y := itob(x)
		z := btoi(y)
		t.Log(idx, y, x, z)
	}
}
