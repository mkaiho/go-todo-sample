package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mkaiho/go-todo-sample/util"
)

func main() {
	router := gin.Default()

	router.GET("/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		if util.IsEmptyString(id) {
			statusCode := http.StatusBadRequest
			c.JSON(statusCode, gin.H{
				"message": fmt.Sprintf("%s: id=\"%v\"", http.StatusText(statusCode), id),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"id":   id,
			"name": "test_user",
		})
	})
	router.Run(":3000")

	fmt.Println("Hello World!")
}
