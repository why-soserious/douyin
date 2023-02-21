package main

import (
	"github.com/gin-gonic/gin"
	"github.com/why-soserious/douyin/controller"
)

func initRouter(router *gin.Engine) {
	router.Static("/static", "/public")
	apirouter := router.Group("/douyin")

	// basic apis
	apirouter.GET("/feed", controller.Feed)
	

}