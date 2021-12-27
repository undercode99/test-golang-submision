package main

import (
	"server-grpc/config"
	"server-grpc/rest-api/handler"
)

func main() {
	r := handler.SetupRouters()
	r.Run(config.RestPort())
}
