package apps

import (
	"fmt"
	"github.com/elisaovivo/restful-api-demo/apps/host"
)

var (
	HostService host.Service
	// 维护当前所有的服务
	svcs = map[string]Service{}
)

func Registry(svc Service) {
	// 服务实例注册到svcs map当中
	if _, ok := svcs[svc.Name()]; ok {
		panic(fmt.Sprintf("service %s has registried", svc.Name()))
	}

	svcs[svc.Name()] = svc
	// 更加对象满足的接口来注册具体的服务
	if v, ok := svc.(host.Service); ok {
		HostService = v
	}
}

// 用户初始化 注册到Ioc容器里面的所有服务
func Init() {
	for _, v := range svcs {
		v.Config()
	}
}

type Service interface {
	Config()
	Name() string
}
