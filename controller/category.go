package controller

import (
	"fmt"
	"sort"

	"github.com/google/uuid"
	"main.go/models"
)

func (c Controller) CreateCategory() {
	category := getCategoryInfo()
	id, err := c.Store.CategoryStorage.Insert(category)
	if err != nil {
		fmt.Println("error inserting category ")
		return
	}
	fmt.Println("id ", id)

}
func (c Controller) GetListCategory() {
	categories, err := c.Store.CategoryStorage.GetList()
	if err != nil {
		fmt.Println("error list category ", err.Error())
		return
	}
	fmt.Println(categories)

}
func (c Controller) UpdateCategory() {
	cat := getCategoryInfo()
	err := c.Store.CategoryStorage.UpdateCategory(cat)
	if err != nil {
		fmt.Println("error updating category ", err.Error())
		return
	}
	if cat.Id.String() != "" {
		fmt.Println("UPDATED")
	} else {
		fmt.Println("CREATED")
	}

}
func (c Controller) DeleteCategory() {
	idstr := ""
	fmt.Print("enter id ")
	fmt.Scan(&idstr)
	id, err := uuid.Parse(idstr)
	if err != nil {
		fmt.Println("error while parsing 50 ", err.Error())
		return
	}
	err = c.Store.CategoryStorage.DeleteCategory(id)
	if err != nil {
		fmt.Println("error while deleting ", err.Error())
		return
	}
	fmt.Println("deleted category with this id :",idstr)
}

/////////////////////////////////////////////////

type Sort struct {
	Name  string
	Price int
}

func (c Controller) GetListCategoryAndTotalPrice() {
	categories, err := c.Store.CategoryStorage.GetList()
	if err != nil {
		fmt.Println("error list category ", err.Error())
		return
	}

	products, err := c.Store.ProductStorage.Getlist()
	if err != nil {
		fmt.Println("error getting list product ", err.Error())
		return
	}
	categoryTotalPrices := make(map[string]int)

	for _, category := range categories {
		totalCategoryPrice := 0
		for _, product := range products {
			if product.CtegoryId == category.Id {
				totalCategoryPrice += product.Price
			}
		}
		fmt.Println(category.Name, totalCategoryPrice)
		fmt.Println("----------------")
		categoryTotalPrices[category.Name] = totalCategoryPrice
	}
	sort1 := []Sort{}
	for i, v := range categoryTotalPrices {
		sort1 = append(sort1, Sort{i, v})
	}
	sort.Slice(sort1, func(i, j int) bool {
		return sort1[i].Price > sort1[i].Price

	})
	for _, price := range sort1 {
		fmt.Println(price.Name, price.Price)
	}
	fmt.Println(sort1[:1])

}

////////////////////////////////////////////////////////

func getCategoryInfo() models.Category {
	var (
		idstr string
		name  string
		cmd   int
	)
a:
	fmt.Println("enter command : 1.create 2. update ")
	fmt.Scan(&cmd)
	if cmd == 2 {
		fmt.Println("enter id to update: ")
		fmt.Scan(&idstr)
		fmt.Println("enter new name : ")
		fmt.Scan(&name)
	} else if cmd == 1 {

		fmt.Println("enter name : ")
	} else {
		fmt.Println("not found ")
		goto a
	}

	if idstr != "" {
		return models.Category{
			Id:   uuid.MustParse(idstr),
			Name: name,
		}
	}
	return models.Category{
		Name: name,
	}

}
