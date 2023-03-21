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

// 写一个实体，把内部的接口通过http协议暴露出去
// 依赖内部接口的实现
// 该实体，会实现Gin的Http Handler
type Handler struct {
	svc host.Service
}

// http handler注册
func (h *Handler) Registry(r gin.IRouter) {
	r.POST("/hosts", h.createHost)
}
