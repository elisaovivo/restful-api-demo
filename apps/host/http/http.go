package http

import (
	"github.com/elisaovivo/restful-api-demo/apps"
	"github.com/elisaovivo/restful-api-demo/apps/host"
	"github.com/gin-gonic/gin"
)

var handler = &Handler{}

// 通过写一个实例类, 把内部的接口通过HTTP协议暴露出去
// 所以需要依赖内部接口的实现
// 该实体类, 会实现Gin的Http Handler
type Handler struct {
	svc host.Service
}

func (h *Handler) Config() {
	h.svc = apps.GetImpl(host.AppName).(host.Service)
}

// 完成了 Http Handler的注册
func (h *Handler) Registry(r gin.IRouter) {
	r.POST("/hosts", h.createHost)
}

func (h *Handler) Name() string {
	return host.AppName
}

// 完成Http Handler注册
func init() {
	apps.RegistryGin(handler)
}
