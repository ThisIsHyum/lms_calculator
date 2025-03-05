package calc

import (
	"strconv"
	"time"

	"github.com/ThisIsHyum/lms_calculator/internal/taskmanager"
)

func Calculate(task taskmanager.Task) int {
	time.Sleep(task.Operation_time)
	arg1, _ := strconv.Atoi(task.Arg1)
	arg2, _ := strconv.Atoi(task.Arg2)
	switch task.Operation {
	case "+":
		return arg1 + arg2
	case "-":
		return arg1 - arg2
	case "*":
		return arg1 * arg2
	case "/":
		return arg1 / arg2
	}
	return 0
}