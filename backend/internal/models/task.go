package models

import "time"

type Task struct {
	Id         int           `json:"id"`
	UserId     int           `json:"user_id"`
	Task       string        `json:"task"`
	StartedAt  time.Time     `json:"started_at"`
	FinishedAt time.Time     `json:"finished_at,omitempty"`
	Duration   time.Duration `json:"duration,omitempty"`
}
