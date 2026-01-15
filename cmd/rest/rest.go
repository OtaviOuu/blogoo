package rest

import (
	"encoding/json"
	"net/http"

	"github.com/otaviouu/blogoo/internal/posts"
	usercases "github.com/otaviouu/blogoo/user_cases"
)

type RestPostsHandler struct {
	PostsRepo posts.IPostRepo
}

func NewRestPostsHandler(postsRepo posts.IPostRepo) *RestPostsHandler {
	return &RestPostsHandler{
		PostsRepo: postsRepo,
	}
}

func (h *RestPostsHandler) HandlePublishPost(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("posts"))
}

func (h *RestPostsHandler) HandleListPosts(w http.ResponseWriter, r *http.Request) {
	useCase, err := usercases.NewListPostsUseCase(h.PostsRepo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	posts, err := useCase.Execute()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(posts)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
