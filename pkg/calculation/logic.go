package calc

import (
	"github.com/ThisIsHyum/lms_calculator/internal/taskmanager"
)

func ToTasks(tokens []token) taskmanager.Tasks {
	var tasks = taskmanager.Tasks{}
	stack := Stack{}
	for _, t := range tokens {
		if t.TokenType == number {
			stack.Push(t)
			continue
		}
		t1 := stack.Pop()
		t2 := stack.Pop()
		task := tasks.Add(t1.String, t2.String, t.String)
		stack.Push(
			token{
				TokenType: number,
				String:    task.String(),
			},
		)
	}
	return tasks
}
