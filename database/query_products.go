package database

import (
	"eCommerce/models"
	"log"
)

func GetProductList() []models.Product {
	result, err := Db.Query("SELECT * FROM product")
	if err != nil {
		log.Fatalf("Cannot get products - %s", err.Error())
	}
	var products []models.Product
	for result.Next() {
		var currentProduct models.Product
		err = result.Scan(&currentProduct.ProductId, &currentProduct.Name, &currentProduct.Price, &currentProduct.Category)
		if err != nil {
			log.Fatalf("Error while scanning data rows - %s", err.Error())
		}
		products = append(products, currentProduct)
	}
	return products
}
