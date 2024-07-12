package main

import (
	"app/adapter"
	"app/api"
	"app/router"
	"app/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	repo_init := adapter.Init()
	usecase_init := usecase.Init(repo_init)
	router_init := router.Init(usecase_init)

	// calling the rest api calls
	api.Router(r, router_init)

	r.Run()
}
