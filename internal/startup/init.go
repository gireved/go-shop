package startup

import "go-shop/config"

func Init() {
	config.InitConfig()
	_ = InitMySQL()
	InitLogger()
}
