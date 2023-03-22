package apps

import (
	"fmt"
	"github.com/elisaovivo/restful-api-demo/apps/host"
	"github.com/gin-gonic/gin"
)

var (
	HostService host.Service
	// 维护当前所有的服务
	implApps = map[string]ImplService{}
	ginApps  = map[string]GinService{}
)

func RegistryImpl(svc ImplService) {
	// 服务实例注册到svcs map当中
	if _, ok := implApps[svc.Name()]; ok {
		panic(fmt.Sprintf("service %s has registried", svc.Name()))
	}

	implApps[svc.Name()] = svc
	// 更加对象满足的接口来注册具体的服务
	if v, ok := svc.(host.Service); ok {
		HostService = v
	}
}

// 用户初始化 注册到Ioc容器里面的所有服务
func GetImpl(name string) interface{} {
	for k, v := range implApps {
		if k == name {
			return v
		}
	}

	return nil
}

func RegistryGin(svc GinService) {
	// 服务实例注册到svcs map当中
	if _, ok := ginApps[svc.Name()]; ok {
		panic(fmt.Sprintf("service %s has registried", svc.Name()))
	}

	ginApps[svc.Name()] = svc
}

func InitImpl() {
	for _, v := range implApps {
		v.Config()
	}
}

func InitGin(r gin.IRouter) {
	// 先初始化好所有对象
	for _, v := range ginApps {
		v.Config()
	}

	// 完成Http Handler的注册
	for _, v := range ginApps {
		v.Registry(r)
	}
}

type ImplService interface {
	Config()
	Name() string
}

// 注册Gin编写的Handler
// 比如 编写了Http服务A, 只需要实现Registry方法, 就能把Handler注册给Root Router
type GinService interface {
	Registry(r gin.IRouter)
	Config()
	Name() string
}
