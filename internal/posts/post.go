package posts

import (
	"time"

	"github.com/otaviouu/blogoo/internal/use_cases/dtos"
)

type IPostRepo interface {
	GetPostById(id int) (*dtos.PostOutput, error)
	ListPosts() ([]*dtos.PostOutput, error)
}

type Post struct {
	ID        int
	Title     string
	Content   string
	CreatedAt time.Time
}
