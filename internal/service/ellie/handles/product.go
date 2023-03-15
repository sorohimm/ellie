package handles

import "context"

type ProductHandle struct {
}

func (o *ProductHandle) handleGetProduct() {}

func (o *ProductHandle) handleGetProducts() {}

func (o *ProductHandle) handleGetRecommendedProducts() {}

func (o *ProductHandle) handleGetProductImage() {}

func (o *ProductHandle) ApplyRotes(ctx context.Context) {}
