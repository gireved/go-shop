package main

import (
	"go-shop/config"
	"go-shop/internal/startup"
)

func main() {
	loading()
}

func loading() {
	config.InitConfig()
	startup.InitMySQL()
}
