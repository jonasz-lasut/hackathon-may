package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

/*
Routing
*/
func (h DatabaseHandler) HealthcheckHandler(w http.ResponseWriter, r *http.Request) {
	if h.DB.Ping() != nil {
		generateErrorResponse(w, "Connection to the database has dropped", http.StatusInternalServerError)
		return
	}
	w.WriteHeader((http.StatusOK))
	w.Write([]byte("Everything looks healthy"))
}

func (h DatabaseHandler) ArticleGetter(w http.ResponseWriter, r *http.Request) {
	articleID := chi.URLParam(r, "articleID")
	if articleID == "" {
		generateErrorResponse(w, "Request ID was nil", http.StatusBadRequest)
		return
	}

	article, err := h.dbGetArticleByID(articleID)
	if err != nil {
		generateErrorResponse(w, fmt.Sprintf("article with ID: %s was not found", articleID), http.StatusNotFound)
		return
	}

	body, _ := json.Marshal(article)
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func (h DatabaseHandler) ArticleCreator(w http.ResponseWriter, r *http.Request) {
	var a article

	err := json.NewDecoder(r.Body).Decode(&a)
	if err != nil {
		generateErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	if a.Title == "" || a.AuthorID == 0 {
		generateErrorResponse(w, "`title` and `author_id` can not be empty", http.StatusBadRequest)
		return
	}

	id, err := h.dbCreateArticle(a)
	if err != nil {
		generateErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	a.ID = id

	//remove mock call to comercial attacher: lasut
	err = attachComercials()
	if err != nil {
		generateErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//TODO

	//remove mock call to propagate: lasut
	err = propagateNotification()
	if err != nil {
		generateErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//TODO

	body, _ := json.Marshal(a)
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func (h DatabaseHandler) ArticleListGetter(w http.ResponseWriter, r *http.Request) {
	artList, err := h.dbGetArticleList()
	if err != nil {
		generateErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	render.RenderList(w, r, newArticleList(artList))
}

func (h DatabaseHandler) ArticleUpdater(w http.ResponseWriter, r *http.Request) {
	articleID := chi.URLParam(r, "articleID")
	if articleID == "" {
		generateErrorResponse(w, "Request ID was nil", http.StatusBadRequest)
		return
	}

	var a article

	err := json.NewDecoder(r.Body).Decode(&a)
	if err != nil {
		generateErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.dbUpdateArticle(articleID, a)
	if err != nil {
		generateErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	body, _ := json.Marshal(a)
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func (h DatabaseHandler) articleDeleter(w http.ResponseWriter, r *http.Request) {
	articleID := chi.URLParam(r, "articleID")
	if articleID == "" {
		generateErrorResponse(w, "Request ID was nil", http.StatusBadRequest)
		return
	}

	err := h.dbDeleteArticle(articleID)
	if err != nil {
		generateErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	res := response{ResponseMessage: fmt.Sprintf("article with ID: %s was deleted", articleID)}
	body, _ := json.Marshal(res)
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

/*
Admin routing
*/
func AdminRouter(h DatabaseHandler) http.Handler {
	r := chi.NewRouter()
	r.Use(adminOnly)
	r.Delete("/article/{articleID}", h.articleDeleter)
	return r
}

func adminOnly(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		perm, err := validatePermissions()
		if !perm || err != nil {
			generateErrorResponse(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}

/*
Helpers
*/
func generateErrorResponse(w http.ResponseWriter, err string, errorCode int) {
	res := errorResponse{ErrorMessage: err}
	body, _ := json.Marshal(res)
	log.Printf("[ERROR] Sending error response: %s", body)
	w.WriteHeader(errorCode)
	w.Write(body)
}
