package router

import "github.com/gin-gonic/gin"

func Initializar() {
	router := gin.Default()
	initializeRoutes(router)
	router.Run(":3000")
}
