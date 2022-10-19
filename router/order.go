package router

import (
	"go-assignment-2/controller"

	"github.com/gin-gonic/gin"
)

type orderRouter struct {
	router          *gin.Engine
	orderController controller.OrderController
}

func NewOrderRouter(r *gin.Engine, oc controller.OrderController) *orderRouter {
	return &orderRouter{
		router:          r,
		orderController: oc,
	}

}

func (or *orderRouter) Start() {
	or.router.POST("/orders", or.orderController.CreateOrder)
	or.router.GET("/orders", or.orderController.GetOrder)
	or.router.PUT("/orders/:id", or.orderController.UpdateOrder)
	or.router.DELETE("/orders/:id", or.orderController.DeleteOrder)
}
