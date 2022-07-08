package main

import (
	"qgen/challange/controllers"
	"qgen/challange/databases"
	"qgen/challange/middlewares"

	"github.com/gin-gonic/gin"
)

func main() {
	ginEngin := gin.Default()
	DBAdapter := databases.NewSQLAdapter()
	ginEngin.Use(middlewares.CORSMiddleware())
	ginEngin.Use(middlewares.DBMiddleware(DBAdapter))

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
