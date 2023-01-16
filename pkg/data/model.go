package data

import (
	"gorm.io/gorm"
	"time"
)

type Model struct {
	Id        uint64 `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func NewCreateModel() *Model {
	return &Model{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func NewUpdateModel(id uint64) *Model {
	return &Model{
		Id:        id,
		UpdatedAt: time.Now(),
	}
}
