package main

import (
	"fmt"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Content  string
	Author   string
	Comments []Comment
}

type Comment struct {
	gorm.Model
	Content string
	Author  string
	PostID  int
}

var Db *gorm.DB

// connect to the Db
func init() {
	var err error
	dsn := "host=localhost user=gwp dbname=gwp password=gwp sslmode=disable"
	Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	Db.AutoMigrate(&Post{}, &Comment{})
}

func main() {
	// Create a post
	post := Post{Content: "Hello World!", Author: "Sau Sheong"}
	Db.Create(&post)

	// Add a comment
	comment := Comment{Content: "Good post!", Author: "Joe"}
	Db.Model(&post).Association("Comments").Append(&comment)

	// Get all comments from a post
	var readPost Post
	Db.Where("author = $1", "Sau Sheong").First(&readPost)
	var comments []Comment
	Db.Model(&readPost).Association("Comments").Find(&comments)

	// show the first comment
	fmt.Println(comments[0])
}
