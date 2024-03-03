package model

type TodoResponse struct {
	Title  string `json:"title"`
	IsDone bool   `json:"is_done"`
}
