package models

import "time"

type Task struct {
	Id         int       `json:"id" example:"1"`
	UserId     int       `json:"user_id" example:"1"`
	TaskName   string    `json:"task_name" example:"processing mail inbox"`
	StartedAt  time.Time `json:"started_at" example:"2024-06-30 14:30:01.001"`
	FinishedAt time.Time `json:"finished_at" example:"2024-06-30 14:59:06.113"`
	Duration   int       `json:"duration,int" example:"1741"`
}

type RequestNewTask struct {
	UserId   int    `json:"user_id" example:"1"`
	TaskName string `json:"task_name" example:"processing mail inbox"`
}

type ResponseTasks struct {
	Tasks []*Task `json:"tasks"`
}
