package databases

import (
	"qgen/challange/models"
)

type MockDB struct {
}

func NewMOCKAdapter() *MockDB {
	return &MockDB{}
}

// Connect to the database
func (d *MockDB) Connect() error {
	return nil
}

// Close the database connection
func (d *MockDB) Close() error {
	return nil
}

// Create a new record in the database
func (d *MockDB) Create(model interface{}) error {
	return nil
}

// Read a record from the database
func (d *MockDB) Find(models_list interface{}, pagination *models.Pagination, conds ...interface{}) error {
	return nil
}

// Update a record in the database
func (d *MockDB) Update(model interface{}, values interface{}) error {
	return nil
}

// Delete a record from the database
func (d *MockDB) Delete(model interface{}) error {
	return nil
}
