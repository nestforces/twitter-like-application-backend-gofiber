package models

import "time"

type Reply struct {
    ID        int       `json:"id"`
    UserID    int       `json:"user_id"`
    TweetID   int       `json:"tweet_id"`
    Content   string    `json:"content"`
    CreatedAt time.Time `json:"created_at"`
}
