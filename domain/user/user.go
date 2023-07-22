package entities

import "time"
 
func (User) TableName() string { return "users" }
type User struct { 
	Id int `json:"id,primary_key"` 
	UserName  string `gorm:"size:255"` 
	Email  string `gorm:"size:255"` 
	Password  string `gorm:"size:255"` 
	PhoneNumber  string `gorm:"size:255"` 
	Avatar  string `gorm:"size:255"` 
	RoleId int
	Status bool
	Gender bool
	Dob time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}