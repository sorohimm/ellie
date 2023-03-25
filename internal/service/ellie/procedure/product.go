package procedure

import (
	"context"

	source "github.com/sorohimm/ellie/internal/storage/postgres/product"
)

type ProductProcedure struct {
	source *source.ProductsRequester
}

func (o *ProductProcedure) GetProduct(ctx context.Context, r *GetProductRequest) (*Product, error) {
	p, err := o.source.GetProduct(ctx, r.ID)
	if err != nil {
		return nil, err
	}

	product := Product{
		ID:          p.ID,
		Name:        p.Name,
		CategoryID:  p.CategoryID,
		ImageURL:    p.ImageURL,
		Price:       p.Price,
		Description: p.Description,
		IsStock:     p.IsStock,
	}

	return &product, nil
}

func (o *ProductProcedure) GetProducts() {

}

func (o *ProductProcedure) GetRecommendedProducts() {

}

func (o *ProductProcedure) GetProductImage() {

}
