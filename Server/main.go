package main

import (
	"fmt"
	"mMoviez/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("So it begins...")

	var engine *gin.Engine = gin.Default()

	routes.SetupUnprotectedRoutes(engine)
	routes.SetupProtectedRoutes(engine)

	err := engine.Run(":8080")
	if err != nil {
		fmt.Println("Failed to start server:", err)
		return
	}
}
