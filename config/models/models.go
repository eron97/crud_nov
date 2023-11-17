package models

type Task struct {
	ID        int    `json:"id"`
	Task_Name string `json:"task_name"`
	Priority  string `json:"priority"`
}

type TaskPost struct {
	Task_Name string `json:"task_name" binding:"required,min=3"`
	Priority  string `json:"priority" binding:"required,min=3"`
}
