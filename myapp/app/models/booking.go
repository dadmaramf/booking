// Ваш файл models/booking.go

package models

import "time"

type Booking struct {
	ID        int64     `gorm:"primaryKey"`
	UserID    int64     // Идентификатор пользователя, связанный с бронированием
	Roomtype  string    // Идентификатор номера, связанного с бронированием
	CheckIn   time.Time // Дата заезда
	CheckOut  time.Time // Дата выезда
	CreatedAt time.Time // Дата создания записи
}
