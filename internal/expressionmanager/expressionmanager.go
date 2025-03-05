package expressionmanager

import (
	"slices"
	"strconv"

	"github.com/ThisIsHyum/lms_calculator/internal/taskmanager"
	"github.com/ThisIsHyum/lms_calculator/internal/types"
)

type Expression struct {
	Id int `json:"id"`
	Status types.Status `json:"status"`
	Result int `json:"result"`
	Tasks []taskmanager.Task `json:"-"`
}

type expressions []Expression

func Id(id, result int) {
	taskString := "t" + strconv.Itoa(id)
	for i, expression := range Expressions {
		if len(expression.Tasks) == 0 && expression.Status == types.NotResolved {
			Expressions[i].Result = result
			Expressions[i].Status = types.Solved
			return 
		}
		for j, task := range expression.Tasks {
			if task.Arg1 == taskString {
				Expressions[i].Tasks[j].Arg1 = strconv.Itoa(result)
			}
			if task.Arg2 == taskString {
				Expressions[i].Tasks[j].Arg2 = strconv.Itoa(result)
			}
		}
	}
}

func (e expressions) ById(id int) Expression{
	for _, expression := range e {
		if expression.Id == id {
			return expression
		}
	}
	return Expression{}
}

func (e *expressions) Add(tasks taskmanager.Tasks) {
	id := 1
	if len(*e) != 0 {
		id = (*e)[len(*e)-1].Id+1
	}
	
	*e = append(*(e), Expression{
		Id: id,
		Status: types.NotResolved,
		Tasks: tasks,
	})
}
var Expressions = expressions{}

func NextTask() taskmanager.Task{
	if len(Expressions) == 0 {
		return taskmanager.Task{}
	}
	for i, expression := range Expressions {
		if expression.Status == types.Solved {
			continue
		}
		for j, task := range expression.Tasks {
			if string(task.Arg1[0]) != "t" && string(task.Arg2[0]) != "t" {
				Expressions[i].Tasks = slices.Delete(Expressions[i].Tasks, j, j+1)
				return task
			}
		}	
	}
	return taskmanager.Task{}
}