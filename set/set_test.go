package set

import "testing"

func TestSet(t *testing.T) {
	a := New().Add(1).Add(2).AddMany(2, 3, 4, 5)
	b := New().AddMany(2, 3)
	a.Subtract(b)
	t.Logf("%v", "a")
}
