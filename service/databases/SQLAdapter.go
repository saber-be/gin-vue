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
	database, err := gorm.Open(sqlite.Open("users.db"), &gorm.Config{})

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
func (d *SQLDB) Find(models_list interface{}, pagination *models.Pagination, conds ...interface{}) error {
	if pagination != nil {
		p := *pagination
		offset := (p.Page - 1) * p.Limit
		queryBuilder := d.db.Model(models_list)
		if len(conds) > 1 {
			queryBuilder = queryBuilder.Where(conds[0], conds[1:]...)
		}
		queryBuilder.Count(&pagination.Total)
		queryBuilder.Limit(p.Limit).Offset(offset).Order(p.Sort)
		return queryBuilder.Find(models_list).Error
	}

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
