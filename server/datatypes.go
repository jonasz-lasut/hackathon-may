package server

import (
	"net/http"

	"github.com/go-chi/render"
)

/*
Data structs
*/
type article struct {
	AuthorID uint   `json:"author_id,omitempty" db:"author_id"`
	ID       uint   `json:"id,omitempty" db:"id,omitempty"`
	Title    string `json:"title,omitempty" db:"title"`
}

func (a *article) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

type errorResponse struct {
	ErrorMessage string `json:"error"`
}

type response struct {
	ResponseMessage string `json:"msg"`
}

/*
Data funcs
*/
func newArticleList(articles []*article) []render.Renderer {
	list := []render.Renderer{}
	for _, art := range articles {
		list = append(list, art)
	}
	return list
}
