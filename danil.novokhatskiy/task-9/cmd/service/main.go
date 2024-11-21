package main

import (
	"fmt"

	"github.com/katagiriwhy/task-9/internal/config"
)

type Contact struct {
	name  string
	phone string
}

var testTable []Contact = []Contact{
	{
		name:  "Nikita",
		phone: "89289019785",
	},
	{
		name:  "",
		phone: "",
	},
}

func main() {

	pathOfCfg := config.ReadFlag()

	cfg := config.LoadConfig(pathOfCfg)

	fmt.Println(cfg)

}
