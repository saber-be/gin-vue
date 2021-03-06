package models

type User struct {
	ID    uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Name  string `json:"name"`
	Email string `json:"email" gorm:"type:varchar(100);unique"`
	Phone string `json:"phone"`
	Age   uint   `json:"age"`
}

type UserInput struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required,email"`
	Phone string `json:"phone" binding:"required"`
	Age   uint   `json:"age" binding:"required,min=18"`
}
