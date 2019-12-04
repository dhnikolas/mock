package main

import (
	"mock/internal/routes"
	"mock/pkg/jsonconfig"
)

func main() {
	cm, err := jsonconfig.GetConfigMap()
	if err != nil {
		panic(err)
	}

	routes.Init(cm)
}
