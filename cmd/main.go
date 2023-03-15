package main

import (
	service "github.com/sorohimm/ellie/internal/service/ellie"
)

var version, buildTime string

func main() {
	service.NewService().Init("pythia-apigw", version, buildTime)
}
