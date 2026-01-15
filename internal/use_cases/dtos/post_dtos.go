package dtos

import "time"

type PostInput struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type PostOutput struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}
