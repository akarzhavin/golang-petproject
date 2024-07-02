package main

import (
	"errors"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

type ArticleResponse struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Image     string `json:"image"`
	Text      string `json:"text"`
	Editable  bool   `json:"editable"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func (app *Config) GetList(w http.ResponseWriter, r *http.Request) {
	userID, err := app.getUserID(r)
	if err != nil {
		app.errorJSON(w, err)
	}

	articles, err := app.Models.Article.GetAll()
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	var articlesResponse []*ArticleResponse
	for _, article := range articles {
		articleResponse := ArticleResponse{
			ID:        article.ID,
			Title:     article.Title,
			Image:     article.Image,
			Text:      article.Text[0:500] + "...",
			Editable:  article.AuthorID == userID,
			CreatedAt: article.CreatedAt.Format("2006 January 02"),
			UpdatedAt: article.UpdatedAt.Format("2006-01-02 15:04:05"),
		}
		articlesResponse = append(articlesResponse, &articleResponse)
	}

	var payload jsonResponse
	payload.Error = false
	payload.Data = articlesResponse

	app.writeJSON(w, http.StatusAccepted, payload)
}

func (app *Config) GetArticle(w http.ResponseWriter, r *http.Request) {
	_, err := app.getUserID(r)
	if err != nil {
		app.errorJSON(w, err)
	}

	strArticleId := chi.URLParam(r, "id")
	articleId, err := strconv.Atoi(strArticleId)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	article, err := app.Models.Article.GetArticle(articleId)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	if article == nil {
		app.errorJSON(w, errors.New("Article not found"), http.StatusNotFound)
		return
	}

	var articleResponse ArticleResponse
	articleResponse.ID = article.ID
	articleResponse.Title = article.Title
	articleResponse.Image = article.Image
	articleResponse.Text = article.Text
	articleResponse.Editable = true
	articleResponse.CreatedAt = article.CreatedAt.Format("2006 January 02")
	articleResponse.UpdatedAt = article.UpdatedAt.Format("2006-01-02 15:04:05")

	var payload jsonResponse
	payload.Error = false
	payload.Data = articleResponse

	app.writeJSON(w, http.StatusAccepted, payload)
}
