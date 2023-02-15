package entity

import (
	"gorm.io/gorm"
)

type Employee struct {
	gorm.Model
	Name       string `valid:"null~name not blank"`
	Email      string
	EmployeeID string `valid:"matches(^[JMS]\\d{8}$)"`
}
