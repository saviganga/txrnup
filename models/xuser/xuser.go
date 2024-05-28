package xuser

import (
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
	IsActive   bool      `gorm:"type:boolean;default:true" json:"is_active"`
	IsVerified bool      `gorm:"type:boolean;default:false" json:"is_verified"`
	LastLogin  time.Time `gorm:"type:timestamp with time zone" json:"last_login"`
	CreatedAt  time.Time `gorm:"type:timestamp with time zone;default:now()" json:"created_at"`
	UpdatedAt  time.Time `gorm:"type:timestamp with time zone" json:"updated_at"`
}

func (user *Xuser) BeforeCreate(*gorm.DB) (err error) {
	pass := []byte(user.Password)
	hash, err := bcrypt.GenerateFromPassword(pass, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	user.Id = uuid.New()
	user.Password = string(hash)
	return
}
