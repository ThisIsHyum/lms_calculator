package types

type Status string

const (
	NotResolved Status = "not resolved"
	Solved Status = "solved"
)
type ResultRequest struct {
	Id int `json:"id"`
	Result int `json:"result"`
}