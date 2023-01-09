package Models

type Product struct {
	Id              uint    `json:"id"`
	Sku             string  `json:"sku"`
	Name            string  `json:"name"`
	Brand           string  `json:"brand"`
	Size            float64 `json:"size"`
	Price           float64 `json:"price"`
	Principal_image string  `json:"pincpmage"`
	Other_image     string  `json:"othimage"`
}

func (b *Product) TableName() string {
	return "product"
}
