package models

import (
	"context"
	"time"
)

type Article struct {
	ID        int       `json:"id"`
	Image     string    `json:"image,omitempty"`
	Title     string    `json:"title"`
	Text      string    `json:"text"`
	AuthorID  int       `json:"author_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (a *Article) GetAll() ([]*Article, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select id, image, title, text, author_id, created_at, updated_at from articles order by created_at`

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var articles []*Article

	for rows.Next() {
		var article Article
		err := rows.Scan(
			&article.ID,
			&article.Image,
			&article.Title,
			&article.Text,
			&article.AuthorID,
			&article.CreatedAt,
			&article.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		articles = append(articles, &article)
	}

	return articles, nil
}

func (a *Article) GetArticle(id int) (*Article, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select id, image, title, text, author_id, created_at, updated_at from articles where id = $1`

	var article Article
	err := db.QueryRowContext(ctx, query, id).Scan(
		&article.ID,
		&article.Image,
		&article.Title,
		&article.Text,
		&article.AuthorID,
		&article.CreatedAt,
		&article.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &article, nil
}

func (a *Article) Create(article Article) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	var newID int
	stmt := `insert into articles (image, title, text, author_id, created_at, updated_at)
		values ($1, $2, $3, $4, $5, $6) returning id`

	err := db.QueryRowContext(ctx, stmt,
		article.ID,
		article.Image,
		article.Title,
		article.Text,
		time.Now(),
		time.Now(),
	).Scan(&newID)

	if err != nil {
		return 0, err
	}

	return newID, nil
}
