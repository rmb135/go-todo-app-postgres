package models

import "gorm.io/gorm"

type Todo struct {
    gorm.Model
    Title  string `json:"title"`
    Done   bool   `json:"done"`
    UserID uint   `json:"user_id"`
}
