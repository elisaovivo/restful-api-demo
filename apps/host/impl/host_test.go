package impl

import (
	"context"
	"fmt"
	"github.com/elisaovivo/restful-api-demo/apps/host"
	"github.com/elisaovivo/restful-api-demo/conf"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	service host.Service
)

func TestCreate(t *testing.T) {
	should := assert.New(t)
	ins := host.NewHost()

	ins.Name = "test"
	ins.Id = "ins-01"
	ins.Region = "杭州"
	ins.Type = "sm1"
	ins.CPU = 1
	ins.Memory = 2048
	ins, err := service.CreateHost(context.Background(), ins)
	if should.NoError(err) {
		fmt.Println(ins)
	}
	service.CreateHost(context.Background(), ins)
}

func init() {
	err := conf.LoadConfigFromEnv()
	if err != nil {
		panic(err)
	}
	service = NewHostServiceImpl()
}
