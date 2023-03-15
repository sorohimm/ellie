package handles

import "context"

type CategoryHandle struct {
}

func (o *CategoryHandle) handleGetCategories() {}

func (o *CategoryHandle) handleGetSubcategories() {}

func (o *CategoryHandle) ApplyRotes(ctx context.Context) {}
