package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main(){
	router := gin.Default()

	// GET
	router.GET("/", rootHandler)
	router.GET("/hello", helloHandler)
	router.GET("/books/:id/:title", booksHandler)
	router.GET("/query", queryHandler)

	// POST
	router.POST("/books", postBooksHandler)

	router.Run()
	// untuk mengatur port bisa lakukan cara seperti di bawah ini
	// router.Run(":8888") 
	// ^ menggunakan port 8888
}

func rootHandler(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{
		"name":"Ananda Ricky",
		"bio":"Lorem Ipsum Dolor Sit Amet",
	})
}

func helloHandler(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{
		"title":"Hello World",
		"subtitle":"Belajar Go-Lang Web API",
	})
}

func booksHandler(c *gin.Context){
	id := c.Param("id")
	title := c.Param("title")

	c.JSON(http.StatusOK, gin.H{
		"id":id,
		"title":title,
	})
}

func queryHandler(c *gin.Context){
	title := c.Query("title")
	price := c.Query("price")

	c.JSON(http.StatusOK, gin.H{
		"title":title,
		"price":price,
	})
}


// post
type BookInput struct{
	Title string
	Price int
	SubTitle string `json:"sub_title"`
}

func postBooksHandler(c *gin.Context){

	var bookInput BookInput
	err := c.ShouldBindJSON(&bookInput)
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"title":bookInput.Title,
		"price":bookInput.Price,
		"sub_title":bookInput.SubTitle,
	})
}