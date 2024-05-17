package controller

import (
	"database/sql"
	"lunch-be/model"
	"lunch-be/repository"

	"github.com/gin-gonic/gin"
)

type RestaurantController struct {
	Db *sql.DB
}

func NewRestaurantController(db *sql.DB) RestaurantControllerInterface {
	return &RestaurantController{Db: db}
}

func (r *RestaurantController) DeleteRestaurant(c *gin.Context) {
	DB := r.Db
	var uri model.RestaurantUri
	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(400, gin.H{"status": "failed", "msg": err})
		return
	}
	repo := repository.NewRestaurantRepository(DB)
	delete := repo.DeleteRestaurant(uri.ID)
	if delete {
		c.JSON(200, gin.H{"status": "success", "msg": "deleted restaurant successfully"})
		return
	} else {
		c.JSON(500, gin.H{"status": "failed", "msg": "delete restaurant failed"})
		return
	}

}

func (r *RestaurantController) GetAllRestaurants(c *gin.Context) {
	DB := r.Db
	repo := repository.NewRestaurantRepository(DB)
	get := repo.GetAllRestaurants()
	if get != nil {
		c.JSON(200, gin.H{"status": "success", "data": get, "msg": "get restaurants successfully"})
		return
	} else {
		c.JSON(200, gin.H{"status": "success", "data": nil, "msg": "restaurants not found"})
		return
	}
}

func (r *RestaurantController) GetOneRestaurant(c *gin.Context) {
	DB := r.Db
	var uri model.RestaurantUri
	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(400, gin.H{"status": "failed", "msg": err})
		return
	}
	repo := repository.NewRestaurantRepository(DB)
	get := repo.GetOneRestaurant(uri.ID)
	if (get != model.Restaurant{}) {
		c.JSON(200, gin.H{"status": "success", "data": get, "msg": "get restaurant successfully"})
		return
	} else {
		c.JSON(200, gin.H{"status": "success", "data": nil, "msg": "restaurant not found"})
		return
	}
}

func (r *RestaurantController) InsertRestaurant(c *gin.Context) {
	DB := r.Db
	var post model.PostRestaurant
	if err := c.ShouldBind(&post); err != nil {
		c.JSON(400, gin.H{"status": "failed", "msg": err})
		return
	}
	repo := repository.NewRestaurantRepository(DB)
	insert := repo.InsertRestaurant(post)
	if insert {
		c.JSON(200, gin.H{"status": "success", "msg": "insert restaurant successfully"})
		return
	} else {
		c.JSON(500, gin.H{"status": "failed", "msg": "insert restaurant failed"})
		return
	}
}

func (r *RestaurantController) UpdateRestaurant(c *gin.Context) {
	DB := r.Db
	var post model.UpdateRestaurant
	var uri model.RestaurantUri
	if err := c.ShouldBind(&post); err != nil {
		c.JSON(400, gin.H{"status": "failed", "msg": err})
		return
	}
	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(400, gin.H{"status": "failed", "msg": err})
		return
	}
	repository := repository.NewRestaurantRepository(DB)
	update := repository.UpdateRestaurant(uri.ID, post)
	if (update != model.Restaurant{}) {
		c.JSON(200, gin.H{"status": "success", "data": update, "msg": "update restaurant successfully"})
		return
	} else {
		c.JSON(500, gin.H{"status": "failed", "data": nil, "msg": "update restaurant failed"})
		return
	}
}
