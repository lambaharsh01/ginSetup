package models

type Users struct{
	Id uint `json:"id" gorm:"primarykey;autoIncrement"`
	FirstName string `json:"firstName" gorm"type:varchar(100);"`
	LastName string `json:"lastName" gorm"type:varchar(100);"`
	// Email string `json:"email"; gorm:"unique" gorm"type:varchar(100);"`
}