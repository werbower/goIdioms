package generic01

type Stack[T comparable] struct {
	vals []T
}

func (s *Stack[T]) Push(val T) {
	s.vals = append(s.vals, val)
}

func (s *Stack[T]) Pop() (T, bool) {
	sLen := len(s.vals)

	if sLen == 0 {
		var zero T
		return zero, false
	}

	top := s.vals[sLen-1]
	s.vals = s.vals[:sLen-1]
	return top, true
}

func (s *Stack[T]) Contains(val T) bool {
	for _, v := range s.vals {
		if v == val {
			return true
		}
	}
	return false
}
