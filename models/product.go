package models

type ProductPrimaryKey struct {
	Id string `json:"id"`
}

type CreateProduct struct {
	Name       string  `json:"name"`
	Barcode    string  `json:"barcode"`
	Price      float64 `json:"price"`
	ImageUrl   string  `json:"image_url"`
	CategoryId string  `json:"category_id"`
}

type Product struct {
	Id         string      `json:"id"`
	Name       string      `json:"name"`
	Barcode    string      `json:"barcode"`
	Price      float64     `json:"price"`
	ImageUrl   string      `json:"image_url"`
	CategoryId string      `json:"category_id"`
	Category   interface{} `json:"category"`
	CreatedAt  string      `json:"created_at"`
	UpdatedAt  string      `json:"updated_at"`
}

type UpdateProduct struct {
	Id         string  `json:"id"`
	Name       string  `json:"name"`
	Barcode    string  `json:"barcode"`
	Price      float64 `json:"price"`
	ImageUrl   string  `json:"image_url"`
	CategoryId string  `json:"category_id"`
}

type GetListProductRequest struct {
	Offset int64  `json:"offset"`
	Limit  int64  `json:"limit"`
	Search string `json:"search"`
}

type GetListProductResponse struct {
	Count    int        `json:"count"`
	Products []*Product `json:"products"`
}
