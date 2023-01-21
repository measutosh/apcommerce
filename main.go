package main

import (
	"fmt"
	"measutosh/apcommerce/api"

	"github.com/anthdm/weavebox"
)

func main() {
	app := weavebox.New()
	adminRoute := app.Box("/admin")

	productHandler := &api.ProductHandler{}
	adminRoute.Get("/product", productHandler.HandleGetProduct)
	adminRoute.Get("/product", productHandler.HandlePostProduct)
	app.Serve(3001)
	fmt.Println("Server started at :3001")
}
