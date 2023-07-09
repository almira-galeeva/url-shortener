package db

import (
	"context"
	"fmt"

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

func (r *repository) CreateUrl(ctx context.Context, originalUrl string, shortUrl string) (bool, error) {
	// check if original url already is in db
	builderCheckOriginal := sq.Select("COUNT(*) AS cnt").
		PlaceholderFormat(sq.Dollar).
		From(tableName).
		Where(sq.Eq{"original_url": originalUrl}).
		Limit(1)

	query, args, err := builderCheckOriginal.ToSql()
	if err != nil {
		return false, err
	}

	var cnt int
	err = r.pool.QueryRow(ctx, query, args...).Scan(&cnt)
	if err != nil {
		return false, err

	}

	if cnt != 0 {
		return false, fmt.Errorf("Url %s already exists in db", originalUrl)
	}

	// check if generated url has to be regenerated
	builderCheckShort := sq.Select("COUNT(*) AS cnt").
		PlaceholderFormat(sq.Dollar).
		From(tableName).
		Where(sq.Eq{"short_url": shortUrl}).
		Limit(1)

	query, args, err = builderCheckShort.ToSql()
	if err != nil {
		return false, err
	}

	err = r.pool.QueryRow(ctx, query, args...).Scan(&cnt)
	if err != nil {
		return false, err

	}

	if cnt != 0 {
		return true, nil
	}

	// insert original_url, short_url in db
	builderInsert := sq.Insert(tableName).
		PlaceholderFormat(sq.Dollar).
		Columns("original_url, short_url").
		Values(originalUrl, shortUrl)

	query, args, err = builderInsert.ToSql()
	if err != nil {
		return false, err
	}

	_, err = r.pool.Exec(ctx, query, args...)
	if err != nil {
		return false, err
	}

	return false, nil
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
