package repository

import (
	"Taxi_service/db"
	"Taxi_service/logger"
	"Taxi_service/models"
)

func InsertOrder(order models.Order) error {
	err := db.GetconnectDB().Create(&order).Error
	if err != nil {
		logger.Error.Printf("[repository.InsertOrder]error in added order %s\n", err.Error())
	}
	return nil
}
func EditOrder(distance int, driverphone string, userid int, id int) error {
	var order models.Order
	err := db.GetconnectDB().Raw("UPDATE orders o SET distance=?,driver_phone=?,driver_id=? WHERE id=?", distance, driverphone, userid, id).Scan(&order).Error
	if err != nil {
		logger.Error.Printf("[repository.EditOrder]error in update order %s\n", err.Error())
	}
	return nil
}

func SoftDeleteOrder(isdeleted bool, id int) error {
	err := db.GetconnectDB().Model(&models.Order{}).Where("id=?", id).Update("is_deleted", isdeleted).Error
	if err != nil {
		logger.Error.Printf("[repository.SofDeleteOrder]error in deleted order %s\n", err.Error())
	}
	return nil
}

func Report(isresp, isdeletedr, isblocked, isdeletedu bool, price int) (order []models.Reports, err error) {
	err = db.GetconnectDB().Raw("Select o.from, o.into, o.is_response, SUM(DISTINCT o.distance) as distance ,SUM(DISTINCT o.pricekm) as pricekm, SUM(o.all_price/2) as all_price, COUNT(DISTINCT CASE WHEN u.role='user' THEN u.id END) as client_id, COUNT(DISTINCT CASE WHEN u.role='driver' THEN u.id END) as driver_id FROM orders o, users u Where o.client_id=u.id OR o.driver_id=u.id  AND  o.is_response=? AND o.is_deleted=?  AND u.is_blocked=? AND u.is_deleted=? AND all_price<=? GROUP BY  o.from,o.into,o.is_response ORDER BY all_price DESC", isresp, isdeletedr, isblocked, isdeletedu, price).Scan(&order).Error
	if err != nil {
		logger.Error.Printf("[repository.Report]error in report %s\n", err.Error())
		return order, err
	}
	return order, nil
}
func GetAllOrders(isdeleted, isresp bool, price int, uid uint) (order []models.GetOrder, err error) {
	err = db.GetconnectDB().Raw("SELECT DISTINCT o.from, o.into,  o.distance, o.pricekm, o.all_price, o.is_response FROM orders o WHERE o.is_deleted=? AND o.is_response=? AND o.all_price<=? AND (o.driver_id=? OR o.client_id=?)", isdeleted, isresp, price, uid, uid).Scan(&order).Error
	if err != nil {
		logger.Error.Printf("[repository.GetAllOrders]error in getting all order %s\n", err.Error())
		return order, err
	}
	return order, nil
}
func GetAllOrdersByID(isdeleted bool, id, uid uint) (order []models.GetOrder, err error) {
	err = db.GetconnectDB().Raw("SELECT o.from, o.into, o.distance, o.pricekm, o.all_price, o.is_response FROM orders o WHERE o.is_deleted=?  AND o.id=? AND  (o.driver_id=? OR o.client_id=?)", isdeleted, id, uid, uid).Scan(&order).Error
	if err != nil {
		logger.Error.Printf("[repository.GetAllOrdersByID]error in getting all order by id %s\n", err.Error())
		return order, err
	}
	return order, nil
}

func CheckOrdersAsResponse(isresp bool, id, uid uint) error {
	var order models.Order
	err := db.GetconnectDB().Raw("Update orders o SET is_response=? WHERE o.id=? AND (o.client_id=? OR o.driver_id=?)", isresp, id, uid, uid).Scan(&order).Error
	if err != nil {
		logger.Error.Printf("[repository.CheckRoutesAsResponse]error in checked order %s\n", err.Error())

	}
	return nil
}
