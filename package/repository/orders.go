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
func EditOrders(dphone string, did, id int) error {
	var order models.Order
	err := db.GetconnectDB().Model(&order).Where("id=?", id).Updates(models.Order{DriverPhone: dphone, DriverID: did}).Error
	if err != nil {
		logger.Error.Printf("[repository.EditOrders]error in update order %s\n", err.Error())
	}
	return nil
}
func AddDistanceandTotal(distance, startprice, allprice, id int) error {
	var order models.Order
	err := db.GetconnectDB().Model(&order).Select("distance", "start_price", "all_price").Where("id=?", id).Updates(models.Order{Distance: distance, StartPrice: startprice, AllPrice: allprice}).Error
	if err != nil {
		logger.Error.Printf("[repository.AddDistance]error in add order distance %s\n", err.Error())
		return nil
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

func Report(isrespc, isrespd, isdeletedr, isblocked, isdeletedu bool, price int) (order []models.Reports, err error) {
	err = db.GetconnectDB().Raw("Select o.from, o.into, o.is_responsec,o.is_responsed,  SUM(DISTINCT o.distance) as distance ,SUM(DISTINCT o.start_price) as start_price, SUM(o.all_price/2) as all_price, COUNT(DISTINCT CASE WHEN u.role='user' THEN u.id END) as client_id, COUNT(DISTINCT CASE WHEN u.role='driver' THEN u.id END) as driver_id FROM orders o, users u Where o.client_id=u.id OR o.driver_id=u.id  AND  o.is_responsec=? AND o.is_responsed=? AND o.is_deleted=?  AND u.is_blocked=? AND u.is_deleted=? AND all_price<=? GROUP BY  o.from,o.into,o.is_responsec, o.is_responsed ORDER BY all_price DESC", isrespc, isrespd, isdeletedr, isblocked, isdeletedu, price).Scan(&order).Error
	if err != nil {
		logger.Error.Printf("[repository.Report]error in report %s\n", err.Error())
		return order, err
	}
	return order, nil
}
func GetAllOrders(isdeleted, isrespc, isrespd bool, price int, uid uint) (order []models.GetOrder, err error) {
	err = db.GetconnectDB().Raw("SELECT DISTINCT o.client_phone, o.from, o.into,  o.distance, o.start_price, o.all_price, o.driver_phone, o.is_responsec,o.is_responsed FROM orders o WHERE o.is_deleted=? AND o.is_responsec=? AND o.is_responsed=? AND o.all_price<=? AND (o.driver_id=? OR o.client_id=?)", isdeleted, isrespc, isrespd, price, uid, uid).Scan(&order).Error
	if err != nil {
		logger.Error.Printf("[repository.GetAllOrders]error in getting all order %s\n", err.Error())
		return order, err
	}
	return order, nil
}
func GetAllOrdersByID(isdeleted bool, id uint) (order []models.GetOrder, err error) {
	err = db.GetconnectDB().Raw("SELECT o.client_phone, o.from, o.into, o.distance, o.start_price, o.all_price,o.driver_phone, o.is_responsec,is_responsed FROM orders o WHERE o.is_deleted=?  AND o.id=?", isdeleted, id).Scan(&order).Error
	if err != nil {
		logger.Error.Printf("[repository.GetAllOrdersByID]error in getting all order by id %s\n", err.Error())
		return order, err
	}
	return order, nil
}

func CheckOrdersAsResponse(isrespc, isrespd bool, id int) error {
	err := db.GetconnectDB().Select("is_responsec", "is_responsed").Where("id=?", id).Updates(models.Order{IsResponsec: isrespc, IsResponsed: isrespd}).Error
	if err != nil {
		logger.Error.Printf("[repository.CheckRoutesAsResponse]error in checked order %s\n", err.Error())

	}
	return nil
}
