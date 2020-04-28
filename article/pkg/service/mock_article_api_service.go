// Code generated by mockery v1.0.0. DO NOT EDIT.

package service

import (
	context "context"

	dto "github.com/blog/article/pkg/dto"
	mock "github.com/stretchr/testify/mock"
)

// MockArticleAPIService is an autogenerated mock type for the ArticleAPIService type
type MockArticleAPIService struct {
	mock.Mock
}

// AddArticle provides a mock function with given fields: ctx, title, content, author
func (_m *MockArticleAPIService) AddArticle(ctx context.Context, title string, content string, author string) (string, error) {
	ret := _m.Called(ctx, title, content, author)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string) string); ok {
		r0 = rf(ctx, title, content, author)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string, string) error); ok {
		r1 = rf(ctx, title, content, author)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetArticle provides a mock function with given fields: ctx, id
func (_m *MockArticleAPIService) GetArticle(ctx context.Context, id string) (*dto.Article, error) {
	ret := _m.Called(ctx, id)

	var r0 *dto.Article
	if rf, ok := ret.Get(0).(func(context.Context, string) *dto.Article); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dto.Article)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListArticles provides a mock function with given fields: ctx
func (_m *MockArticleAPIService) ListArticles(ctx context.Context) ([]*dto.Article, error) {
	ret := _m.Called(ctx)

	var r0 []*dto.Article
	if rf, ok := ret.Get(0).(func(context.Context) []*dto.Article); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*dto.Article)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}