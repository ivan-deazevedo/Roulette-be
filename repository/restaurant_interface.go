package repository

import "lunch-be/model"

type RestaurantRepositoryInterface interface {
	InsertRestaurant(model.PostRestaurant) bool
	GetAllRestaurants() []model.Restaurant
	DeleteRestaurant(uint) bool
}
