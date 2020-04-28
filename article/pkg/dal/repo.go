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

	var artDALObj dto.Article

	err := repo.db.QueryRow("SELECT * FROM articles WHERE id=$1;", id).Scan(&artDALObj.Id, &artDALObj.Title, &artDALObj.Content, &artDALObj.Author)

	switch err {
	case sql.ErrNoRows:
		return &dto.Article{}, err
	case nil:
		return &artDALObj, nil
	default:
		return &dto.Article{}, RepoError
	}
}

func (repo *repo) ListArticles(ctx context.Context) ([]*dto.Article, error) {

	var articles []*dto.Article

	rows, err := repo.db.Query("SELECT * FROM articles LIMIT $1;", 80)

	if rows == nil {
		return []*dto.Article{}, nil
	}

	for rows.Next() {
		var artDALObj dto.Article
		err = rows.Scan(&artDALObj.Id, &artDALObj.Title, &artDALObj.Content, &artDALObj.Author)
		if err != nil {
			return []*dto.Article{}, RepoError
		}
		articles = append(articles, &artDALObj)
	}

	return articles, nil
}
