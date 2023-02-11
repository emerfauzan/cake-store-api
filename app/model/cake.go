package model

import "time"

type Cake struct {
	Id            uint      `json:"id"`
	Title         string    `json:"title"`
	Description   string    `json:"description"`
	Rating        float32   `json:"rating"`
	ImageUrl      string    `json:"image_url"`
	IsDeletedFlag bool      `json:"is_deleted_flag"`
	CreatedBy     uint      `json:"created_by"`
	UpdatedBy     uint      `json:"updated_by"`
	DeletedBy     uint      `json:"deleted_by"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	DeletedAt     time.Time `json:"deleted_at"`
}
