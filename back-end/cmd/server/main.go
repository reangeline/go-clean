package main

import "github.com/reangeline/go-clean/configs"

func main() {
	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}
}
