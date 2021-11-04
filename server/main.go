package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Endpoints
	router.GET("/devs", getDevs)
	router.POST("/devs", postDevs)

	router.Run(":8085")
}

// dev represents data about a Dev
type dev struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Role string `json:"role"`
}

// devs slice to seed record dev data.
var devs = []dev{
	{ID: "1", Name: "Josefina J", Role: "frontend dev"},
}

// getDevs GET a list of all Devs as JSON
func getDevs(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, devs)
}

// postDevs POST a new Dev to the list in JSON
func postDevs(c *gin.Context) {
	var newDev dev

	if err := c.BindJSON(&newDev); err != nil {
		return
	}

	devs = append(devs, newDev)
	c.IndentedJSON(http.StatusCreated, newDev)
}
