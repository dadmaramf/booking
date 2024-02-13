package controllers

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"myapp/app/models"
	"myapp/app/modules"

	gormc "github.com/revel/modules/orm/gorm/app/controllers"
	"github.com/revel/revel"
	"golang.org/x/crypto/bcrypt"
)

type App struct {
	gormc.TxnController
}

func (c App) AddUser() revel.Result {
	if user := c.connected(); user != nil {
		c.ViewArgs["user"] = user
	}
	return nil
}

func (c App) connected() *models.User {
	if c.ViewArgs["user"] != nil {
		return c.ViewArgs["user"].(*models.User)
	}
	if username, ok := c.Session["user"]; ok {
		return c.getUser(username.(string))
	}
	return nil
}

func (c App) getUser(username string) (user *models.User) {
	user = &models.User{}
	_, err := c.Session.GetInto("fulluser", user, false)
	if user.Username == username {
		return user
	}

	if c.Txn.Where("username = ?", username).First(user).RecordNotFound() {
		c.Log.Error("Failed to find user", "user", username, "error", err)

	}

	c.Session["fulluser"] = user
	return
}

func (c App) Login(username, password string, remember bool) revel.Result {
	user := c.getUser(username)
	if user != nil {
		err := bcrypt.CompareHashAndPassword(user.HashedPassword, []byte(password))
		if err == nil {
			c.Session["user"] = username
			if remember {
				c.Session.SetDefaultExpiration()
			} else {
				c.Session.SetNoExpiration()
			}
			c.Flash.Success("Welcome, " + username)
			return c.Redirect(Hotels.Index)
		}
	}

	c.Flash.Out["username"] = username
	c.Flash.Error("Login failed")
	return c.Redirect(App.Index)
}

func (c App) Logout() revel.Result {
	for k := range c.Session {
		delete(c.Session, k)
	}
	return c.Redirect(App.Index)
}

func (c App) Index() revel.Result {
	if c.connected() != nil {
		return c.Redirect(App.Register)
	}
	c.Flash.Error("Please log in first")
	return c.Render()
}

func (c App) Register() revel.Result {
	return c.Render()
}
func generateVerificationCode() (string, error) {
	bytes := make([]byte, 16)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(bytes), nil
}

func (c App) SaveUser(user models.User, verifyPassword string, userPassword string) revel.Result {
	c.Validation.Required(verifyPassword)
	c.Validation.Required(verifyPassword == userPassword).
		MessageKey("Password does not match")
	ValidatePassword(c.Validation, userPassword).
		Key("userPassword")
	user.Validate(c.Validation)
	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(App.Register)
	}
	user.FileName = "Null"
	user.Active = false
	user.SetNewPassword(userPassword)

	verificationCode, err := generateVerificationCode()
	if err != nil {
		fmt.Print(err)
	}
	user.VerifiCode = verificationCode
	emails := []string{user.Email}
	err = modules.UserMailer{}.SendReport(verificationCode, emails)
	if err != nil {
		fmt.Print(err)
	}
	err2 := c.Txn.Save(&user)
	if err2 != nil {
		fmt.Print(err2)
	}
	c.Flash.Success("Welcome, " + user.Name + "! You have received an email, please confirm your email")
	return c.Redirect(App.Index)
}

func ValidatePassword(v *revel.Validation, password string) *revel.ValidationResult {
	return v.Check(password,
		revel.Required{},
		revel.MaxSize{24},
		revel.MinSize{8},
	)
}
