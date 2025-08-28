package models

type User struct {
	Username string `json:"Username" gorm:"primaryKey"`
	Password string `json:"Password"`
	Email    string `json:"Email"`
	Phone    string `json:"Phone"`
	Name     string `json:"name"`
}
