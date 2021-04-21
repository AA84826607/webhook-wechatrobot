package router

import (
	"ceph/gin/handler"
	"github.com/gin-gonic/gin"
)

type Router struct {
	handler *handler.Handler
	port string
}

func NewRouter(h *handler.Handler,port string) (result Router) {
	return Router{
		handler: h,
		port: port,
	}
}
func (r *Router) Router() {

	router := gin.Default()
	userGroup := router.Group("/v1")
	{
		// TODO create_order格式如左边
		userGroup.POST("/prometheus_send", r.handler.PrometheusSend)
	}
	router.Run(":"+r.port)
}
