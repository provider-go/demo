package router

import (
	"github.com/gin-gonic/gin"
	"github.com/provider-go/demo/api"
)

type Group struct {
	Router
}

var GroupApp = new(Group)

type Router struct{}

func (s *Router) InitRouter(Router *gin.RouterGroup) {
	{
		Router.GET("a", api.SayA)
		Router.GET("b", api.SayB)
	}
}
