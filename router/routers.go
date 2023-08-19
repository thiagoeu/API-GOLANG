package router

import (
	"github.com/gin-gonic/gin"
	"github.com/thiagoeu/GOstd/handler"
)

func initializeRouters(router *gin.Engine) {
	//initialize handler
	handler.InitializeHandler()

	v1 := router.Group("/api/v1")
	{
		// show opening
		v1.GET("/opening", handler.ShowOpeningHandle)
		// post
		v1.POST("/opening", handler.CreateOpeningHandle)
		// delete
		v1.DELETE("/opening", handler.DeleteOpeningHandle)
		// update
		v1.PUT("/opening", handler.UpdateOpeningHandle)
		// list
		v1.GET("/openings", handler.ListOpeningsHandle)
	}

}
