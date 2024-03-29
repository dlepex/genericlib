package slice

import "testing"

var ops Basic
var pred Pred

func TestSlice(t *testing.T) {
	var a, aa []E
	a = append(a, 1)
	a = append(a, 2)
	a = append(a, 3)
	ops.DeleteAt(a, 2)
	aa = append(aa, a...)
	ops.Reverse(aa)

	pred.FilterMut(&a, func(e E) bool {
		return e.(int) != 2
	})
	b, x := ops.Pop(a)
	b, ok := ops.Delete(b, 1)

	t.Logf("%v %v %v %v %v", a, b, x, ok, aa)

}
