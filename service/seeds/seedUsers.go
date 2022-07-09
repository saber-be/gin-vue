package seeds

import (
	"qgen/challange/databases"
	"qgen/challange/models"

	"github.com/bxcodec/faker/v3"
)

func SeedUsers(d databases.IDBAdapter) {
	var users []models.User
	d.Connect()
	d.Find(&users, nil)
	if len(users) > 0 {
		return
	}

	for i := 0; i < 100; i++ {
		user := models.User{
			Name:  faker.Name(),
			Email: faker.Email(),
			Phone: faker.Phonenumber(),
		}
		d.Create(&user)
	}
}
