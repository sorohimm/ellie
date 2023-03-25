package controller

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/sorohimm/ellie/internal/service/ellie/procedure"
)

func NewProductController() *ProductController {
	return &ProductController{}
}

type ProductController struct {
	proc *procedure.ProductProcedure
}

func (o *ProductController) GetProducts(c *gin.Context) {

}

func (o *ProductController) GetProduct(ctx context.Context, c *gin.Context) {
	id := c.Param("productId")

	product, err := o.proc.GetProduct(ctx, &procedure.GetProductRequest{ID: id})
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, product)
}

func (o *ProductController) GetRecommendedProducts(c *gin.Context) {

}

func (o *ProductController) GetProductImage(c *gin.Context) {

}
