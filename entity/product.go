package entity

import (
	"gorm.io/gorm"
	"time"
)

type Product struct {
	gorm.Model
	ID string `json:"id" xml:"id" form:"id" binding:"required,min=5,max=7" gorm:"primary_key;type:varchar(7)"`
	Name string `json:"name" xml:"name" form:"name" binding:"required" gorm:"type:varchar(100)"`
	Color string `json:"color" xml:"color" form:"color" binding:"required" gorm:"type:varchar(20)"`
	Description string `json:"description" xml:"description" form:"description" gorm:"type:varchar(256)"`
	SalePrice float64 `json:"sale-price" xml:"sale-price" form:"sale-price" binding:"required,gte=0,lte=999999999"`
	OrderCost float64 `json:"order-cost" xml:"order-cost" form:"order-cost" binding:"required,gte=0,lte=999999999"`
	Quantity int16 `json:"quantity" xml:"quantity" form:"quantity" binding:"required"`
	Type string `json:"type" xml:"type" form:"type" binding:"required" gorm:"type:varchar(20)"`
	Image string `json:"image" xml:"image" form:"type" binding:"url" gorm:"type:varchar(255)"`
	CreatedAt time.Time `json:"-" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `json:"-" gorm:"default:CURRENT_TIMESTAMP"`
	BrandID string `json:"brand-id" xml:"brand-id" form:"brand-id" binding:"required,min=5,max=7" gorm:"primary_key;type:varchar(7)"`
}