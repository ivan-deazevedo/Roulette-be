package controller

import "github.com/gin-gonic/gin"

type RestaurantControllerInterface interface {
	InsertRestaurant(*gin.Context)
	GetAllRestaurants(*gin.Context)
	DeleteRestaurant(*gin.Context)
}
