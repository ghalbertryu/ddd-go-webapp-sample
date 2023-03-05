package main

import (
	"m/infrastructure/repo"
	"m/interface/rest"
)

func main() {
	repo.InitRepository()
	router := rest.InitRouter()
	router.Run()
}
