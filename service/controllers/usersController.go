package controllers

import (
	"net/http"
	"qgen/challange/databases"
	"qgen/challange/models"

	"github.com/gin-gonic/gin"
)

// GET /users
// Get all users
func AllUsers(c *gin.Context) {
	var users []models.User
	DBInstance := c.MustGet("db").(databases.IDBAdapter)
	DBInstance.Find(&users)

	c.JSON(http.StatusOK, gin.H{"data": users})
}

// GET /users/:id
// Find a user
func FindUser(c *gin.Context) {
	// Get model if exist
	var user []models.User
	DBInstance := c.MustGet("db").(databases.IDBAdapter)

	DBInstance.Find(&user, "id = ?", c.Param("id"))
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

	DBInstance.Find(&user, "id = ?", c.Param("id"))
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

	DBInstance.Find(&user, "id = ?", c.Param("id"))
	if len(user) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	DBInstance.Delete(&user)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
