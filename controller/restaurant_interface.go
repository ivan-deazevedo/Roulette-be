package controller

import "github.com/gin-gonic/gin"

type RestaurantControllerInterface interface {
	InsertRestaurant(*gin.Context)
	GetAllRestaurants(*gin.Context)
	GetOneRestaurant(*gin.Context)
	DeleteRestaurant(*gin.Context)
	UpdateRestaurant(*gin.Context)
}
