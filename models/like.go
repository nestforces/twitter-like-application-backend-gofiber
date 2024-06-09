package models

import "time"

type Like struct {
    UserID   int       `json:"user_id"`
    TweetID  int       `json:"tweet_id"`
    LikedAt  time.Time `json:"liked_at"`
}