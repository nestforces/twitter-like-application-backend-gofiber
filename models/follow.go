package models

import "time"

type Follow struct {
    FollowerID int       `json:"follower_id"`
    FolloweeID int       `json:"followee_id"`
    FollowedAt time.Time `json:"followed_at"`
}