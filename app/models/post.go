package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Post struct {
	ID      string `gorm:"primary_key"`
	Title   string
	Content string `gorm:"type:text"`
}

func (post *Post) BeforeCreate(tx *gorm.DB) error {
	post.ID = uuid.NewString()
	return nil
}
