package impl

import "github.com/elisaovivo/restful-api-demo/apps/host"

// 编译器的静态检查
var _ host.Service = (*HostServiceImpl)(nil)

func NewHostServiceImpl() *HostServiceImpl {
	return &HostServiceImpl{}
}

type HostServiceImpl struct {
}
