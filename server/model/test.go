package model

import "time"

type Test struct {
	Id        uint `gorm:"primary_key" json:"id"`
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index" json:"-"`
}
