package model

type User struct {
	Id       string `gorm:"type:uuid;primaryKey"`
	Name     string
	Email    string
	Password string
	Status   string
}
