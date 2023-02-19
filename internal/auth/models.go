package auth

import "time"

type User struct {
    ID        uint   `gorm:"primaryKey"`
    Username  string `gorm:"unique"`
    Password  string
    Role      UserRole
    CreatedAt time.Time
    UpdatedAt time.Time
}

type Item struct {
    ID          uint   `gorm:"primaryKey"`
    Name        string `gorm:"unique"`
    Description string
    Price       uint
    Quantity    uint
    CreatedAt   time.Time
    UpdatedAt   time.Time
}

type Order struct {
    ID        uint      `gorm:"primaryKey"`
    UserID    uint      `gorm:"index"`
    ItemID    uint      `gorm:"index"`
    Quantity  uint      `gorm:"not null"`
    TotalCost uint      `gorm:"not null"`
    CreatedAt time.Time `gorm:"index"`
    UpdatedAt time.Time
}

type UserRole string

const (
    RoleAdmin UserRole = "admin"
    RoleUser  UserRole = "user"
)
