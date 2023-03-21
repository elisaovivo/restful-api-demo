package impl

import (
	"database/sql"
	"github.com/elisaovivo/restful-api-demo/apps/host"
)

// 编译器的静态检查
var _ host.Service = (*HostServiceImpl)(nil)

// 保证调用该函数前，全局已经初始化
func NewHostServiceImpl() *HostServiceImpl {
	return &HostServiceImpl{}
}

type HostServiceImpl struct {
	db *sql.DB
}
