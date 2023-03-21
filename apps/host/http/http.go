package http

import (
	"github.com/elisaovivo/restful-api-demo/apps/host"
	"github.com/gin-gonic/gin"
)

func NewHostHTTPHandler(svc host.Service) *Handler {
	return &Handler{
		svc: svc,
	}
}

// 通过写一个实例类, 把内部的接口通过HTTP协议暴露出去
// 所以需要依赖内部接口的实现
// 该实体类, 会实现Gin的Http Handler
type Handler struct {
	svc host.Service
}

// 完成了 Http Handler的注册
func (h *Handler) Registry(r gin.IRouter) {
	r.POST("/hosts", h.createHost)
}
