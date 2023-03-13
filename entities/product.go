package entities

import "time"

type Product struct {
	ID               uint64 `gorm:"primary_key;auto_increment" json:"id"`
	Name             string `gorm:"type:varchar(255);unique" json:"name" binding:"required,max=255"`
	Description      string `json:"description" binding:"max=65535"`
	Price            float32 `json:"price" binding:"required"`
	Weight           float32 `json:"weight" binding:"required"`
	Ingredients      string `json:"ingredients" binding:"max=65535" validate:"should-not-contain-methylparaben"`

	CreatedAt time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `json:"-" gorm:"default:CURRENT_TIMESTAMP"`
}
