package taskmanager

import (
	"fmt"
	"strconv"
	"time"

	"github.com/ThisIsHyum/lms_calculator/internal/config"
)

type Task struct {
	ExpressionId int `json:"-"`
	Id int `json:"id"`
	Arg1 string `json:"arg1"`
	Arg2 string `json:"arg12"`
	Operation string `json:"operation"`
	Operation_time time.Duration `json:"operation_time"`
}

type Tasks []Task
var GlobalTasks = []Tasks{}

/*
func NextTask() Task{
	if len(GlobalTasks) == 0 {
		return Task{}
	}
	for i, task := range GlobalTasks[0] {
		if string(task.Arg1[0]) != "t" && string(task.Arg2[0]) != "t" {
			GlobalTasks = slices.Delete(GlobalTasks, i, i+1)
			return task
		}
	}
	return Task{}
}
*/

func (t *Tasks) Add(arg1, arg2, operation string) Task{
	id := 1
	if len(*t) != 0 {
		id = (*t)[len(*t)-1].Id+1
	}
	var duration time.Duration
	switch operation {
		case "+":
			duration, _ = time.ParseDuration(strconv.Itoa(config.Config.TimeAdditionMs) + "ms")
		case "-":
			duration, _ = time.ParseDuration(strconv.Itoa(config.Config.TimeSubtractionMs) + "ms")
		case "*":
			duration, _ = time.ParseDuration(strconv.Itoa(config.Config.TimeMultiplicationsMs) + "ms")
		case "/":
			duration, _ = time.ParseDuration(strconv.Itoa(config.Config.TimeDivisionsMs) + "ms")
	}
	task := Task{
		Id: id,
		Arg1: arg1,
		Arg2: arg2,
		Operation: operation,
		Operation_time: duration,
	}
	*t = append(*(t), task)
	return task
}

func (t Task) String() string {
	return fmt.Sprintf("t%d", t.Id)
}