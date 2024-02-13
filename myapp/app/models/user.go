package models

import (
	"fmt"
	"regexp"

	"github.com/jinzhu/gorm"
	"github.com/revel/revel"
	"golang.org/x/crypto/bcrypt"
)

// User model.
type User struct {
	gorm.Model
	Name           string `gorm:"size:255"`
	Email          string `gorm:"type:varchar(100);unique_index"`
	Username       string `gorm:"size:255"`
	HashedPassword []byte
	Active         bool
	FileName       string `gorm:"size:255"`
	VerifiCode	   string `gorm:"size:255"`
}

// SetNewPassword set a new hashsed password to user.
func (u *User) SetNewPassword(passwordString string) {
	bcryptPassword, _ := bcrypt.GenerateFromPassword([]byte(passwordString), bcrypt.DefaultCost)
	u.HashedPassword = bcryptPassword
}

// Validate filds
func (u *User) String() string {
	return fmt.Sprintf("User(%s)", u.Username)
}

var userRegex = regexp.MustCompile("^\\w*$")
var emailRegex = regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)

func (u *User) Validate(v *revel.Validation) {
	v.Check(u.Username,
		revel.Required{},
		revel.MaxSize{15},
		revel.MinSize{4},
		revel.Match{userRegex},
	)
	v.Check(u.Email,
		revel.Match{emailRegex},
	)
}
