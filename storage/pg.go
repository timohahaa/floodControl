package storage

import (
	"context"
	"time"

	"github.com/timohahaa/postgres"
)

type storage struct {
	db *postgres.Postgres
}

func NewStorage(pg *postgres.Postgres) *storage {
	return &storage{
		db: pg,
	}
}

func (s *storage) AddUserReq(ctx context.Context, userID int64) error {
	sql, args, _ := s.db.Builder.
		Insert("requests").
		Columns("user_id, req_timestamp").
		Values(userID, time.Now()).
		ToSql()

	_, err := s.db.ConnPool.Exec(ctx, sql, args)
	if err != nil {
		return ErrFailedToSaveRequest
	}

	return nil
}

func (s *storage) ReqInLastN(ctx context.Context, userID int64, N int) (int, error) {
	// если такой пользователь еще не делал запросы, вернем ноль, все логично
	sql, args, _ := s.db.Builder.
		Select("COUNT(*)").
		From("requests").
		Where("user_id = ? AND req_timestamp > (NOW() - interval '? seconds')", userID, N).
		ToSql()

	var count int
	err := s.db.ConnPool.QueryRow(ctx, sql, args).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}
