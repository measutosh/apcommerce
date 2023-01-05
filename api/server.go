package api

import (
	"measutosh/apcommerce/types"
	"net/http"

	"github.com/anthdm/weavebox"
)

type ProductHandler struct {

}

func (h *ProductHandler) HandleGetProduct(c *weavebox.Context) error{
	return c.JSON(http.StatusOK, &types.Product{SKU: "SHOW-1111"})
}