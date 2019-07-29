package slice

import "github.com/dlepex/genericlib/std"

// D - destination type for map or reduce
type D = interface{}

// Map - slice map operation
type Map struct{}

// Reduce - slice reduce right
type Reduce struct{}

// ByChunks - for each chunk of slice
type ByChunks struct{}

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

func (ByChunks) Apply(a []E, chunkSize int, each func([]E) bool) {
	n := len(a)
	if n <= chunkSize {
		if n != 0 {
			each(a)
		}
		return
	}

	start := 0
	for {
		end := std.Min(start+chunkSize, n)
		if !each(a[start:end]) {
			return
		}
		if end == n {
			return
		}
		start = end
	}
}
