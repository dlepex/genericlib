package conv

type E = interface{}

type K = interface{}
type V = interface{}

type MapKeys struct{}
type MapValues struct{}
type SliceToMap struct{}
type SliceToSet struct{}
type SliceToChan struct{}
type SetToSlice struct{}

func (MapKeys) Apply(m map[K]V) []K {
	a := make([]K, 0, len(m))
	for k := range m {
		a = append(a, k)
	}
	return a
}

func (MapValues) Apply(m map[K]V) []V {
	a := make([]V, 0, len(m))
	for _, v := range m {
		a = append(a, v)
	}
	return a
}

func (SliceToMap) Apply(a []V, key func(V) K) map[K]V {
	m := make(map[K]V)
	for _, v := range a {
		m[key(v)] = v
	}
	return m
}

func (SliceToSet) Apply(a []E) map[E]struct{} {
	m := make(map[E]struct{})
	for _, e := range a {
		m[e] = struct{}{}
	}
	return m
}

func (SetToSlice) Apply(s map[E]struct{}) []E {
	a := make([]E, 0, len(s))
	for e := range s {
		a = append(a, e)
	}
	return a
}
