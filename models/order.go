package models

import (
    "time"
)

type (

    Order struct {
        ID          	uint        `gorm:"primary_key" json:"id" gorm:"column:id"`
        Status			string      `json:"status" gorm:"column:status"`
        TotalPrice		uint      	`json:"total_price" gorm:"column:total_price"`
        UserID			uint      	`json:"user_id" gorm:"column:user_id"`
        CreatedAt		time.Time   `json:"created_at" gorm:"column:created_at"`
        UpdatedAt   	time.Time   `json:"updated_at" gorm:"column:updated_at"`
        User    		User   		`json:"-" gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
        OrderItem    	[]OrderItem `json:"-"`
    }
)