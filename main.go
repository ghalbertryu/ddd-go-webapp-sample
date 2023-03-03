package main

import (
	"m/interface/rest"
)

func main() {
	router := rest.InitRouter()
	router.Run()
}
