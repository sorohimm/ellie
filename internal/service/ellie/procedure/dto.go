package procedure

type GetProductRequest struct {
	ID string
}

type GetProductsRequest struct {
	Limit    int
	Offset   int
	Category string
	Sort     string
}

type CreateProductRequest struct {
	Name        string  `json:"name,omitempty"`
	CategoryID  string  `json:"categoryID,omitempty"`
	Price       float32 `json:"price,omitempty"`
	Description string  `json:"description,omitempty"`
}

type Product struct {
	ID          string  `json:"ID,omitempty"`
	Name        string  `json:"name,omitempty"`
	CategoryID  string  `json:"categoryID,omitempty"`
	ImageURL    string  `json:"imageURL,omitempty"`
	Price       float32 `json:"price,omitempty"`
	Description string  `json:"description,omitempty"`
	IsStock     string  `json:"isStock,omitempty"`
}
