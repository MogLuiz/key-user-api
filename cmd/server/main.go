package main

import "github.com/MogLuiz/key-user-api/configs"

func main() {
	config, _ := configs.LoadConfig(".")
	println(config.DBName)
}
