package routes

import (
	"mMoviez/controller"
	"mMoviez/middleware"

	"github.com/gin-gonic/gin"
)

func SetupProtectedRoutes(engine *gin.Engine) {
	engine.Use(middleware.AuthMiddleware())

	engine.GET("/movie/:imdb_id", controller.GetMovieByID())
	engine.POST("/movie", controller.AddMovie())
}
