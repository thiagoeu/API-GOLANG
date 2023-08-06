package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// inicializa o router utilizando as configurações do gin
	r := gin.Default()

	// definindo uma rota
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
