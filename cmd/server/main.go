package main

import "github.com/Msaorc/Go-APIs/configs"

func main() {
	config, _ := configs.LoadConfigs(".")
	println(config.DBDriver)
}
