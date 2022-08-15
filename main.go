package main

import "url-get/initRouter"

func main() {
	router := initRouter.SetupRouter()
	router.Run()
}
