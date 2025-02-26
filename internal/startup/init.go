package startup

import "go-shop/config"

func Init() {
	config.InitConfig()
	InitPrometheus()
	_ = InitMySQL()
	InitLogger()
}
