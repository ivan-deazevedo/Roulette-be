package app

import (
	"database/sql"
	"lunch-be/controller"
	"lunch-be/database"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type App struct {
	DB     *sql.DB
	Router *gin.Engine
}

func (a *App) CreateConnection() {
	db := database.Connectdb()
	a.DB = db
}

func (a *App) Routes() {
	r := gin.Default()
	r.Use(cors.Default())

	controllerResto := controller.NewRestaurantController(a.DB)
	controllerUser := controller.NewUserController(a.DB)

	r.POST("/restaurant", controllerResto.InsertRestaurant)
	r.GET("/restaurant", controllerResto.GetAllRestaurants)
	r.GET("/restaurant/:id", controllerResto.GetOneRestaurant)
	r.DELETE("/restaurant/:id", controllerResto.DeleteRestaurant)
	r.PUT("/restaurant/:id", controllerResto.UpdateRestaurant)

	r.POST("/user", controllerUser.InsertUser)
	r.GET("/user", controllerUser.GetAllUsers)
	r.GET("/user/:id", controllerUser.GetOneUser)
	r.DELETE("/user/:id", controllerUser.DeleteUser)
	r.PUT("/user/:id", controllerUser.UpdateUser)

	a.Router = r
}

func (a *App) Run() {
	a.Router.Run(":4000")
}
