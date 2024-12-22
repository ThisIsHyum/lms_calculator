package calc

import (
	"fmt"
	"strconv"
)

func Answer(tokens []token) (float64, error) {
	stack := Stack{}
	for _, t := range tokens {
		if t.TokenType == number {
			stack.Push(t)
			continue
		}
		t1 := stack.Pop()
		t2 := stack.Pop()

		tn1, _ := strconv.ParseFloat(t1.String, 64)
		tn2, _ := strconv.ParseFloat(t2.String, 64)
		var exp float64 = 0
		switch t.TokenType {
		case plus:
			exp = tn1 + tn2
		case minus:
			exp = tn2 - tn1
		case multiply:
			exp = tn1 * tn2
		case divide:
			if tn1 == 0 {
				return 0, fmt.Errorf("divided on zero")
			}
			exp = tn2 / tn1
		}

		stack.Push(
			token{
				TokenType: number,
				String:    strconv.FormatFloat(exp, 'f', 2, 64),
			},
		)
	}
	ans, _ := strconv.ParseFloat(stack.Top().String, 64)
	return float64(ans), nil
}
func Calc(expression string) (float64, error) {
	if expression == "" {
		return 0, fmt.Errorf("expression is empty")
	}
	tokens, err := tokenize(expression)
	if err != nil {
		return 0, err
	}
	rpn, err := toRPN(tokens)
	if err != nil {
		return 0, err
	}
	return Answer(rpn)
}
