package main

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type restaurants struct {
	ID           string `json:"id"`
	Naam         string `json:"naam"`
	Omschrijving string `json:"omschrijving"`
}

type opties []restaurants

var data = opties{
	{ID: "1", Naam: "Subway", Omschrijving: "Broodjes zaak"},
	{ID: "2", Naam: "Khun Kaew", Omschrijving: "Thais"},
	{ID: "3", Naam: "Dominos", Omschrijving: "Pizza"},
	{ID: "4", Naam: "Minh Anh", Omschrijving: "Vietnamees"},
	{ID: "5", Naam: "Pat's Tosti Bar", Omschrijving: "Broodjes zaak"},
	{ID: "6", Naam: "Cho Pain", Omschrijving: "Broodjes zaak"},
}

func main() {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "http://github.com"
		},
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
