package routes

import (
	"mMoviez/controller"

	"github.com/gin-gonic/gin"
)

func SetupUnprotectedRoutes(engine *gin.Engine) {
	engine.GET("/movies", controller.GetMovies())
	engine.POST("/register", controller.AddUser())
	engine.POST("/login", controller.LoginUser())
}
