package impl

import (
	"context"
	"github.com/elisaovivo/restful-api-demo/apps/host"
	"testing"
)

var (
	service host.Service
)

func TestCreate(t *testing.T) {
	ins := host.NewHost()
	ins.Name = "test"
	service.CreateHost(context.Background(), ins)
}

func init() {
	service = NewHostServiceImpl()
}
