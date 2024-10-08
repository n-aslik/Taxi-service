package service

import (
	"Taxi_service/models"
	"Taxi_service/package/repository"
	"fmt"
)

func AddOrder(order models.Order) error {
	err := repository.InsertOrder(order)
	if err != nil {
		fmt.Println(err)
	}
	return nil
}
func UpdateOrder(order models.Order, id int) error {
	err := repository.EditOrders(order.DriverPhone, order.DriverID, id)
	if err != nil {
		fmt.Println(err)

	}
	return nil
}

func AddOrdersDistanceandTotal(order models.Order, id int) error {
	order.StartPrice = 10
	for i := 1; i <= order.StartPrice+order.Distance; i++ {
		order.AllPrice = i
	}
	err := repository.AddDistanceandTotal(order.Distance, order.StartPrice, order.AllPrice, id)
	if err != nil {
		fmt.Println(err)
	}
	return nil
}
func DeleteOrder(isdeleted bool, id int) error {
	err := repository.SoftDeleteOrder(isdeleted, id)
	if err != nil {
		fmt.Println(err)
	}
	return nil
}
func CheckOrderasResponse(order models.Order, id int) error {
	err := repository.CheckOrdersAsResponse(order.IsResponsec, order.IsResponsed, id)
	if err != nil {
		fmt.Println(err)
	}
	return nil
}

func Printreport(isrespc, isrespd, isdeletedr, isblocked, isdeletedu bool, price int) (order []models.Reports, err error) {
	order, err = repository.Report(isrespc, isrespd, isdeletedr, isblocked, isdeletedu, price)
	if err != nil {

		return order, err
	}
	return order, nil
}
func PrintAllOrders(isdeleted, isrespc, isrespd bool, price int, uid uint) (order []models.GetOrder, err error) {
	order, err = repository.GetAllOrders(isdeleted, isrespc, isrespd, price, uid)
	if err != nil {

		return order, err
	}
	return order, nil
}

func PrintAllOrderByID(isdeleted bool, id uint) (route []models.GetOrder, err error) {
	route, err = repository.GetAllOrdersByID(isdeleted, id)
	if err != nil {

		return route, err
	}
	return route, nil
}
