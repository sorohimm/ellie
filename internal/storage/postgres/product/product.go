package product

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
)

type ProductsRequester struct {
	pool *pgxpool.Pool
}

func (o *ProductsRequester) GetProduct(ctx context.Context, id string) (*Product, error) {
	return o.getProduct(id)
}

func (o *ProductsRequester) getProduct(id string) (*Product, error) {
	// Prepare the SQL query
	query := `SELECT id, name, category_id price, description, image_url 
			  FROM products 
			  WHERE id = $1`

	// Execute the query
	row := o.pool.QueryRow(context.Background(), query, id)

	var p Product
	if err := row.Scan(&p.ID, &p.Name, &p.Price, &p.Description, &p.ImageURL); err != nil {
		return nil, err
	}

	return &p, nil
}
