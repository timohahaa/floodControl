package floodcontroller

import (
	"context"
	"errors"

	"github.com/timohahaa/floodControl/storage"
)

type CustomFloodController struct {
	// время в секундах, за которое проверяется количество недавних запросов, то есть N
	checkInterval int
	// максимально допустимое количество запросов за последние N секунд, то есть K
	maxAllowedRequests int
	// общее хранилище данных
	storage storage.Storage
}

func NewFloodController(N, K int, storage storage.Storage) *CustomFloodController {
	return &CustomFloodController{
		checkInterval:      N,
		maxAllowedRequests: K,
		storage:            storage,
	}
}

func (c *CustomFloodController) Check(ctx context.Context, userID int64) (bool, error) {
	lastReqCount, err := c.storage.ReqInLastN(ctx, userID, c.checkInterval)
	if err != nil {
		return false, err
	}

	if lastReqCount > c.maxAllowedRequests {
		// не сохраняем при этом новый запрос - так как все равно привышен лимит запросов
		return false, nil
	}

	// можно сделать еще запрос, поэтому сохраняем запрос и возвращаем true
	// TODO: как лучше всего обработать эту ошибку?
	err = c.addUserReq(ctx, userID)
	return true, err

}

func (c *CustomFloodController) addUserReq(ctx context.Context, userID int64) error {
	err := c.storage.AddUserReq(ctx, userID)
	if err != nil {
		return err
	}

	return nil
}
