package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/devs", getDevs)

	router.Run(":8085")
}

// deb represents data about a Dev
type dev struct {
	ID   string `json:"id"`
	name string `json:"name"`
	role string `json:"role"`
}

// devs slice to seed record dev data.
var devs = []dev{
	{ID: "1", name: "Josefina J", role: "frontend dev"},
}

// getDevs responds with the list of all Devs as JSON
func getDevs(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, devs)
}
