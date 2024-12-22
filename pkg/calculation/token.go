package calc

import (
	"fmt"
	"strconv"
	"strings"
)

var CharToTokens = map[string]tokenType{
	"(": brace_start,
	")": brace_end,
	"+": plus,
	"-": minus,
	"*": multiply,
	"/": divide,
}

func toToken(str string, isNum *bool, prevToken *token) (t token, err error, addition bool) {
	token := token{
		String: str,
	}

	tokenType, ok := CharToTokens[str]

	if !ok {
		_, err := strconv.Atoi(str)
		if err != nil {
			token.String = "error"
			return token, fmt.Errorf("expression has forbidden characters"), false
		}
		if *isNum {
			prevToken.String = prevToken.String + str
			*(isNum) = true
			return *prevToken, nil, true
		}
		token.TokenType = number
		*(isNum) = true
	} else {
		if !*(isNum) && prevToken.TokenType != brace_end && tokenType != brace_start {
			return token, fmt.Errorf("two operations in a row"), false
		}
		token.TokenType = tokenType
		*(isNum) = false
	}
	return token, nil, false
}
func tokenize(expression string) ([]token, error) {
	var isNum bool
	var prevToken token
	tokens := []token{}
	for _, char := range strings.Split(expression, "") {
		tt, err, addition := toToken(char, &isNum, &prevToken)
		prevToken = tt
		if err != nil {
			return nil, err
		}
		if addition {
			tokens[len(tokens)-1] = tt
		} else {
			tokens = append(tokens, tt)
		}
	}
	tokentype := tokens[len(tokens)-1].TokenType
	if tokentype != number && tokentype != brace_end && tokentype != brace_start {
		return tokens, fmt.Errorf("expression is wrong")
	}
	return tokens, nil
}
