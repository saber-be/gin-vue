package databases

type NOSQLDB struct {
	//DB NOSQLDB_DB_instance
}

// Connect to the database
func (d *NOSQLDB) Connect() error {
	return nil
}

// Close the database connection
func (d *NOSQLDB) Close() error {
	return nil
}

// Create a new record in the database
func (d *NOSQLDB) Create(model interface{}) error {
	return nil
}

// Read a record from the database
func (d *NOSQLDB) Read(model interface{}) error {
	return nil
}

// Update a record in the database
func (d *NOSQLDB) Update(model interface{}) error {
	return nil
}

// Delete a record from the database
func (d *NOSQLDB) Delete(model interface{}) error {
	return nil
}
