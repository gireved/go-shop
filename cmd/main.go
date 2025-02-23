package main

import (
	"go-shop/config"
	"go-shop/internal/startup"
)

func main() {

}

func init() {
	config.InitConfig()
	startup.InitMySQL()
}
