package vk

import (
	"context"
	"fmt"

	fc "github.com/timohahaa/floodControl/floodController"
	"github.com/timohahaa/floodControl/storage"
	"github.com/timohahaa/postgres"
)

// FloodControl интерфейс, который нужно реализовать.
// Рекомендуем создать директорию-пакет, в которой будет находиться реализация.
type FloodControl interface {
	// Check возвращает false если достигнут лимит максимально разрешенного
	// кол-ва запросов согласно заданным правилам флуд контроля.
	Check(ctx context.Context, userID int64) (bool, error)
}

func main() {
	// пример использования
	pg, err := postgres.New("postgres://myuser:mypassword@localhost:5432/db_name", postgres.MaxConnPoolSize(5))
	if err != nil {
		// обрабатываем ошибку
	}

	storage := storage.NewStorage(pg)

	myFloodController := fc.NewFloodController(10, 100, storage)

	allowed, err := myFloodController.Check(context.TODO(), 12345)
	if err != nil {
		// обрабатываем ошибку
	}
	// как нибудь дальше используем полученную информацию
	fmt.Println(allowed)
}
