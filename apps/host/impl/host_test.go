package impl_test

import (
	"context"
	"fmt"
	"github.com/elisaovivo/restful-api-demo/apps/host"
	"github.com/elisaovivo/restful-api-demo/apps/host/impl"
	"github.com/elisaovivo/restful-api-demo/conf"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/stretchr/testify/assert"

	"testing"
)

var (
	// 定义对象是满足该接口的实例
	service host.Service
)

func TestCreate(t *testing.T) {
	should := assert.New(t)
	ins := host.NewHost()
	ins.Id = "ins-01"
	ins.Name = "test"
	ins.Region = "cn-hangzhou"
	ins.Type = "sm1"
	ins.CPU = 1
	ins.Memory = 2048
	ins, err := service.CreateHost(context.Background(), ins)
	if should.NoError(err) {
		fmt.Println(ins)
	}
}

func init() {
	// 测试用例的配置文件
	err := conf.LoadConfigFromEnv()
	if err != nil {
		panic(err)
	}

	// 需要初始化全局Logger,
	// 为什么不设计为默认打印, 因为性能
	fmt.Println(zap.DevelopmentSetup())

	// host service 的具体实现
	service = impl.NewHostServiceImpl()
}
