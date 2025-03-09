package agent

import (
	"time"

	"github.com/ThisIsHyum/lms_calculator/internal/apps/agent/client"
	calc "github.com/ThisIsHyum/lms_calculator/pkg/calculation"
)

type App struct {
	client client.Client
	computingPower int
}

func New(computingPower int) App {
	return App{
		client: client.New("orchestrator:8080"),
		computingPower: computingPower,
	}
}

func (a App) Run() error {
	errCh := make(chan error)
	for range a.computingPower {
		go cycle(a.client, errCh)	
	}
	return <-errCh
}

func cycle(client client.Client, errCh chan error) {
	for {
		time.Sleep(1*time.Second)
		task, err := client.GetTask()
		if err != nil {
			errCh <- err
			return
		}
		if task.Id != 0 {
			result := calc.Calculate(task)
			err = client.SendResult(task.Id, result)
			if err != nil {
				errCh <- err
				return 
			}
		}
	}
}