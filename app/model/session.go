package model

import "time"

type Session struct {
	Id         uint       `json:"id"`
	UserId     uint       `json:"user_id"`
	SessionKey string     `json:"session_key"`
	ExpiresAt  *time.Time `json:"expires_at"`
}
