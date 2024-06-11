package models

import "time"

type DirectMessage struct {
    ID        int       `json:"id"`
    SenderID  int       `json:"sender_id"`
    ReceiverID int      `json:"receiver_id"`
    Content   string    `json:"content"`
    CreatedAt time.Time `json:"created_at"`
}
