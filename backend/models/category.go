package models

import (
	"time"
)

// Category is car product categories
type Category struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title" xorm:"varchar(255) not null" binding:"required"`
	Description string    `json:"description" xorm:"varchar(255) not null" binding:"required"`
	ImageUrl    string    `json:"imageUrl" xorm:"varchar(255) not null" binding:"required"`
	CreatedAt   time.Time `json:"createdAt" xorm:"created"`
	UpdatedAt   time.Time `json:"-" xorm:"updated"`
	DeletedAt   time.Time `json:"-" xorm:"deleted"` // soft delete, shows delete time instead of really deleting
}

func (c *Category) TableName() string {
	return "category"
}
