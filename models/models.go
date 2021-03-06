package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	ID       uint
	Name     string
	Email    string `gorm:"unique"`
	Password []byte
}

// Kuroiler structs

type Kuroiler struct {
	gorm.Model

	Title      string     `gorm:"unique" json:"title"`
	Production Production `gorm:"embedded"`
	Hatchery   Hatchery   `gorm:"embedded"`
	Premises   Premises   `gorm:"embedded"`
}
type RainbowRooster struct {
	gorm.Model

	Title      string     `gorm:"unique" json:"title"`
	Production Production `gorm:"embedded"`
	Hatchery   Hatchery   `gorm:"embedded"`
	Premises   Premises   `gorm:"embedded"`
}
type Broilers struct {
	gorm.Model

	Title      string     `gorm:"unique" json:"title"`
	Production Production `gorm:"embedded"`
	Hatchery   Hatchery   `gorm:"embedded"`
	Premises   Premises   `gorm:"embedded"`
}
type Layers struct {
	gorm.Model

	Title      string     `gorm:"unique" json:"title"`
	Production Production `gorm:"embedded"`
	Hatchery   Hatchery   `gorm:"embedded"`
	Premises   Premises   `gorm:"embedded"`
}

type Premises struct {
	Farm  string `json:"farm"`
	House string `json:"house"`
}

type Production struct {
	Eggs         int `json:"eggs"`
	Dirty        int `json:"dirty"`
	WrongShape   int `json:"wrong_shape"`
	WeakShell    int `json:"weak_shell"`
	Damaged      int `json:"damaged"`
	HatchingEggs int `json:"hatching_eggs"`
}

type Hatchery struct {
	Infertile   int `json:"infertile"`
	Early       int `json:"early"`
	Middle      int `json:"middle"`
	Late        int `json:"late"`
	DeadChicks  int `json:"dead_chicks"`
	AliveChicks int `json:"alive_chicks"`
}
type Flock struct {
	Id        int    `gorm:"unique" json:"id"`
	FlockName string `json:"flock_name"`
	Operations	  Operations `gorm:"embedded" json:"operations"`
}
type Operations struct {
	OpPro	string	`json:"op_pro"`
	OpHat	string	`json:"op_hat"`
}

func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&User{}, &Kuroiler{}, &RainbowRooster{}, &Broilers{}, &Layers{}, &Flock{})
	return db
}
