package main

import (
	"github.com/gin-gonic/gin"
	"llm_gateway/routes"
)

func main() {
	r := gin.Default()
	r.GET("/addr", routes.GetAddress)
	r.POST("/addr", routes.UpdateAddress)
	r.Run(":8081") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
