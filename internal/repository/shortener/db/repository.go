package db

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	irepo "github.com/almira-galeeva/url-shortener/internal/repository/shortener"
	"github.com/jackc/pgx/v4/pgxpool"
)

const tableName = "urls"

type repository struct {
	pool *pgxpool.Pool
}

func NewRepository(pool *pgxpool.Pool) irepo.Repository {
	return &repository{
		pool: pool,
	}
}

func (r *repository) CreateUrl(ctx context.Context, originalUrl string, shortUrl string) error {
	builder := sq.Insert(tableName).
		PlaceholderFormat(sq.Dollar).
		Columns("original_url, short_url").
		Values(originalUrl, shortUrl)

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	_, err = r.pool.Exec(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) GetOriginalUrl(ctx context.Context, shortlUrl string) (string, error) {
	builder := sq.Select("original_url").
		PlaceholderFormat(sq.Dollar).
		From(tableName).
		Where(sq.Eq{"short_url": shortlUrl}).
		Limit(1)

	query, args, err := builder.ToSql()
	if err != nil {
		return "", err
	}

	var originalUrl string
	err = r.pool.QueryRow(ctx, query, args...).Scan(&originalUrl)
	if err != nil {
		return "", err
	}

	return originalUrl, nil
}
