package http

import (
	"github.com/elisaovivo/restful-api-demo/apps/host"
	"github.com/gin-gonic/gin"
	"github.com/infraboard/mcube/http/response"
)

// 用于暴露host service接口
func (h *Handler) createHost(c *gin.Context) {
	ins := host.NewHost()
	// 用户传递过来的参数进行解析
	if err := c.Bind(ins); err != nil {
		response.Failed(c.Writer, err)
	}
	ins, err := h.svc.CreateHost(c.Request.Context(), ins)
	if err != nil {
		response.Failed(c.Writer, err)
	}
	response.Success(c.Writer, ins)
}
