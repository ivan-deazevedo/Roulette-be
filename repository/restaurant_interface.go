package repository

import "lunch-be/model"

type RestaurantRepositoryInterface interface {
	InsertRestaurant(model.PostRestaurant) bool
	GetAllRestaurants() []model.Restaurant
	GetOneRestaurant(uint) model.Restaurant
	DeleteRestaurant(uint) bool
	UpdateRestaurant(uint, model.UpdateRestaurant) model.Restaurant
}
