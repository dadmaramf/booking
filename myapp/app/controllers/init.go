package controllers

import (
	"myapp/app/models"

	gorm "github.com/revel/modules/orm/gorm/app"
	"github.com/revel/revel"
)

func initializeDB() {
	gorm.DB.AutoMigrate(&models.User{})
	gorm.DB.AutoMigrate(&models.Goods{})
	gorm.DB.AutoMigrate(&models.Booking{})
}

func init() {
	revel.OnAppStart(initializeDB)
}
