package models

type Task struct {
	ID        int    `json:"id"`
	Task_Name string `json:"task_name"`
	Priority  string `json:"priority"`
}
