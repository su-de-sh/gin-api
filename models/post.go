package models

// models/post.go

import "gorm.io/gorm"

type Post struct {
    gorm.Model
    Title   string `json:"title" binding:"required"`
    Content string `json:"content" binding:"required"`
}