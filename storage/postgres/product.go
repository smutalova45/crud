package postgres

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"main.go/models"
)

type productRepo struct {
	DB *sql.DB
}

func NewproductRepo(db *sql.DB) productRepo {
	return productRepo{
		DB: db,
	}
}

func (p productRepo) Insert(product models.Product) (string, error) {
	id := uuid.New()
	product.CreatedAt=time.Now()
	if _, err := p.DB.Exec(`insert into products (id, name_, category_id,created_at, updated_at, price) values ($1,$2,$3,$4,$5,$6)`, id, product.Name, product.CtegoryId, product.CreatedAt, product.UpdatedAt, product.Price); err != nil {
		return "", err
	}
	return id.String(), nil

}

func (p productRepo) Getlist() ([]models.Product, error) {

	pro := []models.Product{}
	rows, err := p.DB.Query(` select id, name_, price, category_id, created_at, updated_at from products`)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		p := models.Product{}
		if err = rows.Scan(&p.Id, &p.Name, &p.Price, &p.CtegoryId, &p.CreatedAt, &p.UpdatedAt); err != nil {
			return nil, err
		}
		pro = append(pro, p)

	}
	return pro, nil

}

func (p productRepo) UpdateProduct(pro models.Product) error {
	pro.UpdatedAt = time.Now()
	if _, err := p.DB.Exec(`update products set name_ = $1 , updated_at = $2 where id = $3 `, pro.Name, pro.UpdatedAt, pro.Id); err != nil {
		return err
	}
	return nil
}

func (p productRepo) DeleteProduct(id uuid.UUID) error {
	if _, err := p.DB.Exec(`delete from products where id = $1`, id); err != nil {
		return err
	}
	return nil

}
