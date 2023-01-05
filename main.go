package main

import (
	"fmt"
	"measutosh/apcommerce/api"

	"github.com/anthdm/weavebox"
)

func main() {
	app := weavebox.New()
	productHandler := &api.ProductHandler{}
	app.Get("/product", productHandler.HandleGetProduct)
	app.Serve(3001)
	fmt.Println("Working fine..")
}  
