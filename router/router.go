package router

import "github.com/gin-gonic/gin"

// USAR LETRA MAIUSCULA  NA FUNÇÃO FAZ COM QUE ELA SEJA EXPORTADA PARA OUTROS MODULOS, SE USAR MINUSCULA NÃO FICA DISPONIVEL.

func Initialize() {
	// inicializa o router utilizando as configurações do gin
	router := gin.Default()

	// definindo uma rota
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.Run(":8080") // listen and serve on 0.0.0.0:8080
}
