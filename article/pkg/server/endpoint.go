package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-kit/kit/endpoint"

	articlepb "github.com/blog/article/pkg/pb"
	"github.com/blog/article/pkg/service"
)

type Endpoints struct {
	AddArticle   endpoint.Endpoint
	GetArticle   endpoint.Endpoint
	ListArticles endpoint.Endpoint
}

func MakeEndpoints(s service.ArticleAPIService) Endpoints {
	return Endpoints{
		AddArticle:   makeAddArticleEndpoint(s),
		GetArticle:   makeGetArticleEndpoint(s),
		ListArticles: makeListArticleEndpoint(s),
	}
}

func makeAddArticleEndpoint(s service.ArticleAPIService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(articlepb.AddArticleRequest)

		valError := req.Validate()

		if valError != nil {
			return articlepb.AddArticleResponse{
				Status:  http.StatusBadRequest,
				Message: valError.(articlepb.AddArticleRequestValidationError).Reason(),
				Data:    nil,
			}, nil
		}

		id, err := s.AddArticle(ctx, req.Title, req.Content, req.Author)

		return articlepb.AddArticleResponse{
			Status:  http.StatusCreated,
			Message: "Success",
			Data: &articlepb.AddArticleResponse_ArticleId{
				Id: id,
			},
		}, err
	}
}

func makeGetArticleEndpoint(s service.ArticleAPIService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(articlepb.GetArticleRequest)

		article, err := s.GetArticle(ctx, req.Id)

		if err != nil {
			return articlepb.AddArticleResponse{
				Status:  http.StatusBadRequest,
				Message: "Request cannot fulfill",
				Data:    nil,
			}, nil
		}

		return articlepb.GetArticleResponse{
			Status:  http.StatusOK,
			Message: "Success",
			Data: &articlepb.Article{
				Id:      article.Id,
				Title:   article.Title,
				Content: article.Content,
				Author:  article.Author,
			},
		}, err
	}
}

func makeListArticleEndpoint(s service.ArticleAPIService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var articlesList []*articlepb.Article

		articles, err := s.ListArticles(ctx)

		if err != nil {
			return articlepb.AddArticleResponse{
				Status:  http.StatusBadRequest,
				Message: "Request cannot fulfill",
				Data:    nil,
			}, nil
		}

		for _, article := range articles {
			fmt.Println(article)
			a := &articlepb.Article{
				Id:      article.Id,
				Title:   article.Title,
				Content: article.Content,
				Author:  article.Author,
			}
			articlesList = append(articlesList, a)

		}

		return articlepb.ListArticlesResponse{
			Status:  http.StatusOK,
			Message: "Success",
			Data:    articlesList,
		}, err
	}
}
