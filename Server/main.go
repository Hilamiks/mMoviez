package main

import (
	"fmt"
	"mMoviez/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("So it begins...")

	var engine *gin.Engine = gin.Default()

	engine.GET("/hello", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello mMoviez")
	})

	engine.GET("/movies", controller.GetMovies())
	engine.GET("/movie/:imdb_id", controller.GetMovieByID())
	engine.POST("/movie", controller.AddMovie())

	err := engine.Run(":8080")
	if err != nil {
		fmt.Println("Failed to start server:", err)
		return
	}
}
