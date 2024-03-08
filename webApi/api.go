package webapi

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostOptions(c *gin.Context) {
	var newOption Restaurants

	if err := c.BindJSON(&newOption); err != nil {
		return
	}

	Data = append(Data, newOption)
	c.IndentedJSON(http.StatusCreated, newOption)
}

func GetOptions(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, Data)
}
