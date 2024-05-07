package model

type Restaurant struct {
	Id   uint   `json:"id"`
	Naam string `json:"naam"`
}

type PostRestaurant struct {
	Naam string `json:"naam"`
}

type RestaurantUri struct {
	ID uint `uri:"id" binding"required,number"`
}
