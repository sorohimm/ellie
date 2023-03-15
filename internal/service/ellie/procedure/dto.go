package procedure

type GetProduct struct {
	ID string
}

type GetProductsRequest struct {
	Limit    int
	Offset   int
	Category string
	Sort     string
}

type Product struct {
	ID     string
	Name   string
	Status string
	Model  string
	Price  string
}
