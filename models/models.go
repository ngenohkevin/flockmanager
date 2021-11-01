package models

import "github.com/jinzhu/gorm"


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

func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&Kuroiler{}, &RainbowRooster{}, &Broilers{}, &Layers{})
	return db
}
