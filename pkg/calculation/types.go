package calc

type tokenType int

const (
	number tokenType = iota

	brace_start
	brace_end

	plus
	minus
	multiply
	divide
)

func (t tokenType) getPriority() int {
	switch t {
	case number:
		return 0
	case brace_start:
		return -5
	case brace_end:
		return -4
	case plus:
		return 1
	case minus:
		return 1
	case multiply:
		return 2
	case divide:
		return 2
	}
	return -1
}

type token struct {
	TokenType tokenType
	String    string
}

type Stack struct {
	items []token
}

func (s *Stack) Push(data token) {
	s.items = append(s.items, data)
}

func (s *Stack) Pop() token {
	if len(s.items) == 0 {
		return token{}
	}
	l := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return l
}

func (s *Stack) Top() token {
	if len(s.items) == 0 {
		return token{TokenType: number}
	}
	return s.items[len(s.items)-1]
}
