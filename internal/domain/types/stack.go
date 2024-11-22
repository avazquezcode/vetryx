package types

type Stack []interface{}

func NewStack() Stack {
	var s Stack
	return s
}

func (s *Stack) Push(element interface{}) {
	*s = append(*s, element)
}

func (s *Stack) Pop() {
	if s.Length() == 0 {
		return // nothing to pop
	}
	*s = (*s)[:s.Length()-1]
}

func (s *Stack) Peek() interface{} {
	return s.getElementAt(s.Length() - 1)
}

func (s *Stack) Length() int {
	return len(*s)
}

func (s *Stack) getElementAt(index int) interface{} {
	return (*s)[index]
}
