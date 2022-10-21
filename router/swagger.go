package router

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type swaggerRouter struct {
	router *gin.Engine
}

func NewSwaggerRouter(r *gin.Engine) *swaggerRouter {
	return &swaggerRouter{
		router: r,
	}

}

func (sr *swaggerRouter) Start() {
	sr.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
