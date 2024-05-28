package model

type Restaurant struct {
	Id     uint   `json:"id"`
	Naam   string `json:"naam"`
	Teller uint   `json:"teller"`
	WebUrl string `json:"weburl"`
}

type PostRestaurant struct {
	Naam   string `json:"naam"`
	WebUrl string `json:"weburl"`
}

type RestaurantUri struct {
	ID uint `uri:"id" binding:"required,number"`
}

type UpdateRestaurant struct {
	Teller uint   `json:"teller"`
	WebUrl string `json:"weburl"`
}
