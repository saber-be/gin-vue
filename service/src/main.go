package main

import (
	"qgen/challange/controllers"
	"qgen/challange/databases"
	"qgen/challange/middlewares"
	"qgen/challange/seeds"

	"github.com/gin-gonic/gin"
)

func main() {
	ginEngin := gin.Default()
	ginEngin.Use(middlewares.LoggerToFile(), gin.Recovery())
	ginEngin.Use(middlewares.CORSMiddleware())

	DBAdapter := databases.NewSQLAdapter()
	seeds.SeedUsers(DBAdapter)

	ginEngin.Use(middlewares.DBMiddleware(DBAdapter))

	// Routes
	v1 := ginEngin.Group("/api/v1")
	{
		eg := v1.Group("/users")
		{
			eg.GET("/", controllers.GetUsers)
			eg.GET("/:id", controllers.GetUserByID)
			eg.POST("/", controllers.CreateUser)
			eg.PATCH("/:id", controllers.UpdateUser)
			eg.DELETE("/:id", controllers.DeleteUser)
		}
	}

	// Run the server
	ginEngin.Run(":8080")
}
