// main.go
package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB
var err error

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Email    string `json:"email" gorm:"unique_index"`
	Password string `json:"password"`
	Username string `json:"username" gorm:"unique_index"`
}

func main() {
	db, err = gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/go_base_db?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("Không thể kết nối cơ sở dữ liệu")
	}
	defer db.Close()
	db.AutoMigrate(&User{})

	r := gin.Default()
	r.GET("/users", getUsers)
	r.POST("/users", addUser)
	r.PUT("/users/:id", updateUser)
	r.DELETE("/users/:id", deleteUser)
	r.Run(":8080")
}

func getUsers(c *gin.Context) {
	var users []User
	db.Find(&users)
	c.JSON(200, users)
}

func addUser(c *gin.Context) {
	var user User
	c.BindJSON(&user)
	db.Create(&user)
	c.JSON(200, user)
}

func updateUser(c *gin.Context) {
	var user User
	id := c.Params.ByName("id")
	if err := db.Where("id = ?", id).First(&user).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&user)
	db.Save(&user)
	c.JSON(200, user)
}

func deleteUser(c *gin.Context) {
	id := c.Params.ByName("id")
	var user User
	d := db.Where("id = ?", id).Delete(&user)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "đã bị xóa"})
}
