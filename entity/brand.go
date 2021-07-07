package entity

import (
	"gorm.io/gorm"
	"time"
)

type Brand struct {
	gorm.Model
	BrandID string `json:"brand-id" xml:"brand-id" form:"brand-id" binding:"required,min=5,max=7" gorm:"primary_key;type:varchar(7)"`
	BrandName string `json:"brand-name" xml:"brand-name" form:"brand-name" binding:"required" gorm:"type:varchar(100)"`
	BrandLogo string `json:"brand-logo" xml:"brand-logo" form:"brand-logo" binding:"url" gorm:"type:varchar(256)"`
	CreatedAt time.Time `json:"-" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `json:"-" gorm:"default:CURRENT_TIMESTAMP"`
}