package models

import (
    "gorm.io/gorm"
)

var DB *gorm.DB

type Task struct {
    gorm.Model
    Title  string `json:"title"`
    Status string `json:"status"`
}

type InputTask struct {
    Title  string `json:"title" binding:"required"`
    Status string `json:"status" binding:"required"`
}
