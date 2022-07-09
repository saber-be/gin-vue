package databases

import (
	"qgen/challange/models"

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
func (d *SQLDB) Connect() error {
	database, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(&models.User{})
	d.db = database
	return nil
}

// Close the database connection
func (d *SQLDB) Close() error {
	return nil
}

// Create a new record in the database
func (d *SQLDB) Create(model interface{}) error {
	return d.db.Create(model).Error
}

// Read a record from the database
func (d *SQLDB) Find(models_list interface{}, conds ...interface{}) error {
	return d.db.Find(models_list, conds).Error
}

// Update a record in the database
func (d *SQLDB) Update(model interface{}, values interface{}) error {
	return d.db.Model(model).Updates(values).Error
}

// Delete a record from the database
func (d *SQLDB) Delete(model interface{}) error {
	return d.db.Delete(model).Error
}
