package dto

import (
	"context"
	"encoding/json"
	"io"
)

type Article struct {
	Id      string `json:"id,omitempty"`
	Title   string `json:"title,omitempty"`
	Content string `json:"content,omitempty"`
	Author  string `json:"author,omitempty"`
}

type Repository interface {
	AddArticle(ctx context.Context, art *Article) error
	GetArticle(ctx context.Context, id string) (*Article, error)
	ListArticles(ctx context.Context) ([]*Article, error)
}

func (a *Article) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)

	return e.Encode(a)
}
