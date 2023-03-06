package main

import (
	"m/infrastructure/repo"
	"m/interface/rest"
)

func main() {
	repo.Init()
	router := rest.InitRouter()
	router.Run()
}
