package router

import "github.com/gin-gonic/gin"

func Initialize() {
	// initialize router
	router := gin.Default()

	// initializeRoutes
	initializeRouters(router)

	// run server
	router.Run(":8080") // listen and serve on 0.0.0.0:8080
}
