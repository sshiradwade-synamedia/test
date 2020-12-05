package models

import (
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
  ID uint32 `gorm:"primary_key;auto_increment" json:"id"`
  UserName string `gorm:"size:255;not null;unique" json:"username"`
  Email string `gorm:"size:100;not null;unique" json:"email"`
  Password string `gorm:"size:100;not null;unique" json:"password"`
  CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"create_at"`
  UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
  DeleteAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"deleted_at"`
  Role uint32 `gorm:"not null" json:"role"`
  Address
}

type Address struct {
  Line1 string `gorm:"size:100;not null" json:"line1"`
  City string `gorm:"size:100;not nul" json:"city"`
  State string `gorm:"size:60; not null" json:"state"`
  LatLong string `gorm:"size:100" json:"latlong"`
}

// Hash will has the password
func Hash (password string)([]byte, error) {
  return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}


// BeforeSave will convert hash into string for saving in DB
func (u *User) BeforeSave()  error{
  hashedPassword, err := Hash(u.Password)
  if err != nil {
    return err
  }
  u.Password = string(hashedPassword)
  return nil
}

// Prepare -
func (u *User) Prepare() {
	u.ID = 0
	u.UserName = html.EscapeString(strings.TrimSpace(u.UserName))
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
}

// SaveUser will save user details to the DB
func (u *User) SaveUser(db *gorm.DB) (*User, error )  {
  var err error

  err = db.Debug().Create(&u).Error
  if err != nil {
    return &User{}, err
  }
  return u,nil
}


