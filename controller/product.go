package controller

import (
	"fmt"

	"github.com/google/uuid"
	_ "github.com/google/uuid"
	"main.go/models"
)

func (c Controller) Createproduct() {
	product := getProductInfo()
	id, err := c.Store.ProductStorage.Insert(product)
	if err != nil {
		fmt.Println("error creating products ", err.Error())
		return
	}
	fmt.Println("id:", id)
}

func (c Controller) Getlistproduct() {
	products, err := c.Store.ProductStorage.Getlist()
	if err != nil {
		fmt.Println("error getting list product ", err.Error())
		return
	}
	fmt.Println(products)

}
func (c Controller) DeleteProduct() {
	idstr := ""
	fmt.Print("enter id : ")
	fmt.Scan(&idstr)
	id, err := uuid.Parse(idstr)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	err = c.Store.ProductStorage.DeleteProduct(id)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("deleted product with this id ", idstr)

}

func (c Controller) Updateproduct() {
	prod := getProductInfo()
	err := c.Store.ProductStorage.UpdateProduct(prod)
	if err != nil {
		fmt.Println("error in updating product ", err.Error())
		return
	}
	if prod.Id.String() != "" {
		fmt.Println("UPDATED")
	} else {
		fmt.Println("CREATED")
	}

}

func getProductInfo() models.Product {
	var (
		name       string
		categoryid string
		price, cmd int
		idstr      string
	)

	fmt.Println("enter cmd : 1.Create 2.Update ")
	fmt.Scan(&cmd)

a:
	if cmd == 2 {
		fmt.Print("enter id to update : ")
		fmt.Scan(&idstr)
		fmt.Print("enter new name of product : ")
		fmt.Scan(&name)

	} else if cmd == 1 {
		fmt.Print("enter name of product : ")
		fmt.Scan(&name)
		fmt.Print("enter category id ")
		fmt.Scan(&categoryid)
		fmt.Println("enter price : ")
		fmt.Scan(&price)

	} else {

		fmt.Println("error cmd ")
		goto a
	}

	if categoryid != "" {
		return models.Product{
			CtegoryId: uuid.MustParse(categoryid),
		}

	}
	if idstr != "" {
		return models.Product{
			Id:   uuid.MustParse(idstr),
			Name: name,
		}
	}

	return models.Product{
		Name:  name,
		Price: price,
	}
}
