package schema

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Email    string `json:"email" gorm:"unique_index"`
	Password string `json:"password"`
	Username string `json:"username" gorm:"unique_index"`
}
