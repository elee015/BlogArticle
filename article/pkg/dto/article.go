package dto

import (
	"context"
	"encoding/json"
	"io"
)

type Article struct {
	Id      string `json:"id,omitempty"`
	Title   string `json:"title,omitempty" validate:"string,min=1,required"`
	Content string `json:"content,omitempty" validate:"string"`
	Author  string `json:"author,omitempty" validate:"string"`
}

type Repository interface {
	AddArticle(ctx context.Context, art *Article) error
	GetArticle(ctx context.Context, id string) (*Article, error)
	// ListArticles(ctx context.Context) ([]*Article, error)
}

func (a *Article) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)

	return e.Encode(a)
}

func AddArticle(ctx context.Context, art *Article) error {
	return nil
}

func GetArticle(ctx context.Context, id string) (*Article, error) {
	return product, nil
}

var product = &Article{
	Id:      "1",
	Title:   "Test",
	Content: "Test_Content",
	Author:  "Test_AUthor",
}
