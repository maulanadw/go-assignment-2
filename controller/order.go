package controller

import (
	"fmt"
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
		msg := fmt.Sprintf("Failed to create order: %s", err.Error())
		jsonResponse = helper.JsonResponse(http.StatusNotAcceptable, &msg, nil)
		ctx.JSON(jsonResponse.Status, jsonResponse)
		return
	}

	err = or.itemService.CreateItem(order.ID, request.Items)
	if err != nil {
		msg := fmt.Sprintf("Failed to create item order: %s", err.Error())
		jsonResponse = helper.JsonResponse(http.StatusNotAcceptable, &msg, nil)
		ctx.JSON(jsonResponse.Status, jsonResponse)
		return
	}

	jsonResponse = helper.JsonResponse(http.StatusOK, nil, nil)
	ctx.JSON(jsonResponse.Status, jsonResponse)
}

func (or *orderController) GetOrder(ctx *gin.Context) {
	var jsonResponse param.Response
	orders, err := or.orderService.GetOrder()
	if err != nil {
		msg := fmt.Sprintf("Failed to get orders data: %s", err.Error())
		jsonResponse = helper.JsonResponse(http.StatusBadRequest, &msg, nil)
		ctx.JSON(jsonResponse.Status, jsonResponse)
		return
	}

	jsonResponse = helper.JsonResponse(http.StatusOK, nil, orders)
	ctx.JSON(jsonResponse.Status, jsonResponse)
}

func (or *orderController) UpdateOrder(ctx *gin.Context) {
	var request param.Order
	var jsonResponse param.Response

	IDstr := ctx.Param("id")
	ID, err := strconv.Atoi(IDstr)
	if err != nil {
		msg := "something is wrong with the param given"
		jsonResponse = helper.JsonResponse(http.StatusBadRequest, &msg, nil)
		ctx.AbortWithStatusJSON(jsonResponse.Status, jsonResponse)
		return
	}

	err = ctx.ShouldBindJSON(&request)
	if err != nil {
		msg := "something is wrong with the body request given"
		jsonResponse = helper.JsonResponse(http.StatusBadRequest, &msg, nil)
		ctx.AbortWithStatusJSON(jsonResponse.Status, jsonResponse)
		return
	}

	err = or.orderService.UpdateOrder(ID, request)
	if err != nil {
		msg := fmt.Sprintf("Failed to update order: %s", err.Error())
		jsonResponse = helper.JsonResponse(http.StatusBadRequest, &msg, nil)
		ctx.JSON(jsonResponse.Status, jsonResponse)
		return
	}

	if request.Items != nil {
		err = or.itemService.UpdateItemByOrderID(ID, request.Items)
		if err != nil {
			msg := fmt.Sprintf("Failed to update order items: %s", err.Error())
			jsonResponse = helper.JsonResponse(http.StatusBadRequest, &msg, nil)
			ctx.JSON(jsonResponse.Status, jsonResponse)
			return
		}
	}

	jsonResponse = helper.JsonResponse(http.StatusOK, nil, nil)
	ctx.JSON(jsonResponse.Status, jsonResponse)
}

func (or *orderController) DeleteOrder(ctx *gin.Context) {
	var jsonResponse param.Response
	IDStr := ctx.Param("id")
	ID, err := strconv.Atoi(IDStr)
	if err != nil {
		msg := "something is wrong with the param given"
		jsonResponse = helper.JsonResponse(http.StatusBadRequest, &msg, nil)
		ctx.AbortWithStatusJSON(jsonResponse.Status, jsonResponse)
		return
	}

	err = or.orderService.DeleteOrder(ID)
	if err != nil {
		msg := fmt.Sprintf("Failed to delete order: %s", err.Error())
		jsonResponse = helper.JsonResponse(http.StatusBadRequest, &msg, nil)
		ctx.JSON(jsonResponse.Status, jsonResponse)
		return
	}

	jsonResponse = helper.JsonResponse(http.StatusOK, nil, "the data has been deleted")
	ctx.JSON(jsonResponse.Status, jsonResponse)
}
