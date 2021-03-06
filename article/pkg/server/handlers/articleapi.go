package handlers

import (
	"context"

	"github.com/go-kit/kit/log"
	uuid "github.com/satori/go.uuid"

	"github.com/blog/article/pkg/dto"
	"github.com/blog/article/pkg/service"
	"github.com/go-kit/kit/log/level"
)

type articleAPIService struct {
	repo dto.Repository
	log  log.Logger
}

func NewArticleAPIServiceImpl(rep dto.Repository, l log.Logger) service.ArticleAPIService {
	return &articleAPIService{
		repo: rep,
		log:  l,
	}
}

func (s *articleAPIService) AddArticle(ctx context.Context, title, content, author string) (string, error) {
	logger := log.With(s.log, "method", "AddArticle")

	uuid := uuid.NewV4()
	id := uuid.String()

	article := &dto.Article{
		Id:      id,
		Title:   title,
		Content: content,
		Author:  author,
	}

	if err := s.repo.AddArticle(ctx, article); err != nil {
		level.Error(logger).Log("err", err)
		return "", err
	}

	logger.Log("Create Article", id)

	return id, nil
}

func (s *articleAPIService) GetArticle(ctx context.Context, id string) (*dto.Article, error) {
	logger := log.With(s.log, "method", "GetArticle")

	article, err := s.repo.GetArticle(ctx, id)

	if err != nil {
		level.Error(logger).Log("err", err)
		return &dto.Article{}, err
	}

	logger.Log("Get Article Object", id)

	return article, nil
}

func (s *articleAPIService) ListArticles(ctx context.Context) ([]*dto.Article, error) {
	logger := log.With(s.log, "method", "ListArticles")

	articles, err := s.repo.ListArticles(ctx)

	if err != nil {
		level.Error(logger).Log("err", err)
		return []*dto.Article{}, err
	}

	logger.Log("Listing all Articles")

	return articles, nil
}
