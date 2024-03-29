package cmd

import (
	"github.com/elisaovivo/restful-api-demo/apps"
	_ "github.com/elisaovivo/restful-api-demo/apps/all"
	"github.com/elisaovivo/restful-api-demo/conf"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

var (
	// pusher service config option
	confType string
	confFile string
	confETCD string
)

// 程序的启动时 组装都在这里进行
// 1.
// StartCmd represents the base command when called without any subcommands
var StartCmd = &cobra.Command{
	Use:   "start",
	Short: "启动 demo 后端API",
	Long:  "启动 demo 后端API",
	RunE: func(cmd *cobra.Command, args []string) error {
		// 加载程序配置
		err := conf.LoadConfigFromToml(confFile)
		if err != nil {
			panic(err)
		}

		apps.InitImpl()

		g := gin.Default()
		apps.InitGin(g)

		return g.Run(conf.C().App.HttpAddr())
	},
}

// 还没有初始化Logger实例

func init() {
	StartCmd.PersistentFlags().StringVarP(&confFile, "config", "f", "etc/demo.toml", "demo api 配置文件路径")
	RootCmd.AddCommand(StartCmd)
}
