package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/ThisIsHyum/lms_calculator/internal/expressionmanager"
	calc "github.com/ThisIsHyum/lms_calculator/pkg/calculation"
	"github.com/labstack/echo/v4"
)

type Request struct {
	Expression string `json:"expression"`
}

type ExpressionsRequest struct {
	Expressions []expressionmanager.Expression `json:"expression"`
}
func AddExpression(c echo.Context) error {
	p := make([]byte, 1024)
	n, err := c.Request().Body.Read(p)
	if err != nil && err != io.EOF{
		return c.String(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError)+"\n")
	}
	
	request := Request{}
	err = json.Unmarshal(p[:n], &request)
	if err != nil {
		return c.String(http.StatusUnprocessableEntity, http.StatusText(http.StatusUnprocessableEntity)+"\n")
	}

	tokens, err := calc.Tokenize(request.Expression)
	if err != nil {
		return c.String(http.StatusUnprocessableEntity, "error: "+err.Error()+"\n")
	}

	rpn, err := calc.ToRPN(tokens)
	if err != nil {
		return c.String(http.StatusUnprocessableEntity, "error: "+err.Error()+"\n")
	}

	tasks := calc.ToTasks(rpn)
	expressionmanager.Expressions.Add(tasks)
	
	return c.String(http.StatusCreated, http.StatusText(http.StatusCreated)+"\n")
}

func GetExpressions(c echo.Context) error {
	return c.JSON(200, ExpressionsRequest{Expressions: expressionmanager.Expressions})
}

func GetExpression(c echo.Context) error {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)
	expression := expressionmanager.Expressions.ById(id)
	if expression.Id == 0 {
		return c.String(http.StatusNotFound, http.StatusText(http.StatusNotFound)+"\n")
	}
	return c.JSON(http.StatusOK, expressionmanager.Expressions.ById(id))
}
