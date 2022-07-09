package controllers

import (
	"net/http"
	"qgen/challange/databases"
	"qgen/challange/models"

	"github.com/gin-gonic/gin"
)

// GET /users
// Get all users
func GetUsers(c *gin.Context) {
	var users []models.User
	DBInstance := c.MustGet("db").(databases.IDBAdapter)

	// init pagination
	pagination := models.Pagination{}
	pagination.Init(c)

	// get search key from query
	search := c.Query("search")
	if search != "" {
		DBInstance.Find(&users, &pagination, "name like ?", "%"+search+"%")
	} else {
		DBInstance.Find(&users, &pagination)
	}

	c.JSON(http.StatusOK, gin.H{"data": users, "pagination": pagination})
}

// GET /users/:id
// Find a user
func GetUserByID(c *gin.Context) {
	// Get model if exist
	var user []models.User
	DBInstance := c.MustGet("db").(databases.IDBAdapter)

	DBInstance.Find(&user, nil, "id = ?", c.Param("id"))
	if len(user) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user[0]})
}

// POST /users
// Create a user
func CreateUser(c *gin.Context) {
	// Validate input
	var input models.CreateUserInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create user
	user := models.User{Name: input.Name, Email: input.Email, Phone: input.Phone}
	DBInstance := c.MustGet("db").(databases.IDBAdapter)
	DBInstance.Create(&user)

	c.JSON(http.StatusOK, gin.H{"data": user})
}

// PATCH /users/:id
// Update a user
func UpdateUser(c *gin.Context) {
	// Get model if exist
	var user []models.User
	DBInstance := c.MustGet("db").(databases.IDBAdapter)

	DBInstance.Find(&user, nil, "id = ?", c.Param("id"))
	if len(user) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input models.UpdateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	DBInstance.Update(&user[0], models.User{Name: input.Name, Email: input.Email, Phone: input.Phone})

	c.JSON(http.StatusOK, gin.H{"data": user})
}

// DELETE /users/:id
// Delete a user
func DeleteUser(c *gin.Context) {
	// Get model if exist
	var user []models.User
	DBInstance := c.MustGet("db").(databases.IDBAdapter)

	DBInstance.Find(&user, nil, "id = ?", c.Param("id"))
	if len(user) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	DBInstance.Delete(&user)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
