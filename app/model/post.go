package model

import (
	"errors"
	"github.com/jinzhu/gorm"
)

type Post struct {
	gorm.Model
	Title string `json:"title" form:"title"`
	Body  string `json:"body" form:"body"`
}

func FindPost(db *gorm.DB, id int) (*Post, error) {
	var post Post

	db.First(&post, id)
	if post.ID == 0 {
		return nil, errors.New("Not found Post")
	}

	return &post, nil
}

func AllPost(db *gorm.DB) []Post {
	var posts []Post

	db.Limit(10).Find(&posts)

	return posts
}

func (p Post) Create(db *gorm.DB) {
	db.Create(&p)
}

func (p Post) Update(db *gorm.DB, post *Post) {
	db.Model(&p).Updates(post)
}

func (p Post) Delete(db *gorm.DB) {
	db.Delete(&p)
}
