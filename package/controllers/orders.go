package controllers

import (
	"Taxi_service/errs"
	"Taxi_service/logger"
	"Taxi_service/models"
	"Taxi_service/package/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateOrder
// @Summary Create Order
// @Security AKA
// @Tags orders
// @Description create new order
// @ID create-order
// @Accept json
// @Produce json
// @Param input body models.Order true "new order info"
// @Success 200 {object} defaultResponse
// @Failure 400 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /api/orders [post]
func CreateOrder(c *gin.Context) {
	userID := c.GetUint(userIDCtx)
	urole := c.GetString(userRoleCtx)
	uphone := c.GetString(userPhoneCtx)
	uaddr := c.GetString(userAddressCtx)
	if urole == "" {
		HandleError(c, errs.ErrValidationFailed)
		return
	}
	if urole != "user" {
		HandleError(c, errs.ErrPermissionDenied)
		return
	}
	var newroute models.Order
	err := c.BindJSON(&newroute)
	if err != nil {
		HandleError(c, errs.ErrValidationFailed)
		return
	}

	newroute.ClientID = int(userID)
	newroute.ClientPhone = uphone
	newroute.ClientAddress = uaddr
	logger.Info.Printf("[controllers.AddRoute] add order is succesful")

	err = service.AddOrder(newroute)
	if err != nil {
		HandleError(c, errs.ErrRoutesNotFound)
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Succesful created"})

}

// Report
// @Summary Get Report
// @Security AKA
// @Tags report
// @Description get list of report
// @ID get-report
// @Produce json
// @Param q query string false "fill if you need search"
// @Param is_response query bool true "fill if you need search"
// @Param all_price query int true "fill if you need search"
// @Success 200 {array} models.Reports
// @Failure 400 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /api/report [get]
func Report(c *gin.Context) {
	userID := c.GetUint(userIDCtx)
	if userID == 0 {
		HandleError(c, errs.ErrRoutesNotFound)
		return
	}
	urole := c.GetString(userRoleCtx)
	if urole == "" {
		HandleError(c, errs.ErrValidationFailed)
		return
	}
	if urole != "admin" {
		HandleError(c, errs.ErrPermissionDenied)
		return
	}

	isRespStr := c.Query("is_response")
	isResp, err := strconv.ParseBool(isRespStr)
	if err != nil {
		HandleError(c, errs.ErrValidationFailed)
		return
	}
	priceStr := c.Query("all_price")
	price, err := strconv.Atoi(priceStr)
	if err != nil {
		HandleError(c, errs.ErrValidationFailed)
		return
	}
	report, err := service.Printreport(isResp, false, false, false, price)
	if err != nil {
		HandleError(c, errs.ErrRoutesNotFound)
		return
	}
	c.JSON(http.StatusOK, gin.H{"report": report})
}

// GetAllOrders
// @Summary Get All Orders
// @Security AKA
// @Tags orders
// @Description get list of all orders
// @ID get-all-orders
// @Produce json
// @Param q query string false "fill if you need search"
// @Param is_response query bool true "fill if you need search"
// @Param all_price query int true "fill if you need search"
// @Success 200 {array} models.GetOrder
// @Failure 400 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /api/orders [get]
func GetAllOrders(c *gin.Context) {
	userID := c.GetUint(userIDCtx)
	if userID == 0 {
		HandleError(c, errs.ErrRoutesNotFound)
		return
	}
	urole := c.GetString(userRoleCtx)
	if urole == "" {
		HandleError(c, errs.ErrValidationFailed)
		return
	}

	isRespStr := c.Query("is_response")
	isResp, err := strconv.ParseBool(isRespStr)
	if err != nil {
		HandleError(c, errs.ErrValidationFailed)
		return
	}
	priceStr := c.Query("all_price")
	price, err := strconv.Atoi(priceStr)
	if err != nil {
		HandleError(c, errs.ErrValidationFailed)
		return
	}

	routes, err := service.PrintAllOrders(false, isResp, price, uint(userID))
	if err != nil {
		HandleError(c, errs.ErrRoutesNotFound)
		return
	}
	c.JSON(http.StatusOK, gin.H{"orders": routes})
}

// GetOrdersByID
// @Summary Get Order By ID
// @Security AKA
// @Tags orders
// @Description get order by ID
// @ID get-order-by-id
// @Produce json
// @Param id path integer true "id of the route"
// @Success 200 {object} models.GetOrder
// @Failure 400 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /api/orders/{id} [get]
func GetAllOrdersByID(c *gin.Context) {
	userID := c.GetUint(userIDCtx)
	urole := c.GetString(userRoleCtx)
	if userID == 0 {
		HandleError(c, errs.ErrRecordNotFound)
		return
	}
	if urole == "" {
		HandleError(c, errs.ErrValidationFailed)
		return
	}
	rid, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.Error.Printf("[controllers.GetAllOrdersByID] invalid order_id path parameter: %s\n", c.Param("id"))
		HandleError(c, errs.ErrValidationFailed)
		return
	}
	route, err := service.PrintAllOrderByID(false, uint(rid))
	if err != nil {
		HandleError(c, errs.ErrRoutesNotFound)
		return
	}
	c.JSON(http.StatusOK, gin.H{"order": route})
}

// UpdateOrderByID
// @Summary Update Order
// @Security AKA
// @Tags orders
// @Description update existed order
// @ID update-order
// @Accept json
// @Produce json
// @Param id path integer true "id of the order"
// @Param input body models.EditOrder true "order update info"
// @Success 200 {object} defaultResponse
// @Failure 400 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /api/orders/{id} [put]
func UpdateOrderByID(c *gin.Context) {
	userID := c.GetUint(userIDCtx)
	if userID == 0 {
		HandleError(c, errs.ErrRecordNotFound)
		return
	}
	urole := c.GetString(userRoleCtx)
	if urole == "" {
		HandleError(c, errs.ErrValidationFailed)
		return
	}
	if urole != "admin" {
		HandleError(c, errs.ErrPermissionDenied)
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.Error.Printf("[controllers.UpdateOrderByID] invalid order_id path parameter: %s\n", c.Param("id"))
		HandleError(c, errs.ErrValidationFailed)
		return
	}

	var route models.Order
	err = c.BindJSON(&route)
	if err != nil {
		HandleError(c, errs.ErrValidationFailed)
		return
	}
	err = service.UpdateOrder(route, id)
	if err != nil {
		HandleError(c, errs.ErrRoutesNotFound)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Update is successful"})

}

// ChecksOrderasResponse
// @Summary Check order as response
// @Security AKA
// @Tags orders
// @Description  check as response existed order
// @ID check-order-as-response
// @Accept json
// @Produce json
// @Param id path integer true "id of the route"
// @Param input body models.Checkresponse true " check order as response info"
// @Success 200 {object} defaultResponse
// @Failure 400 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /api/orders/{id} [patch]
func ChecksOrderasResponse(c *gin.Context) {
	urole := c.GetString(userRoleCtx)
	if urole == "" {
		HandleError(c, errs.ErrValidationFailed)
		return
	}
	if urole != "user" && urole != "driver" {
		HandleError(c, errs.ErrPermissionDenied)
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.Error.Printf("[controllers.ChecksOrderasResponse] invalid order_id path parameter: %s\n", c.Param("id"))
		HandleError(c, errs.ErrValidationFailed)
		return
	}
	userID := c.GetUint(userIDCtx)
	if userID == 0 {
		HandleError(c, errs.ErrRecordNotFound)
		return
	}
	var order models.Order
	err = c.BindJSON(&order)
	if err != nil {
		HandleError(c, errs.ErrValidationFailed)
		return
	}
	err = service.CheckOrderasResponse(order, int(userID), int(userID), int(id))
	if err != nil {
		HandleError(c, errs.ErrRoutesNotFound)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Check as response is succesfuly"})

}

// DeleteOrderByID
// @Summary Delete Order By ID
// @Security AKA
// @Tags orders
// @Description delete route by ID
// @ID delete-route-by-id
// @Param id path integer true "id of the route"
// @Success 200 {object} defaultResponse
// @Failure 400 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /api/orders/{id} [delete]
func DeleteOrderByID(c *gin.Context) {
	urole := c.GetString(userRoleCtx)
	if urole == "" {
		HandleError(c, errs.ErrValidationFailed)
		return
	}
	if urole != "admin" {
		HandleError(c, errs.ErrPermissionDenied)
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.Error.Printf("[controllers.DeleteOrderByID] invalid order_id path parameter: %s\n", c.Param("id"))
		HandleError(c, errs.ErrValidationFailed)
		return
	}
	err = service.DeleteOrder(true, id)
	if err != nil {
		HandleError(c, errs.ErrRoutesNotFound)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Deleted is succesfuly"})
}
