package databases

import (
	"github.com/gin-gonic/gin"
)

type IDBAdapter interface {
	Connect(c *gin.Context) error
	Close(c *gin.Context) error
	Create(c *gin.Context, model interface{}) error
	Find(c *gin.Context, models interface{}, conds ...interface{}) error
	Update(c *gin.Context, model interface{}, values interface{}) error
	Delete(c *gin.Context, model interface{}) error
}

func DBMiddleware(d IDBAdapter) gin.HandlerFunc {
	return func(c *gin.Context) {
		d.Connect(c)
		defer d.Close(c)
		c.Next()
	}
}
