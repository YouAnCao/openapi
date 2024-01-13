package initialize

import (
	"fmt"
	"openapi/global"

	"github.com/spf13/viper"
)

// LoadConfig 加载配置文件
func LoadConfig() {
	viper.AddConfigPath("./")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("读取配置异常！")
		return
	}
	if err := viper.Unmarshal(&global.Config); err != nil {
		fmt.Println("引入配置异常！")
	}
}
