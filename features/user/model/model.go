package model

import (
	"time"
)

type Users struct {
	ID         string    `gorm:"type:varchar(50);primaryKey;not null" json:"id"`
	Name       string    `gorm:"varchar(50);not null" json:"username"`
	Image      string    `json:"image"`
	Email      string    `gorm:"varchar(50);not null" json:"email"`
	Password   string    `gorm:"varchar(50);not null" json:"password"`
	Role       string    `gorm:"Varchar(25);not null" json:"role"`
	Point      string    `gorm:"Varchar(100);not null" json:"point"`
	TotalPoint string    `gorm:"Varchar(100);not null" json:"total_point"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"update_at"`
}

type UserPoint struct {
	Id        string `gorm:"type:varchar(50);primaryKey;not null" json:"id"`
	UserId    string
	Type      string
	TaskName  string
	Point     int
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"update_at"`
}
