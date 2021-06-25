package database

import (
	"eCommerce/models"
	"strconv"
)

var Products []models.Product
var Orders []models.Order
var Users []models.User
var freeId = 0

func GetId() string {
	freeId++
	return strconv.Itoa(freeId)
}
//MySQL, interface