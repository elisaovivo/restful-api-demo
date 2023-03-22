package impl

import (
	"database/sql"
	"github.com/elisaovivo/restful-api-demo/apps"
	"github.com/elisaovivo/restful-api-demo/apps/host"
	"github.com/elisaovivo/restful-api-demo/conf"

	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
)

// 接口实现的静态检查
var impl = &HostServiceImpl{}

// NewHostServiceImpl 保证调用该函数之前, 全局conf对象已经初始化
func NewHostServiceImpl() *HostServiceImpl {
	return &HostServiceImpl{
		// Host service 服务的子Loggger
		// 封装的Zap让其满足 Logger接口
		// 为什么要封装:
		// 		1. Logger全局实例
		// 		2. Logger Level的动态调整, Logrus不支持Level共同调整
		// 		3. 加入日志轮转功能的集合
		l:  zap.L().Named("Host"),
		db: conf.C().MySQL.GetDB(),
	}
}

type HostServiceImpl struct {
	l  logger.Logger
	db *sql.DB
}

func (i *HostServiceImpl) Config() {
	// Host service 服务的子Loggger
	// 封装的Zap让其满足 Logger接口
	// 为什么要封装:
	// 		1. Logger全局实例
	// 		2. Logger Level的动态调整, Logrus不支持Level共同调整
	// 		3. 加入日志轮转功能的集合
	i.l = zap.L().Named("Host")
	i.db = conf.C().MySQL.GetDB()
}

// 服务服务的名称
func (i *HostServiceImpl) Name() string {
	return host.AppName
}

// _ import app 自动执行注册逻辑
func init() {
	//  对象注册到ioc层
	apps.Registry(impl)
}
