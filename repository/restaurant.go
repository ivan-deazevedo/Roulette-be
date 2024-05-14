package repository

import (
	"database/sql"
	"log"
	"lunch-be/model"
)

type RestaurantRepository struct {
	Db *sql.DB
}

func NewRestaurantRepository(db *sql.DB) RestaurantRepositoryInterface {
	return &RestaurantRepository{Db: db}
}

func (r *RestaurantRepository) DeleteRestaurant(id uint) bool {
	_, err := r.Db.Exec("DELETE FROM restaurants WHERE restaurant_id = $1", id)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

func (r *RestaurantRepository) GetAllRestaurants() []model.Restaurant {
	query, err := r.Db.Query("SELECT * FROM restaurants")
	if err != nil {
		log.Println(err)
		return nil
	}
	var restaurants []model.Restaurant
	if query != nil {
		for query.Next() {
			var (
				id     uint
				naam   string
				teller int
			)
			err := query.Scan(&id, &naam, &teller)
			if err != nil {
				log.Println(err)
			}
			resto := model.Restaurant{Id: id, Naam: naam, Teller: teller}
			restaurants = append(restaurants, resto)
		}
	}
	return restaurants
}

func (r *RestaurantRepository) InsertRestaurant(post model.PostRestaurant) bool {
	statement, err := r.Db.Prepare("INSERT INTO restaurants(naam) VALUES($1)")
	if err != nil {
		log.Println(err)
		return false
	}
	defer statement.Close()

	_, err2 := statement.Exec(post.Naam)
	if err2 != nil {
		log.Println(err2)
		return false
	}
	return true
}

func (r *RestaurantRepository) UpdateRestaurant(id uint, post model.UpdateRestaurant) model.Restaurant {
	_, err := r.Db.Exec("UPDATE restaurants SET picked = $1 WHERE restaurant_id = $2", post.Teller, id)
	if err != nil {
		log.Println(err)
	}
	return model.Restaurant{}
}
