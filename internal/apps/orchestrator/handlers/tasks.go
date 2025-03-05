package handlers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/ThisIsHyum/lms_calculator/internal/expressionmanager"
	"github.com/ThisIsHyum/lms_calculator/internal/types"
	"github.com/labstack/echo/v4"
)

func GetTask(c echo.Context) error {
	task := expressionmanager.NextTask()
	if task.Id == 0 {
		return c.String(http.StatusNotFound, http.StatusText(http.StatusNotFound)+"\n")
	}
	return c.JSON(200, task)
}

func SendResult(c echo.Context) error {
	p := make([]byte, 1024)
	n, err := c.Request().Body.Read(p)
	if err != nil && err != io.EOF{
		return c.String(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError)+"\n")
	}

	resultRequest := types.ResultRequest{}
	err = json.Unmarshal(p[:n], &resultRequest)
	if err != nil {
		return c.String(http.StatusUnprocessableEntity, http.StatusText(http.StatusUnprocessableEntity)+"\n")
	}

	if expressionmanager.Expressions.ById(resultRequest.Id).Id == 0 {
		return c.String(http.StatusNotFound, http.StatusText(http.StatusNotFound)+"\n")
	}
	expressionmanager.Id(resultRequest.Id, resultRequest.Result)
	return c.String(http.StatusOK, http.StatusText(http.StatusOK)+"\n")
}