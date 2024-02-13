package controllers

import (
	"myapp/app/models"
	"time"

	"github.com/revel/revel"
)

type BookingApp struct {
	App
}

func (c BookingApp) ShowBookingForm() revel.Result {
	user := c.connected()
	if user == nil {
		return c.Render(App.Index)
	}
	var bookings []models.Booking
	if err := c.Txn.Find(&bookings, "user_id = ?", user.ID).Error; err != nil {
		c.Log.Error("Failed to get user bookings", "error", err)
	}
	return c.Render(user)
}

func (c BookingApp) BookRoom(userID int64, roomType, checkIn, checkOut string) revel.Result {
	checkInDate, _ := time.Parse("2006-01-02", checkIn)
	checkOutDate, _ := time.Parse("2006-01-02", checkOut)

	booking := models.Booking{
		UserID:   userID,
		Roomtype: roomType,
		CheckIn:  checkInDate,
		CheckOut: checkOutDate,
	}
	if err := c.Txn.Create(&booking); err != nil {
		c.Log.Error("Failed to book room", "error", err)
	}
	return c.RenderText("Room booked successfully!")
}
