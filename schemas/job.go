package schemas

import (
	"gorm.io/gorm"
	"time"
)

type Job struct {
	gorm.Model
	Role        string
	Company     string
	Location    string
	Description string
	Remote      bool
	Link        string
	Salary      int64
}

type JobResponse struct {
	Id          uint      `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	ExpiredAt   time.Time `json:"expired_at"`
	DeletedAt   time.Time `json:"deletedAt,omitempty"`
	Role        string    `json:"role"`
	Company     string    `json:"company"`
	Location    string    `json:"location"`
	Description string    `json:"description"`
	Remote      bool      `json:"isRemote"`
	Link        string    `json:"link"`
	Salary      int64     `json:"salary"`
}
