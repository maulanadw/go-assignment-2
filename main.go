package main

import (
	"go-assignment-2/controller"
	"go-assignment-2/db"
	"go-assignment-2/repository"
	"go-assignment-2/router"
	"go-assignment-2/service"

	"github.com/gin-gonic/gin"
)

const APP_PORT = ":8080"

func main() {
	server := gin.Default()
	database := db.GetDB()

	itemRepo := repository.NewItemRepository(database)
	itemService := service.NewItemService(itemRepo)

	orderRepo := repository.NewOrderRepository(database)
	orderService := service.NewOrderService(orderRepo)
	orderController := controller.NewOrderController(orderService, itemService)
	orderRouter := router.NewOrderRouter(server, orderController)

	orderRouter.Start()
	server.Run(APP_PORT)
}
