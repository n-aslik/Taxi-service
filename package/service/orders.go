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
	for i := 10; i <= order.Distance*order.StartPrice; i++ {
		order.AllPrice += i
	}
	err := repository.EditOrder(order.Distance, order.DriverPhone, order.DriverID, id)
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
func CheckOrderasResponse(isresp bool, id, uid uint) error {
	err := repository.CheckOrdersAsResponse(isresp, id, uid)
	if err != nil {
		fmt.Println(err)
	}
	return nil
}

func Printreport(isresp, isdeletedr, isblocked, isdeletedu bool, price int) (order []models.Reports, err error) {
	order, err = repository.Report(isresp, isdeletedr, isblocked, isdeletedu, price)
	if err != nil {

		return order, err
	}
	return order, nil
}
func PrintAllOrders(isdeleted, isresp bool, price int, uid uint) (order []models.GetOrder, err error) {
	order, err = repository.GetAllOrders(isdeleted, isresp, price, uid)
	if err != nil {

		return order, err
	}
	return order, nil
}

func PrintAllOrderByID(isdeleted bool, id, uid uint) (route []models.GetOrder, err error) {
	route, err = repository.GetAllOrdersByID(isdeleted, id, uid)
	if err != nil {

		return route, err
	}
	return route, nil
}
