package models

type Post struct {
	ID      uint `gorm:"primary_key"`
	Title   string
	Slug    string
	Content string `gorm:"type:text"`
}
