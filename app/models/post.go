package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Post struct {
	ID      string `gorm:"primary_key" json:"id"`
	Title   string `json:"title"`
	Content string `gorm:"type:text" json:"content"`
}

func (post *Post) BeforeCreate(tx *gorm.DB) error {
	post.ID = uuid.NewString()
	return nil
}
