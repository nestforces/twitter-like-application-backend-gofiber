package models

import "time"

type Tweet struct {
    TweetID   int       `json:"tweet_id"`
    UserID    int       `json:"user_id"`
    Content   string    `json:"content"`
    CreatedAt time.Time `json:"created_at"`
}