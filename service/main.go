package main

import (
	"qgen/challange/controllers"
	"qgen/challange/databases"

	"github.com/gin-gonic/gin"
)

func main() {
	ginEngin := gin.Default()
	DBAdapter := databases.NewSQLAdapter()
	ginEngin.Use(databases.DBMiddleware(DBAdapter))

	// Routes
	v1 := ginEngin.Group("/api/v1")
	{
		eg := v1.Group("/users")
		{
			eg.GET("/", controllers.AllUsers)
			eg.GET("/:id", controllers.FindUser)
			eg.POST("/", controllers.CreateUser)
			eg.PATCH("/:id", controllers.UpdateUser)
			eg.DELETE("/:id", controllers.DeleteUser)
		}
	}

	// Run the server
	ginEngin.Run(":8086")
}
