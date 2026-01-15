package html

import (
	"net/http"

	"github.com/otaviouu/blogoo/cmd/html/templates"
	"github.com/otaviouu/blogoo/internal/entity/posts"
)

type HtmlPostsHandler struct {
	PostsRepo posts.IPostRepo
}

func NewHtmlPostsHandler(postRepo posts.IPostRepo) *HtmlPostsHandler {
	return &HtmlPostsHandler{
		PostsRepo: postRepo,
	}
}

func (h *HtmlPostsHandler) HandleIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err := templates.Navbar().Render(w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
