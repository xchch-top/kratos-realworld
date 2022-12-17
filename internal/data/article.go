package data

import "gorm.io/gorm"

type Article struct {
	gorm.Model
	Slug           string `gorm:"size:200"`
	Title          string `gorm:"size:200"`
	Description    string `gorm:"size:200"`
	Body           string
	Tags           []Tag `gorm:"many2many:article_tags;"`
	AuthorID       uint
	Author         User
	FavoritesCount uint32
}

type Tag struct {
	gorm.Model
	Name     string    `gorm:"size:200;uniqueIndex"`
	Articles []Article `gorm:"many2many:article_tags;"`
}
