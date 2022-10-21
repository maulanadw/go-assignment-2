package main

import (
	"go-assignment-2/controller"
	"go-assignment-2/db"
	"go-assignment-2/repository"
	"go-assignment-2/router"
	"go-assignment-2/service"

	_ "go-assignment-2/docs"

	"github.com/gin-gonic/gin"
)

const APP_PORT = ":8080"

// @title           Hacktiv8 Assignment 2
// @version         1.0
// @description     Build Rest API In GO

// @contact.name   Maulana Dwi Wahyudi
// @contact.email  maulana@email.com

// @host      localhost:8080
// @BasePath  /
func main() {
	server := gin.Default()
	database := db.GetDB()

	itemRepo := repository.NewItemRepository(database)
	itemService := service.NewItemService(itemRepo)

	orderRepo := repository.NewOrderRepository(database)
	orderService := service.NewOrderService(orderRepo)
	orderController := controller.NewOrderController(orderService, itemService)
	orderRouter := router.NewOrderRouter(server, orderController)

	swaggerRouter := router.NewSwaggerRouter(server)

	orderRouter.Start()
	swaggerRouter.Start()
	server.Run(APP_PORT)
}
