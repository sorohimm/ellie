package handles

import (
	"context"

	"github.com/gin-gonic/gin"

	"github.com/sorohimm/ellie/internal/service/ellie/controller"
)

type ProductHandle struct {
	ctrl *controller.ProductController
}

func (o *ProductHandle) handleGetProduct(ctx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		o.ctrl.GetProduct(ctx, c)
	}
}

func (o *ProductHandle) handleGetProducts() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func (o *ProductHandle) handleGetRecommendedProducts() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func (o *ProductHandle) handleGetProductImage() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func (o *ProductHandle) handleCreateProduct(ctx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func (o *ProductHandle) ApplyRotes(ctx context.Context, r gin.IRouter) {
	r.GET("/products", o.handleGetProducts())
	r.GET("/products/:productId", o.handleGetProduct(ctx))
	r.POST("/products", o.handleCreateProduct(ctx))
}
