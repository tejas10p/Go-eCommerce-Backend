package workers

import (
	"eCommerce/database"
	"eCommerce/kafka"
)

func InitiatePurchase(id string) (int, error) {
	p := kafka.NewProducer()
	kafka.SendMessage(p, id, "purchases")
	result, err := database.BuyOrder()
	kafka.CloseProducer(p)
	return result, err
}
