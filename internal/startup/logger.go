package startup

import (
	"fmt"
	"go-shop/pkg/logger"
)

func InitLogger() {
	// 初始化日志记录器
	err := logger.InitLog()
	if err != nil {
		fmt.Println("无法初始化日志记录器:", err)
		return
	}
}
