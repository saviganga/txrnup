package xuser

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Xuser struct {
	Id          uuid.UUID `gorm:"type:uuid;primaryKey;not null" json:"id"`
	Email       string    `gorm:"type:varchar(50);not null;unique" json:"email"`
	Password    string    `gorm:"type:varchar(128);not null" json:"password"`
	FirstName   string    `gorm:"type:varchar(50);not null" json:"first_name"`
	LastName    string    `gorm:"type:varchar(50);not null" json:"last_name"`
	PhoneNumber string    `gorm:"type:varchar(15)" json:"phone_number"`
	IsActive    bool      `gorm:"type:boolean;default:true" json:"is_active"`
	IsVerified  bool      `gorm:"type:boolean;default:false" json:"is_verified"`
	LastLogin   time.Time `gorm:"type:timestamp with time zone" json:"last_login"`
	CreatedAt   time.Time `gorm:"type:timestamp with time zone;default:now()" json:"created_at"`
	UpdatedAt   time.Time `gorm:"type:timestamp with time zone" json:"updated_at"`
}

func (user *Xuser) BeforeCreate(*gorm.DB) (err error) {
	if len(user.Email) == 0 {
		return errors.New("email cannot be empty")
	}
	if len(user.Password) == 0 {
		return errors.New("password cannot be empty")
	}
	if len(user.FirstName) == 0 {
		return errors.New("first name cannot be empty")
	}
	if len(user.LastName) == 0 {
		return errors.New("last name cannot be empty")
	}
	pass := []byte(user.Password)
	hash, err := bcrypt.GenerateFromPassword(pass, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	user.Id = uuid.New()
	user.Password = string(hash)
	return
}

// func (u *Xuser) ValidateNotNullFieldsBeforeCreate(tx *gorm.DB) (err error) {
// 	if len(u.Email) == 0 {
// 		return errors.New("email cannot be empty")
// 	}
// 	if len(u.Password) == 0 {
// 		return errors.New("password cannot be empty")
// 	}
// 	if len(u.FirstName) == 0 {
// 		return errors.New("first name cannot be empty")
// 	}
// 	if len(u.LastName) == 0 {
// 		return errors.New("last name cannot be empty")
// 	}
// 	return
// }
