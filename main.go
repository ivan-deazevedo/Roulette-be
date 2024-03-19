package main

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type restaurants struct {
	ID   string `json:"id"`
	Naam string `json:"naam"`
}

type opties []restaurants

var data = opties{
	{ID: "1", Naam: "Subway"},
	{ID: "2", Naam: "Khun Kaew"},
	{ID: "3", Naam: "Dominos"},
	{ID: "4", Naam: "Minh Anh"},
	{ID: "5", Naam: "Pat's Tosti Bar"},
	{ID: "6", Naam: "Cho Pain"},
	{ID: "7", Naam: "Enzo's"},
}

func main() {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	router.GET("/opties", getOptions)
	router.GET("/opties/:id", getRestaurantByID)
	router.POST("/opties", postOptions)

	router.Run(":4000")
}

func getOptions(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, data)
}

func postOptions(c *gin.Context) {
	var newRestaurant restaurants

	if err := c.BindJSON(&newRestaurant); err != nil {
		return
	}

	data = append(data, newRestaurant)
	c.IndentedJSON(http.StatusCreated, newRestaurant)
}

func getRestaurantByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range data {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Restaurant not found"})
}
