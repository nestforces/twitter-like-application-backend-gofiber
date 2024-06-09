package models

import "time"

type User struct {
    UserID       int       `json:"user_id"`
    Username     string    `json:"username"`
    Email        string    `json:"email"`
    PasswordHash string    `json:"-"`
    Name         string    `json:"name"`
    Bio          string    `json:"bio"`
    Location     string    `json:"location"`
    Website      string    `json:"website"`
    CreatedAt    time.Time `json:"created_at"`
}