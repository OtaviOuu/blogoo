package db

import (
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/otaviouu/blogoo/internal/posts"
	"github.com/otaviouu/blogoo/user_cases/dtos"
)

type PostRepository struct {
	Db *sqlx.DB
}

func NewPostRepository(db *sqlx.DB) (posts.IPostRepo, error) {
	err := db.Ping()
	if err != nil {
		return nil, err
	}

	return &PostRepository{
		Db: db,
	}, nil
}

func (pr *PostRepository) GetPostById(id int) (*dtos.PostOutput, error) {
	return &dtos.PostOutput{
		Title:     "eita trem bão",
		Content:   "bão de mais da conta sô",
		CreatedAt: time.Now(),
	}, nil
}

func (pr *PostRepository) ListPosts() ([]*dtos.PostOutput, error) {
	return []*dtos.PostOutput{
		{
			Title:     "eita trem bão",
			Content:   "bão de mais da conta sô",
			CreatedAt: time.Now(),
		},
		{
			Title:     "eita trem bão",
			Content:   "bão de mais da conta sô",
			CreatedAt: time.Now(),
		},
	}, nil
}
