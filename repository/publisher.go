package repository

import "gorm.io/gorm"

type Publisher struct {
	gorm.Model
	Name    string `db:"name"`
	Address string `db:"address"`
	Phone   string `db:"phone"`
	Email   string `db:"email"`
}