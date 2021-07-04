package database

import (
	"database/sql"
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

func GetSpecificProduct(id int) models.Product {
	result, err := Db.Query("SELECT * FROM product WHERE id = ?", id)
	if err != nil {
		log.Fatalf("Cannot get product - %s", err.Error())
	}
	var product models.Product
	for result.Next() {
		err = result.Scan(&product.ProductId, &product.Name, &product.Price, &product.Category)
		if err != nil {
			log.Fatalf("Error while scanning data rows - %s", err.Error())
		}
	}
	return product
}

func DeleteProduct(id int) models.Product {
	result, err := Db.Query("SELECT * FROM product WHERE id = ?", id)
	if err != nil {
		log.Fatalf("Cannot get product - %s", err.Error())
	}
	var product models.Product
	for result.Next() {
		err = result.Scan(&product.ProductId, &product.Name, &product.Price, &product.Category)
		if err != nil {
			log.Fatalf("Error while scanning data rows - %s", err.Error())
		}
	}
	_, err = Db.Exec("DELETE FROM product WHERE id = ?", id)
	if err != nil {
		log.Fatalf("Cannot delete product - %s", err.Error())
	}
	return product
}

func AddProduct(product models.Product) models.Product {
	_, err := Db.Exec("INSERT INTO product (name, price, category) VALUES (?, ?, ?)", product.Name, product.Price, product.Category)
	if err != nil {
		log.Fatalf("Cannot add product - %s", err.Error())
	}
	var result *sql.Rows
	result, err = Db.Query("SELECT MAX(id) FROM product")
	if err != nil {
		log.Fatalf("Error retrieving ID for product - %s", err.Error())
	}
	result.Scan(&product.ProductId)
	return product
}

func UpdateProduct(id int, product models.Product) models.Product {
	_, err := Db.Exec("UPDATE product SET name = ?, price = ?, category = ? WHERE id = ?", product.Name, product.Price, product.Category, id)
	if err != nil {
		log.Fatalf("Cannot update product - %s", err.Error())
	}
	return product
}
