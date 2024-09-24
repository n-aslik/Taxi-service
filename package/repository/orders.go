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
func EditOrder(cphone, address, from, into string, distance, startprice, allprice int, dphone string, cid, did, id int) error {
	err := db.GetconnectDB().Omit("client_phone", "client_address", "from", "into", "client_id").Save(&models.Order{ID: id, ClientPhone: cphone, ClientAddress: address, From: from, Into: into, Distance: distance, StartPrice: startprice, AllPrice: allprice, DriverPhone: dphone, ClientID: cid, DriverID: did}).Error
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
	err = db.GetconnectDB().Raw("Select o.from, o.into, o.is_response, SUM(DISTINCT o.distance) as distance ,SUM(DISTINCT o.start_price) as start_price, SUM(o.all_price/2) as all_price, COUNT(DISTINCT CASE WHEN u.role='user' THEN u.id END) as client_id, COUNT(DISTINCT CASE WHEN u.role='driver' THEN u.id END) as driver_id FROM orders o, users u Where o.client_id=u.id OR o.driver_id=u.id  AND  o.is_response=? AND o.is_deleted=?  AND u.is_blocked=? AND u.is_deleted=? AND all_price<=? GROUP BY  o.from,o.into,o.is_response ORDER BY all_price DESC", isresp, isdeletedr, isblocked, isdeletedu, price).Scan(&order).Error
	if err != nil {
		logger.Error.Printf("[repository.Report]error in report %s\n", err.Error())
		return order, err
	}
	return order, nil
}
func GetAllOrders(isdeleted, isresp bool, price int, uid uint) (order []models.GetOrder, err error) {
	err = db.GetconnectDB().Raw("SELECT DISTINCT o.client_phone, o.from, o.into,  o.distance, o.start_price, o.all_price, o.is_response FROM orders o WHERE o.is_deleted=? AND o.is_response=? AND o.all_price<=? AND (o.driver_id=? OR o.client_id=?)", isdeleted, isresp, price, uid, uid).Scan(&order).Error
	if err != nil {
		logger.Error.Printf("[repository.GetAllOrders]error in getting all order %s\n", err.Error())
		return order, err
	}
	return order, nil
}
func GetAllOrdersByID(isdeleted bool, id uint) (order []models.GetOrder, err error) {
	err = db.GetconnectDB().Raw("SELECT o.client_phone, o.from, o.into, o.distance, o.start_price, o.all_price, o.is_response FROM orders o WHERE o.is_deleted=?  AND o.id=?", isdeleted, id).Scan(&order).Error
	if err != nil {
		logger.Error.Printf("[repository.GetAllOrdersByID]error in getting all order by id %s\n", err.Error())
		return order, err
	}
	return order, nil
}

func CheckOrdersAsResponse(isresp bool, cid, did, id int) error {
	err := db.GetconnectDB().Model(&models.Order{}).Where("id=? AND (client_id=? OR driver_id=?)", id, cid, did).Select("is_response").Updates(models.Order{IsResponse: isresp}).Error
	if err != nil {
		logger.Error.Printf("[repository.CheckRoutesAsResponse]error in checked order %s\n", err.Error())

	}
	return nil
}
