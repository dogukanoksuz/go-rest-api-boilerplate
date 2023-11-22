package models

type Post struct {
	Base
	Title   string `json:"title"`
	Content string `gorm:"type:text" json:"content"`
}
