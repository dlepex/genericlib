package set

type E = interface{}

type Set map[E]struct{}

func New() Set {
	return make(map[E]struct{})
}

func (s Set) Add(elem E) Set {
	s[elem] = struct{}{}
	return s
}

func (s Set) AddMany(a ...E) Set {
	for _, x := range a {
		s[x] = struct{}{}
	}
	return s
}

func (s Set) Contains(elem E) bool {
	_, ok := s[elem]
	return ok
}

func (s Set) Copy() Set {
	m := make(map[E]struct{})
	for k := range s {
		m[k] = struct{}{}
	}
	return m
}

func (s Set) AddSet(other Set) {
	for k := range other {
		s[k] = struct{}{}
	}
}

func (s Set) Subtract(other Set) {
	for k := range other {
		delete(s, k)
	}
}
