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
	_, err := r.Db.Exec("DELETE FROM restaurants WHERE restaurantid = $1", id)
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
				id   uint
				naam string
			)
			err := query.Scan(&id, &naam)
			if err != nil {
				log.Println(err)
			}
			resto := model.Restaurant{Id: id, Naam: naam}
			restaurants = append(restaurants, resto)
		}
	}
	return restaurants
}

func (r *RestaurantRepository) InsertRestaurant(post model.PostRestaurant) bool {
	statement, err := r.Db.Prepare("INSERT INTO restaurants(restonaam) VALUES($1)")
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
