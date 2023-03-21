package conf

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/caarlos0/env/v6"
)

//将配置成config对象

func LoadConfigFromToml(filePath string) error {
	config = NewDefaultConfig()

	//读取Toml配置
	_, err := toml.DecodeFile(filePath, config)
	if err != nil {
		return fmt.Errorf("load config from file,error,path:%s,%s", filePath, err)
	}
	return nil
}

func LoadConfigFromEnv() error {
	cfg := NewDefaultConfig()
	if err := env.Parse(cfg); err != nil {
		return err
	}
	// 加载全局配置单例
	config = cfg
	return nil
}
