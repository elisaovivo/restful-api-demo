package http

import (
	"github.com/elisaovivo/restful-api-demo/apps"
	"github.com/elisaovivo/restful-api-demo/apps/host"
	"github.com/gin-gonic/gin"
)

func NewHostHTTPHandler() *Handler {
	return &Handler{}
}

// 通过写一个实例类, 把内部的接口通过HTTP协议暴露出去
// 所以需要依赖内部接口的实现
// 该实体类, 会实现Gin的Http Handler
type Handler struct {
	svc host.Service
}

func (h *Handler) Config() {
	if apps.HostService == nil {
		panic("dependence host service required")
	}
	// 从IOC里面获取HostService的实例对象
	h.svc = apps.HostService
}

// 完成了 Http Handler的注册
func (h *Handler) Registry(r gin.IRouter) {
	r.POST("/hosts", h.createHost)
}
