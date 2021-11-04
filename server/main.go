package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/devs", getDevs)
	router.POST("/devs", postDevs)

	router.Run(":8085")
}

// deb represents data about a Dev
type dev struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Role string `json:"role"`
}

// devs slice to seed record dev data.
var devs = []dev{
	{ID: "1", Name: "Josefina J", Role: "frontend dev"},
}

// getDevs responds with the list of all Devs as JSON
func getDevs(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, devs)
}

func postDevs(c *gin.Context) {
	var newDev dev

	if err := c.BindJSON(&newDev); err != nil {
		return
	}

	devs = append(devs, newDev)
	c.IndentedJSON(http.StatusCreated, newDev)
}
