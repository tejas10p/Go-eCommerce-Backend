package models

type Product struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Price int `json:"price"`
	Quantity int `json:"quantity"`
	Category string `json:"category"`
}

func (product Product) Validate() bool {
	for _, c := range product.Name {
		if !(c >= 'A' && c <= 'Z') && !(c >= '1' && c <= '9') {
			return false
		}
	}
	return true
}