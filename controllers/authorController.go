package controllers

import (
	"library-Go/initializers"
	"library-Go/models"
	"log"
	"net/http"
	"strconv"
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

// fetch paginated authors
// http://localhost:7070/authors?pageNum=1&pageSize=5
func AuthorIndex(c *gin.Context) {

	pageNumStr := c.Query("pageNum")
	pageSizeStr := c.Query("pageSize")

	pageNum, err1 := strconv.Atoi(pageNumStr)
	if err1 != nil || pageNum < 1 {
		pageNum = 1
	}
	pageSize, err2 := strconv.Atoi(pageSizeStr)
	if err2 != nil || pageSize < 1 {
		pageSize = 5
	}

	offset := (pageNum - 1) * pageSize

	var authors []models.Author
	var totalCount int64

	countResult := initializers.DB.Model(&models.Author{}).Count(&totalCount)
	fetchResult := initializers.DB.Order("id").
		Limit(pageSize).
		Offset(offset).
		Find(&authors)

	if countResult.Error != nil || fetchResult.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch authors' data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"authors":    authors,
		"totalCount": totalCount,
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
