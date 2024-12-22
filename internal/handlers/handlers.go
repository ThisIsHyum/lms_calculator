package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	calc "github.com/ThisIsHyum/lms_calculator/pkg/calculation"
)

type Request struct {
	Expression string `json:"expression"`
}
type Response struct {
	Result string `json:"result,omitempty"`
	Error  string `json:"error,omitempty"`
}

func Calculate(w http.ResponseWriter, r *http.Request) {
	response := Response{}

	if r.Method != "POST" {
		response.Error = "Wrong method"
		w.WriteHeader(http.StatusMethodNotAllowed)
		sendResponse(w, response)
		return
	}

	request := Request{}

	data := make([]byte, 200)
	n, err := r.Body.Read(data)
	if isInternalError(w, err) {
		return
	}

	err = json.Unmarshal(data[:n], &request)
	if isInternalError(w, err) {
		return
	}

	result, err := calc.Calc(request.Expression)

	if err != nil {
		response.Error = err.Error()
		w.WriteHeader(http.StatusUnprocessableEntity)
		sendResponse(w, response)
		return
	}

	response.Result = strconv.FormatFloat(result, 'g', 32, 64)

	sendResponse(w, response)
}

func sendResponse(w http.ResponseWriter, response Response) {
	jsonResponse, _ := json.Marshal(response)
	fmt.Fprintln(w, string(jsonResponse))
}

func isInternalError(w http.ResponseWriter, err error) bool {
	if err == io.EOF {
		return false
	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		sendResponse(w, Response{Error: "Internal server error"})
		return true
	}
	return false
}
