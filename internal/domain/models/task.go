package models

type Task struct {
	ID          int     `json:"id"`
	UserId      string  `json:"user_id"`
	Data        string  `json:"data"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
	CompletedAt float64 `json:"completed_at"`
}
