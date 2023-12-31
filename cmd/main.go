package main

import (
	"log"

	_ "github.com/lib/pq"
	"main.go/config"
	"main.go/controller"
	"main.go/storage/postgres"
)
//test1 databaseda
func main() {
	cfg := config.Load()
	store, err := postgres.New(cfg)
	if err != nil {
		log.Fatalln("Error connecting", err.Error())
		return
	}
	defer store.DB.Close()

	con := controller.New(store)
	// con.GetListCategory()
	// con.Getlistproduct()
	// con.GetListCategoryAndTotalPrice()
	// con.UpdateCategory()
	// con.Updateproduct()
	// con.Createproduct()
	con.DeleteCategory()

}
