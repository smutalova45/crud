package postgres

import (
	"database/sql"
	"fmt"

	"main.go/config"
)

type Storage struct {
	DB              *sql.DB
	ProductStorage  productRepo
	CategoryStorage categoryRepo
}

func New(cfg config.Config) (Storage, error) {
	url := fmt.Sprintf(`host=%s port=%s user=%s password=%s database=%s sslmode=disable`, cfg.PostgresHost, cfg.PostgresPort, cfg.PostgresUser, cfg.PostgresPassword, cfg.PostgresDB)

	db, err := sql.Open("postgres", url)
	if err != nil {
		return Storage{}, err
	}

	productRepo := NewproductRepo(db)
	categoryRepo := NewCategoryRepo(db)

	return Storage{
		DB:              db,
		ProductStorage:  productRepo,
		CategoryStorage: categoryRepo,
	}, nil
}
