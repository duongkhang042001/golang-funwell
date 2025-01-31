package user

import "time"

type Entity struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"size:255;not null"`
	Email     string `gorm:"size:255;unique;not null"`
	Password  string `gorm:"size:255;not null"`
	Role      string `gorm:"size:50;not null;default:'user'"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
