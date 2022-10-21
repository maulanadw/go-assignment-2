package controller

import (
	"go-assignment-2/helper"
	"go-assignment-2/param"
	"go-assignment-2/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type OrderController interface {
	CreateOrder(c *gin.Context)
	GetOrder(c *gin.Context)
	UpdateOrder(c *gin.Context)
	DeleteOrder(c *gin.Context)
}

type orderController struct {
	orderService service.OrderService
	itemService  service.ItemService
}

func NewOrderController(orderService service.OrderService, itemService service.ItemService) OrderController {
	return &orderController{
		orderService: orderService,
		itemService:  itemService,
	}
}

// @Create Order
// @tag.name Order
// @Description API for create order
// @Accept json
// @Produce json
// @Tags Orders
// @param Body body param.Order true "Create order"
// @Success 200 {object} param.OrderResponse
// @Router /orders [POST]
func (or *orderController) CreateOrder(ctx *gin.Context) {
	var request param.Order
	var jsonResponse param.Response

	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	order, err := or.orderService.CreateOrder(*request.CustomerName)
	if err != nil {
		msgError := "failed to create order"
		errorInfo := err.Error()
		jsonResponse = helper.JsonResponse(http.StatusInternalServerError, &msgError, &errorInfo, nil)
		ctx.JSON(jsonResponse.Status, jsonResponse)
		return
	}

	err = or.itemService.CreateItem(order.ID, request.Items)
	if err != nil {
		msgError := "failed to create order"
		errorInfo := err.Error()
		jsonResponse = helper.JsonResponse(http.StatusInternalServerError, &msgError, &errorInfo, nil)
		ctx.JSON(jsonResponse.Status, jsonResponse)
		return
	}

	msgSuccess := "order successfully created"
	jsonResponse = helper.JsonResponse(http.StatusCreated, &msgSuccess, nil, nil)
	ctx.JSON(jsonResponse.Status, jsonResponse)
}

// @Get Order
// @tag.name Order
// @Description API for Get All order data
// @Produce json
// @Tags Orders
// @Success 200 {object} param.OrderResponse
// @Router /orders [GET]
func (or *orderController) GetOrder(ctx *gin.Context) {
	var jsonResponse param.Response

	orders, err := or.orderService.GetOrder()
	if err != nil {
		msgError := "failed to get order"
		errorInfo := err.Error()
		jsonResponse = helper.JsonResponse(http.StatusInternalServerError, &msgError, &errorInfo, nil)
		ctx.JSON(jsonResponse.Status, jsonResponse)
		return
	}

	jsonResponse = helper.JsonResponse(http.StatusOK, nil, nil, orders)
	ctx.JSON(jsonResponse.Status, jsonResponse)
}

// @Update Order
// @tag.name Order
// @Description API for update order data
// @Accept json
// @Produce json
// @Tags Orders
// @param Body body param.Order true "Update Order"
// @Success 200 {object} param.OrderResponse
// @Router /orders/:id [PUT]
func (or *orderController) UpdateOrder(ctx *gin.Context) {
	var request param.Order
	var jsonResponse param.Response

	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		msgError := "something is wrong with the param given"
		errorInfo := err.Error()
		jsonResponse = helper.JsonResponse(http.StatusBadRequest, &msgError, &errorInfo, nil)
		ctx.AbortWithStatusJSON(jsonResponse.Status, jsonResponse)
		return
	}

	err = ctx.ShouldBindJSON(&request)
	if err != nil {
		msgError := "something is wrong with the body request given"
		errorInfo := err.Error()
		jsonResponse = helper.JsonResponse(http.StatusBadRequest, &msgError, &errorInfo, nil)
		ctx.AbortWithStatusJSON(jsonResponse.Status, jsonResponse)
		return
	}

	err = or.orderService.UpdateOrder(id, request)
	if err != nil {
		msgError := "failed to update order"
		errorInfo := err.Error()
		jsonResponse = helper.JsonResponse(http.StatusBadRequest, &msgError, &errorInfo, nil)
		ctx.JSON(jsonResponse.Status, jsonResponse)
		return
	}

	if request.Items != nil {
		err = or.itemService.UpdateItemByOrderID(id, request.Items)
		if err != nil {
			msgError := "failed to update order"
			errorInfo := err.Error()
			jsonResponse = helper.JsonResponse(http.StatusBadRequest, &msgError, &errorInfo, nil)
			ctx.JSON(jsonResponse.Status, jsonResponse)
			return
		}
	}

	msgSuccess := "order successfully updated"
	jsonResponse = helper.JsonResponse(http.StatusOK, &msgSuccess, nil, nil)
	ctx.JSON(jsonResponse.Status, jsonResponse)
}

// @Delete Order
// @tag.name Order
// @Description API for Delete order data
// @Produce json
// @Tags Orders
// @Success 200 {object} param.OrderResponse
// @Router /orders/:id [DELETE]
func (or *orderController) DeleteOrder(ctx *gin.Context) {
	var jsonResponse param.Response

	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		msgError := "something is wrong with the param given"
		errorInfo := err.Error()
		jsonResponse = helper.JsonResponse(http.StatusBadRequest, &msgError, &errorInfo, nil)
		ctx.AbortWithStatusJSON(jsonResponse.Status, jsonResponse)
		return
	}

	err = or.orderService.DeleteOrder(id)
	if err != nil {
		msgError := "failed to delete order"
		errorInfo := err.Error()
		jsonResponse = helper.JsonResponse(http.StatusBadRequest, &msgError, &errorInfo, nil)
		ctx.JSON(jsonResponse.Status, jsonResponse)
		return
	}

	msgSuccess := "order successfully deleted"
	jsonResponse = helper.JsonResponse(http.StatusOK, &msgSuccess, nil, nil)
	ctx.JSON(jsonResponse.Status, jsonResponse)
}
