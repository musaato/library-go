package controllers

import (
	"library-Go/initializers"
	"library-Go/models"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func AuthorCreate(c *gin.Context) {
	// Define a struct to hold the request body
	var body struct {
		FirstName       string     `json:"firstName"`
		LastName        string     `json:"lastName"`
		ProfilePicture  *string    `json:"profilePicture"`
		Biography       *string    `json:"biography"`
		NationalityCode *string    `json:"nationalityCode"`
		Birthdate       *time.Time `json:"birthdate"`
	}

	// Bind the request body to the struct
	if err := c.BindJSON(&body); err != nil {
		log.Printf("Binding error: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Create the author model
	author := models.Author{
		FirstName:       body.FirstName,
		LastName:        body.LastName,
		ProfilePicture:  body.ProfilePicture,  // Directly assign *string
		Biography:       body.Biography,       // Directly assign *string
		NationalityCode: body.NationalityCode, // Directly assign *string
		Birthdate:       body.Birthdate,       // Directly assign *time.Time
		CreatedTime:     time.Now(),
		UpdatedTime:     time.Now(),
	}

	log.Printf("Author struct: %+v\n", author)

	// Create the author in the database
	result := initializers.DB.Create(&author)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to create author"})
		return
	}

	// Respond with the created author
	c.JSON(http.StatusCreated, gin.H{"author": author})
}

// fetch all authors
func AuthorIndex(c *gin.Context) {
	// get the authors
	var authors []models.Author

	initializers.DB.Find(&authors)

	// log.Printf("Author struct: %+v\n", authors)
	// respond with them
	c.JSON(200, gin.H{
		"authors": authors,
	})

}

// fetch single author
func AuthorShow(c *gin.Context) {
	// get author id off url
	id := c.Param("id")

	// get the author
	var author models.Author
	initializers.DB.First(&author, id)

	// respond with it
	c.JSON(200, gin.H{
		"author": author,
	})
}

// update author
func AuthorUpdate(c *gin.Context) {
	//  get the id off the url
	id := c.Param("id")
	// get the data off request body
	var body struct {
		FirstName       string     `json:"firstName"`
		LastName        string     `json:"lastName"`
		ProfilePicture  *string    `json:"profilePicture"`
		Biography       *string    `json:"biography"`
		NationalityCode *string    `json:"nationalityCode"`
		Birthdate       *time.Time `json:"birthdate"`
	}
	c.Bind(&body)

	// find the author were updating
	var author models.Author
	initializers.DB.First(&author, id)

	// update it
	initializers.DB.Model(&author).Updates(models.Author{
		FirstName:       body.FirstName,
		LastName:        body.LastName,
		ProfilePicture:  body.ProfilePicture,
		Biography:       body.Biography,
		NationalityCode: body.NationalityCode,
		Birthdate:       body.Birthdate,
		UpdatedTime:     time.Now(),
	})

	// respond with it
	c.JSON(200, gin.H{
		"author": author,
	})

}

// delete author
func AuthorDelete(c *gin.Context) {
	//  get the id off the url
	id := c.Param("id")

	// delete the author(s)
	initializers.DB.Delete(&models.Author{}, id)

	// respond
	c.JSON(200, gin.H{
		"message": "Author deleted.",
	})
}
