package models

import (
    "time"
)

type (

    OrderItem struct {
        ID          	uint        `gorm:"primary_key" json:"id" gorm:"column:id"`
        Quantity		uint      	`json:"quantity" gorm:"column:quantity"`
        SubTotal		uint      	`json:"sub_total" gorm:"column:sub_total"`
        OrderID		    uint      	`json:"order_id" gorm:"column:order_id"`
        ProductID		uint      	`json:"product_id" gorm:"column:product_id"`
        CreatedAt		time.Time   `json:"created_at" gorm:"column:created_at"`
        UpdatedAt   	time.Time   `json:"updated_at" gorm:"column:updated_at"`
        // Order    		Order   	`json:"-" gorm:"foreignKey:OrderID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"` 
        Product    		Product     `json:"-" gorm:"foreignKey:ProductID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
    }
)