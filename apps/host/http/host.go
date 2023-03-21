package http

import (
	"github.com/elisaovivo/restful-api-demo/apps/host"
	"github.com/gin-gonic/gin"
	"github.com/infraboard/mcube/http/response"
)

// 用于暴露Host service接口
func (h *Handler) createHost(c *gin.Context) {
	ins := host.NewHost()
	// 用户传递过来的参数进行解析, 实现了一个json 的unmarshal
	if err := c.Bind(ins); err != nil {
		response.Failed(c.Writer, err)
		return
	}

	// 进行接口调用
	ins, err := h.svc.CreateHost(c.Request.Context(), ins)
	if err != nil {
		response.Failed(c.Writer, err)
		return
	}

	// 成功, 把对象实例返回给HTTP API调用方
	response.Success(c.Writer, ins)
}
