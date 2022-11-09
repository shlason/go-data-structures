package set

type element interface{}

type set struct {
	items map[element]bool
}

func New() *set {
	s := new(set)
	s.items = make(map[element]bool)
	return s
}

func (s *set) Has(e element) bool {
	_, ok := s.items[e]
	return ok
}

func (s *set) Add(e element) bool {
	if s.Has(e) {
		return false
	}
	s.items[e] = true
	return true
}

func (s *set) Remove(e element) bool {
	if s.Has(e) {
		delete(s.items, e)
		return true
	}
	return false
}

func (s *set) Clear() {
	s.items = make(map[element]bool)
}

func (s *set) Size() int {
	return len(s.items)
}

func (s *set) Values() []element {
	keys := make([]element, len(s.items))
	for k := range s.items {
		keys = append(keys, k)
	}
	return keys
}

func (s *set) Union(otherSet *set) *set {
	uSet := new(set)
	for k := range s.items {
		uSet.Add(k)
	}
	for k := range otherSet.items {
		uSet.Add(k)
	}
	return uSet
}

func (s *set) Intersection(otherSet *set) *set {
	var base *set
	var target *set

	iSet := new(set)

	if s.Size() > otherSet.Size() {
		base = otherSet
		target = s
	} else {
		base = s
		target = otherSet
	}

	for k := range base.items {
		if target.Has(k) {
			iSet.Add(k)
		}
	}

	return iSet
}

func (s *set) Diff(otherSet *set) *set {
	dSet := new(set)

	for k := range s.items {
		if !otherSet.Has(k) {
			dSet.Add(k)
		}
	}

	return dSet
}

func (s *set) IsSubset(otherSet *set) bool {
	if s.Size() > otherSet.Size() {
		return false
	}
	for k := range s.items {
		if !otherSet.Has(k) {
			return false
		}
	}
	return true
}
