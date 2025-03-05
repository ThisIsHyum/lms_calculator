package client

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/ThisIsHyum/lms_calculator/internal/taskmanager"
	"github.com/ThisIsHyum/lms_calculator/internal/types"
)

type Client struct {
	BaseURL string
	HTTPclient http.Client
}

func New(baseURL string) Client {
	return Client{
		BaseURL: "http://" + baseURL,
		HTTPclient: http.Client{},
	}
}

func (c Client) GetTask() (taskmanager.Task, error) {
	resp, err := c.HTTPclient.Get(c.BaseURL+"/internal/task")

	if err != nil {
		return taskmanager.Task{}, err
	}

	task := taskmanager.Task{}
	if resp.StatusCode == http.StatusOK {
		p := make([]byte, 1024)
		n, _ := resp.Body.Read(p)
		json.Unmarshal(p[:n], &task)
	}
	return task, nil
}

func (c Client) SendResult(id, result int) error {
	b, _ := json.Marshal(types.ResultRequest{
		Id: id, 
		Result: result,
	})
	_, err := c.HTTPclient.Post(c.BaseURL+"/internal/task", "application/json", bytes.NewReader(b))
	return err
}