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
	controller := controller.NewRestaurantController(a.DB)

	r.Use(cors.Default())

	r.POST("/restaurant", controller.InsertRestaurant)
	r.GET("/restaurant", controller.GetAllRestaurants)
	r.GET("/restaurant/:id", controller.GetOneRestaurant)
	r.DELETE("/restaurant/:id", controller.DeleteRestaurant)
	r.PUT("/restaurant/:id", controller.UpdateRestaurant)
	a.Router = r
}

func (a *App) Run() {
	a.Router.Run(":4000")
}
