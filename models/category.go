package models

import (
    "time"
)

type (
    // AgeRatingCategory
    Category struct {
        ID          uint        `gorm:"primary_key" json:"id" gorm:"column:id"`
        Name        string      `json:"name" gorm:"column:name"`
        CreatedAt   time.Time   `json:"created_at" gorm:"column:created_at"`
        UpdatedAt   time.Time   `json:"updated_at" gorm:"column:updated_at"`
        Products    []Product   `json:"-"`
    }
)