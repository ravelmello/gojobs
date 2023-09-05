package schemas

import "gorm.io/gorm"

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
