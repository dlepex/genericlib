package slice

// E - slice element type
type E = interface{}

// Basic - basic slice operations
type Basic struct{}

// Pred - slice operations that require predicate (filter, exists, all)
type Pred struct{}

func (Basic) Pop(a []E) ([]E, E) {
	l := len(a)
	return a[:l-1], a[l-1]
}

func (Basic) Copy(a []E) []E {
	dest := make([]E, len(a))
	copy(dest, a)
	return dest
}

func (Basic) IndexOf(a []E, elem E) int {
	for i, e := range a {
		if e == elem {
			return i
		}
	}
	return -1
}

func (b Basic) Contains(a []E, elem E) bool {
	return b.IndexOf(a, elem) >= 0
}

func (b Basic) Delete(a []E, elem E) ([]E, bool) {
	idx := b.IndexOf(a, elem)
	if idx < 0 {
		return a, false
	}
	return b.DeleteAt(a, idx), true
}

func (b Basic) DeleteAt(a []E, idx int) []E { return append(a[0:idx], a[idx+1:]...) }

func (Basic) Reverse(a []E) {
	l := len(a)
	for i := l/2 - 1; i >= 0; i-- {
		opp := l - 1 - i
		a[i], a[opp] = a[opp], a[i]
	}
}

func (Pred) FilterTo(dest []E, src []E, pred func(E) bool) []E {
	for _, x := range src {
		if pred(x) {
			dest = append(dest, x)
		}
	}
	return dest
}

func (p Pred) Filter(a []E, pred func(E) bool) []E {
	return p.FilterTo(nil, a, pred)
}

// FilterMut - inplace (mutating) filtering. No allocation.s
func (p Pred) FilterMut(a *[]E, pred func(E) bool) {
	s := *a
	*a = p.FilterTo(s[:0], s, pred)
}

// FindIndex - same as IndexOf, but with custom predicate
func (Pred) FindIndex(a []E, pred func(E) bool) int {
	for i, x := range a {
		if pred(x) {
			return i
		}
	}
	return -1
}

// Exists - aka "some"
func (p Pred) Exists(a []E, pred func(E) bool) bool { return p.FindIndex(a, pred) >= 0 }

// All - aka "every"
func (Pred) All(a []E, pred func(E) bool) bool {
	for _, x := range a {
		if !pred(x) {
			return false
		}
	}
	return true
}
