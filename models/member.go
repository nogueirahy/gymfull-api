package models

import (
	"time"

	"gorm.io/gorm"
)

type Member struct {
	ID                    uint    `json:"memberId" gorm:"primaryKey"`
	FirstName             string  `json:"firstName"`
	LastName              string  `json:"lastName"`
	MobileNumber          string  `json:"mobileNumber"`
	Email                 string  `json:"email"`
	Address               string  `json:"address"`
	BirthDate             string  `json:"birthDate"`
	MedicalHistory        string  `json:"medicalHistory"`
	SubscriptionType      string  `json:"subscriptionType"`
	SubscriptionStartDate string  `json:"subscriptionStartDate"`
	SubscriptionEndDate   string  `json:"subscriptionEndDate"`
	SubscriptionPeriod    string  `json:"subscriptionPeriod"`
	Amount                float64 `json:"amount"`
	DayToPay              string  `json:"dayToPay"`
	PaymentStatus         string  `json:"paymentStatus"`
	DayPeriod             string  `json:"dayPeriod"`
	WorkoutStatus         bool    `json:"workoutStatus"`

	CreatedAt time.Time      `gorm:"<-:create" json:"created"`
	UpdatedAt time.Time      `json:"updated"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted"`
}
