package dal

import (
	"context"
	"database/sql"
	"errors"

	"github.com/go-kit/kit/log"

	"github.com/blog/article/pkg/dto"
)

var RepoError = errors.New("Unable to handle Repo Request")

type repo struct {
	db     *sql.DB
	logger log.Logger
}

func NewRepo(db *sql.DB, logger log.Logger) dto.Repository {
	return &repo{
		db:     db,
		logger: log.With(logger, "DAL", "sql"),
	}
}

func (repo *repo) AddArticle(ctx context.Context, article *dto.Article) error {
	sql := `INSERT INTO articles (id, title, content, author)
			VALUES ($1, $2, $3, $4)`

	_, err := repo.db.ExecContext(ctx, sql, article.Id, article.Title, article.Content, article.Author)
	if err != nil {
		return err
	}

	repo.logger.Log("added row", article.Id)

	return nil
}

func (repo *repo) GetArticle(ctx context.Context, id string) (*dto.Article, error) {

	foundArticle := &dto.Article{}

	err := repo.db.QueryRow("SELECT * FROM articles WHERE id=$1", id).Scan(foundArticle)
	if err != nil {
		return &dto.Article{}, RepoError
	}

	return foundArticle, nil
}
