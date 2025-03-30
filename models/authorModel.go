package models

// Gin has already integrated validator
import (
	"time"
)

type Author struct {
	ID              uint       `json:"id" db:"id"`
	FirstName       string     `json:"first_name" db:"first_name" binding:"required"`
	LastName        string     `json:"last_name" db:"last_name" binding:"required"`
	ProfilePicture  *string    `json:"profile_picture" db:"profile_picture"`
	Biography       *string    `json:"biography" db:"biography"`
	NationalityCode *string    `json:"nationality_code" db:"nationality_code"`
	Birthdate       *time.Time `json:"birthdate" db:"birthdate" validate:"lt"`
	CreatedTime     time.Time  `json:"created_time" db:"created_time"`
	UpdatedTime     time.Time  `json:"updated_time" db:"updated_time"`
}

// specify table name in DB
func (Author) TableName() string {
	return "author"
}
