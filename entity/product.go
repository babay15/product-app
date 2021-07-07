package entity

type Product struct {
	ID string `json:"id" xml:"id" form:"id" binding:"required,min=5,max=7"`
	Name string `json:"name" xml:"name" form:"name" binding:"required"`
	Color string `json:"color" xml:"color" form:"color" binding:"required"`
	Description string `json:"description" xml:"description" form:"description"`
	SalePrice float64 `json:"sale-price" xml:"sale-price" form:"sale-price" binding:"required,gte=0,lte=999999999"`
	OrderCost float64 `json:"order-cost" xml:"order-cost" form:"order-cost" binding:"required,gte=0,lte=999999999"`
	Quantity int16 `json:"quantity" xml:"quantity" form:"quantity" binding:"required"`
	Type string `json:"type" xml:"type" form:"type" binding:"required"`
	Image string `json:"image" xml:"image" form:"type" binding:"url"`
	Brand Brand `json:"brand" xml:"brand" form:"brand" binding:"required"`
}

type Brand struct {
	BrandID string `json:"brand-id" xml:"brand-id" form:"brand-id" binding:"required"`
	BrandName string `json:"brand-name" xml:"brand-name" form:"brand-name" binding:"required"`
	BrandLogo string `json:"brand-logo" xml:"brand-logo" form:"brand-logo" binding:"url"`
}