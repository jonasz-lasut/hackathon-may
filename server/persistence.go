package server

import (
	"errors"

	"github.com/upper/db/v4"
)

/*
Data persistence handlers
*/
type DatabaseHandler struct {
	DB db.Session
}

func (h DatabaseHandler) dbGetArticleByID(id string) (*article, error) {
	var a *article
	err := h.DB.Collection("articles").Find("id=?", id).One(&a)

	return a, err
}

func (h DatabaseHandler) dbGetArticleList() ([]*article, error) {
	var aList []*article
	err := h.DB.Collection("articles").Find().All(&aList)

	return aList, err
}

func (h DatabaseHandler) dbCreateArticle(a article) (uint, error) {
	var dbArticle *article
	articleCollection := h.DB.Collection("articles")

	articleCollection.Find("author_id=? AND title=?", a.AuthorID, a.Title).One(&dbArticle)
	if dbArticle != nil {
		return dbArticle.ID, nil
	}

	res, err := articleCollection.Insert(a)
	return uint(res.ID().(int64)), err
}

func (h DatabaseHandler) dbUpdateArticle(id string, a article) error {
	var dbArticle *article
	articleResult := h.DB.Collection("articles").Find("id=?", id)

	err := articleResult.One(&dbArticle)
	if err != nil {
		return errors.New("request article is not present and therefore can not be updated")
	}

	if a.AuthorID == 0 {
		a.AuthorID = dbArticle.AuthorID
	}

	if a.Title == "" {
		a.Title = dbArticle.Title
	}

	err = articleResult.Update(a)
	return err
}

func (h DatabaseHandler) dbDeleteArticle(id string) error {
	err := h.DB.Collection("articles").Find("id=?", id).Delete()
	return err
}
