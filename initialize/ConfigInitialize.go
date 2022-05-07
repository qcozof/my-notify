//加载配置文件
package initialize

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/qcozof/my-notify/global"
	"github.com/spf13/viper"
)
var config string
func init()  {
	config = "config.yaml"

	v := viper.New()
	v.SetConfigFile(config)
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err := v.Unmarshal(&global.SERVER_CONFIG); err != nil {
			fmt.Println(err)
		}
	})
	if err := v.Unmarshal(&global.SERVER_CONFIG); err != nil {
		fmt.Println(err)
	}
	global.VIPER = v
}