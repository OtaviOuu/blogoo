package usercases

import (
	"time"

	"github.com/otaviouu/blogoo/internal/entity/posts"
	"github.com/otaviouu/blogoo/internal/use_cases/dtos"
)

type PublishPostUseCase struct {
	postRepo posts.IPostRepo
}

func (pb *PublishPostUseCase) Execute(postInput *dtos.PostInput) (*dtos.PostOutput, error) {
	return &dtos.PostOutput{
		Title:     postInput.Title,
		Content:   postInput.Content,
		CreatedAt: time.Now(),
	}, nil
}

func NewPublishPostUseCase(postRepo posts.IPostRepo) *PublishPostUseCase {
	return &PublishPostUseCase{
		postRepo: postRepo,
	}
}
