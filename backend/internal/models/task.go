package models

import "time"

type Task struct {
	Id         int       `json:"id"`
	UserId     int       `json:"user_id"`
	TaskName   string    `json:"task_name"`
	StartedAt  time.Time `json:"started_at"`
	FinishedAt time.Time `json:"finished_at,omitempty"`
	Duration   int       `json:"duration,int,omitempty"`
}

type ResponseTasks struct {
	Tasks []*Task `json:"tasks"`
}
