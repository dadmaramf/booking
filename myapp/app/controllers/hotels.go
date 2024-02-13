package controllers

import (
	"myapp/app/models"

	"github.com/revel/revel"
)

type Hotels struct {
	App
}

func (c Hotels) Index() revel.Result {
	user := c.connected()
	var bookings []models.Booking
	if err := c.Txn.Find(&bookings, "user_id = ?", user.ID).Error; err != nil {
		c.Log.Error("Failed to get user bookings", "error", err)
	}

	return c.Render(user, bookings)
}

func (c Hotels) Settings() revel.Result {
	return c.Render()
}
