package models

import "time"

type Retweet struct {
    UserID     int       `json:"user_id"`
    TweetID    int       `json:"tweet_id"`
    RetweetedAt time.Time `json:"retweeted_at"`
}