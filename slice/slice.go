package slice

//E - element type
type E = interface{}

//Ops - slice operations type
type Ops struct{}

func (Ops) Pop(a []E) ([]E, E) {
	l := len(a)
	return a[:l-1], a[l-1]
}

func (Ops) IndexOf(a []E, elem E) int {
	for i, e := range a {
		if e == elem {
			return i
		}
	}
	return -1
}

func (ops Ops) Contains(a []E, elem E) bool {
	return ops.IndexOf(a, elem) >= 0
}

func (ops Ops) Delete(a []E, elem E) ([]E, bool) {
	idx := ops.IndexOf(a, elem)
	if idx < 0 {
		return a, true
	}
	return ops.DeleteIndex(a, idx), false
}

func (ops Ops) DeleteIndex(a []E, idx int) []E {
	return append(a[0:idx], a[idx+1:]...)
}

func (Ops) Reverse(a []E) {
	for i := len(a)/2 - 1; i >= 0; i-- {
		opp := len(a) - 1 - i
		a[i], a[opp] = a[opp], a[i]
	}
}

func (Ops) FilterTo(from []E, to []E, pred func(E) bool) []E {
	for _, x := range from {
		if pred(x) {
			to = append(to, x)
		}
	}
	return to
}

func (ops Ops) Filter(a []E, pred func(E) bool) []E {
	return ops.FilterTo(a, nil, pred)
}

//FilterSelf - inplace, non-allocating operation
func (ops Ops) FilterSelf(a []E, pred func(E) bool) []E {
	return ops.FilterTo(a, a[:0], pred)
}

func (Ops) Some(a []E, pred func(E) bool) bool {
	for _, x := range a {
		if pred(x) {
			return true
		}
	}
	return false
}

func (Ops) Every(a []E, pred func(E) bool) bool {
	for _, x := range a {
		if !pred(x) {
			return false
		}
	}
	return true
}
