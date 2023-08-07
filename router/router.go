package router

import "github.com/gin-gonic/gin"

func Initialize() {
	// Inicializando o router utilizando as configurações Default do gin
	router := gin.Default()
	// Definindo uma Rota
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.Run(":8080")
}
