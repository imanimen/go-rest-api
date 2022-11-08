package main

import (
	_"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Book struct {
	ID		int32 `json:"id"`
	Title 	string `json:"title"`
	Author  string `json:"author"`
}

var books = []Book {
	{ID: 1, Title: "BELIEVE", Author: "JIM"},
	{ID: 2, Title: "FAITH", Author: "JIM"},
	{ID: 3, Title: "MYTH", Author: "JIM"},
}

func main(){
	r := gin.New()
	r.GET("/", func (c *gin.Context)  {
		c.JSON(http.StatusOK, gin.H {
			"data" : "data",
			"errors": "",
			"messages": "",
			"code": 200,
		})
	})

	r.GET("/index", func (c *gin.Context)  {
		c.JSON(http.StatusOK, gin.H {
			"data" : books,
			"errors": "",
			"messages": "",
			"code": 200,
		})
	})
	r.Run()
}