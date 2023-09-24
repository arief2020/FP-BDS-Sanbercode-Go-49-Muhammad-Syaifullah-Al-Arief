package models

import (
	"time"

)

type (
    Product struct {
        ID                  uint				`json:"id" gorm:"primary_key" gorm:"column:id"`
        Name				string				`json:"name" gorm:"column:name"`
        Description			string				`json:"description" gorm:"column:description"`
        Price				uint				`json:"price" gorm:"column:price"`
        Stock				uint				`json:"stock" gorm:"column:stock"`
        CategoryID			uint			    `json:"category_id" gorm:"column:category_id"`
        CreatedAt           time.Time			`json:"created_at" gorm:"column:created_at"`
        UpdatedAt           time.Time			`json:"updated_at" gorm:"column:updated_at"`
        Category			Category			`json:"-" gorm:"foreignKey:CategoryID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
        OrderItem			[]OrderItem			`json:"-"`
        ReviewProduct		[]ReviewProduct		`json:"-"`
    }
)