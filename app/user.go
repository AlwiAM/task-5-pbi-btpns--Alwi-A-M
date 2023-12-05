package app

import "time"

type User struct {
    ID        uint      `json:"id" gorm:"primary_key"`
    Username  string    `json:"username" gorm:"type:varchar(100);not null"`
    Email     string    `json:"email" gorm:"type:varchar(100);unique;not null"`
    Password  string    `json:"-" gorm:"type:varchar(100);not null"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}
