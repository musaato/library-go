package main

import (
	"library-Go/controllers"
	"library-Go/initializers"

	"github.com/gin-gonic/gin"
)

// CompileDaemon -command="./library-Go"
func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {

	r := gin.Default()

	// create a new author
	r.POST("/authors", controllers.AuthorCreate)

	// fetch all authors
	r.GET("/authors", controllers.AuthorIndex)

	// fetch an author
	r.GET("/authors/:id", controllers.AuthorShow)

	// update author(s)
	r.PUT("/authors/:id", controllers.AuthorUpdate)

	// delete authors(s)
	r.DELETE("/authors/:id", controllers.AuthorDelete)

	r.Run()

}
