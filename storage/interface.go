package storage

import (
	"context"
)

type Storage interface {
	// сохраняет запрос пользователя
	AddUserReq(ctx context.Context, userID int64) error
	// достает количество запросов, которые были сделаны за последние N секунд
	ReqInLastN(ctx context.Context, userID int64, N int) (int, error)
}

var _ Storage = &storage{}
