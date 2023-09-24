package models

import (
    "time"
)

type (

    ReviewProduct struct {
        ID          	uint        `gorm:"primary_key" json:"id" gorm:"column:id"`
        Rating			uint      	`json:"rating" gorm:"column:rating"`
        Comment			string      `json:"comment" gorm:"column:comment"`
        ProductID		uint      	`json:"product_id" gorm:"column:product_id"`
        UserID			uint      	`json:"user_id" gorm:"column:user_id"`
        CreatedAt		time.Time   `json:"created_at"`
        UpdatedAt   	time.Time   `json:"updated_at"`
        // Product    		Product   	`json:"-" gorm:"foreignKey:ProductID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" `
        // User    		User   		`json:"-" gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
    }
)