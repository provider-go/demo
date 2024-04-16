package demo

import (
	"github.com/gin-gonic/gin"
	"github.com/provider-go/demo/router"
)

type Plugin struct{}

func CreatePlugin() *Plugin {
	return &Plugin{}
}

func (*Plugin) Register(group *gin.RouterGroup) {
	router.GroupApp.InitRouter(group)
}

func (*Plugin) RouterPath() string {
	return "demo"
}
