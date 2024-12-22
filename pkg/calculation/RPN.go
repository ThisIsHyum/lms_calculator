package calc

import (
	"fmt"
	"slices"
)

func toRPN(tokens []token) ([]token, error) {
	s := []token{}
	stack := Stack{}
	for _, token := range tokens {
		if token.TokenType == brace_start {
			stack.Push(token)
			continue
		}

		if token.TokenType == brace_end {
			var brace_obtained bool
			for range stack.items {
				if stack.Top().TokenType == brace_start {
					stack.Pop()
					brace_obtained = true
					break
				}
				s = append(s, stack.Pop())
			}
			if !brace_obtained {
				return s, fmt.Errorf("brace not closed")
			}
			continue
		}

		if token.TokenType == number {
			s = append(s, token)
			continue
		}
		if stack.Top().TokenType.getPriority() >= token.TokenType.getPriority() && stack.Top().TokenType.getPriority() != 0 {
			for range stack.items {
				if stack.Top().TokenType == brace_start {
					continue
				}
				s = append(s, stack.Pop())
			}

		}
		stack.Push(token)
	}
	slices.Reverse(stack.items)
	s = append(s, stack.items...)
	return s, nil
}
