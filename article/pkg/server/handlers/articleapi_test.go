package handlers_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/blog/article/pkg/dto"
	"github.com/blog/article/pkg/server/handlers"
	"github.com/blog/article/pkg/service"
	"github.com/go-kit/kit/log"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAddArticle(t *testing.T) {

	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
	}

	var mockRepo dto.MockRepository
	var svc service.ArticleAPIService

	err := errors.New("anything")

	type args struct {
		title   string
		content string
		author  string
	}
	scenarios := []struct {
		desc     string
		args     args
		doMock   func()
		expected error
	}{
		{
			desc: "happy path",
			args: args{title: "test_title1", content: "test_content1", author: "test_author1"},
			doMock: func() {
				mockRepo = dto.MockRepository{}
				svc = handlers.NewArticleAPIServiceImpl(&mockRepo, logger)
				mockRepo.On("AddArticle", mock.Anything, mock.Anything).Return(nil)
			},
			expected: nil,
		},
		{
			desc: "unhappy path",
			args: args{title: "test_title2", content: "test_content2", author: "test_author2"},
			doMock: func() {
				mockRepo = dto.MockRepository{}
				svc = handlers.NewArticleAPIServiceImpl(&mockRepo, logger)
				mockRepo.On("AddArticle", mock.Anything, mock.Anything).Return(err)
			},
			expected: err,
		},
	}
	for _, scenario := range scenarios {
		t.Run(scenario.desc, func(t *testing.T) {
			scenario.doMock()
			id, err := svc.AddArticle(context.Background(), scenario.args.title, scenario.args.content, scenario.args.author)

			if err != nil {
				assert.Equal(t, scenario.expected, err)
				assert.Empty(t, id)
			} else {
				assert.Equal(t, scenario.expected, err)
				assert.NotEmpty(t, id)
			}
		})
	}
}

func TestGetArticle(t *testing.T) {

	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
	}

	err := errors.New("not found")
	mockRepo := &dto.MockRepository{}
	mockArticle := &dto.Article{
		Id:      "test_id1",
		Title:   "test_title1",
		Content: "test_content1",
		Author:  "test_author1",
	}

	svc := handlers.NewArticleAPIServiceImpl(mockRepo, logger)
	mockRepo.On("GetArticle", mock.Anything, mockArticle.Id).Return(mockArticle, nil)
	mockRepo.On("GetArticle", mock.Anything, "test_id2").Return(&dto.Article{}, err)

	type args struct {
		id string
	}
	scenarios := []struct {
		desc     string
		args     args
		expected *dto.Article
	}{
		{
			desc:     "happy path",
			args:     args{id: "test_id1"},
			expected: mockArticle,
		},
		{
			desc:     "unhappy path - not found",
			args:     args{id: "test_id2"},
			expected: nil,
		},
	}
	for _, scenario := range scenarios {
		t.Run(scenario.desc, func(t *testing.T) {
			article, getArtError := svc.GetArticle(context.Background(), scenario.args.id)

			if getArtError != nil {
				assert.Equal(t, err, getArtError)
			} else {
				assert.Equal(t, scenario.expected, article)
			}
		})
	}
}

func TestListArticles(t *testing.T) {

	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
	}

	err := errors.New("anything")

	var mockRepo dto.MockRepository
	var svc service.ArticleAPIService

	mockArticleList := []*dto.Article{
		{
			Id:      "test_id1",
			Title:   "test_title1",
			Content: "test_content1",
			Author:  "test_author1",
		},
		{
			Id:      "test_id2",
			Title:   "test_title2",
			Content: "test_content2",
			Author:  "test_author2",
		},
	}

	scenarios := []struct {
		desc     string
		doMock   func()
		expected []*dto.Article
	}{
		{
			desc: "happy path",
			doMock: func() {
				mockRepo = dto.MockRepository{}
				svc = handlers.NewArticleAPIServiceImpl(&mockRepo, logger)
				mockRepo.On("ListArticles", mock.Anything).Return(mockArticleList, nil)
			},
			expected: mockArticleList,
		},
		{
			desc: "unhappy path - not found",
			doMock: func() {
				mockRepo = dto.MockRepository{}
				svc = handlers.NewArticleAPIServiceImpl(&mockRepo, logger)
				mockRepo.On("ListArticles", mock.Anything).Return([]*dto.Article{}, err)
			},
			expected: []*dto.Article{},
		},
	}
	for _, scenario := range scenarios {
		t.Run(scenario.desc, func(t *testing.T) {
			scenario.doMock()

			articleList, ListArtError := svc.ListArticles(context.Background())

			if ListArtError != nil {
				assert.Equal(t, err, ListArtError)
			} else {
				assert.Equal(t, scenario.expected, articleList)
			}
		})
	}
}
