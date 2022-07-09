package databases

import (
	"qgen/challange/models"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type SQLDB struct {
	db *gorm.DB
}

func NewSQLAdapter() *SQLDB {
	return &SQLDB{}
}

// Connect to the database
func (d *SQLDB) Connect(c *gin.Context) error {
	database, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(&models.User{})
	sqlDB := &SQLDB{db: database}
	c.Set("db", sqlDB)
	return nil
}

// Close the database connection
func (d *SQLDB) Close(c *gin.Context) error {
	return nil
}

// Create a new record in the database
func (d *SQLDB) Create(c *gin.Context, model interface{}) error {
	dba := c.MustGet("db").(*SQLDB)
	return dba.db.Create(model).Error
}

// Read a record from the database
func (d *SQLDB) Find(c *gin.Context, models_list interface{}, conds ...interface{}) error {
	dba := c.MustGet("db").(*SQLDB)
	return dba.db.Find(models_list, conds).Error
}

// Update a record in the database
func (d *SQLDB) Update(c *gin.Context, model interface{}, values interface{}) error {
	dba := c.MustGet("db").(*SQLDB)
	return dba.db.Model(model).Updates(values).Error
}

// Delete a record from the database
func (d *SQLDB) Delete(c *gin.Context, model interface{}) error {
	dba := c.MustGet("db").(*SQLDB)
	return dba.db.Delete(model).Error
}
