package usercases

import (
	"github.com/otaviouu/blogoo/internal/posts"
	"github.com/otaviouu/blogoo/user_cases/dtos"
)

type ListPostsUseCase struct {
	postRepo posts.IPostRepo
}

func NewListPostsUseCase(postRepo posts.IPostRepo) (*ListPostsUseCase, error) {
	return &ListPostsUseCase{
		postRepo: postRepo,
	}, nil
}

func (lp *ListPostsUseCase) Execute() ([]*dtos.PostOutput, error) {
	posts, err := lp.postRepo.ListPosts()
	if err != nil {
		return nil, err
	}

	return posts, nil
}
