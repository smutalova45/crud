package postgres

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"main.go/models"
)

type categoryRepo struct {
	DB *sql.DB
}

func NewCategoryRepo(db *sql.DB) categoryRepo {
	return categoryRepo{
		DB: db,
	}
}
func (c categoryRepo) Insert(category models.Category) (string, error) {
	id := uuid.New()
	category.CreatedAt = time.Now()
	if _, err := c.DB.Exec(`insert into category values($1 ,$2, $3)`, id, category.Name, category.CreatedAt); err != nil {
		return "", err
	}
	return id.String(), nil
}

func (c categoryRepo) GetList() ([]models.Category, error) {
	rows, err := c.DB.Query(`select * from category `)
	if err != nil {
		return nil, err
	}
	c1 := []models.Category{}
	for rows.Next() {
		ca := models.Category{}
		if err = rows.Scan(&ca.Id, &ca.Name, &ca.CreatedAt, &ca.UpdatedAt); err != nil {
			return nil, err
		}
		c1 = append(c1, ca)

	}
	return c1, nil

}

func (c categoryRepo) UpdateCategory(cat models.Category) error {
	cat.UpdatedAt = time.Now()
	_, err := c.DB.Exec(`update category set namec= $1 , updated_at = $2 where id = $3`, cat.Name, cat.UpdatedAt, cat.Id)
	if err != nil {
		return err
	}
	return nil
}

func (c categoryRepo) DeleteCategory(id uuid.UUID) error {

	if _, err := c.DB.Exec("delete from products where category_id = $1", id); err != nil {
		return err
	}

	//then delete category itself
	if _, err := c.DB.Exec(`delete from category where id = $1 `, id); err != nil {
		return err
	}
	return nil

}
