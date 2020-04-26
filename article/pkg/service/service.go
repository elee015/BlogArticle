package service

import (
	"context"

	"github.com/blog/article/pkg/dto"
)

// ArticleService describes the service.
type ArticleAPIService interface {
	// Add your methods here
	AddArticle(ctx context.Context, title, content, author string) (string, error)
	GetArticle(ctx context.Context, id string) (*dto.Article, error)
	// ListArticles(ctx context.Context) ([]*dto.Article, error)
}
