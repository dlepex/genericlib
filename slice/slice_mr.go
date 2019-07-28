package slice

// D - destination type
type D = interface{}

type Map struct{}

type Reduce struct{}

func (Map) ApplyTo(dest []D, src []E, f func(E) D) []D {
	for _, x := range src {
		dest = append(dest, f(x))
	}
	return dest
}

func (m Map) Apply(a []E, f func(E) D) []D { return m.ApplyTo(nil, a, f) }

func (Reduce) Apply(a []E, initial D, reducer func(D, E) D) D {
	d := initial
	for _, x := range a {
		d = reducer(d, x)
	}
	return d
}
