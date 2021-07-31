package repository

import (
	"ProgettoDB/models"
	"time"
)

type Nation struct {
	Id   uint
	Name string
}

type Team struct {
	Id      uint
	Name    string
	Nation  string
	Entries []Entry `gorm:"foreignKey:Team"`
}

type Manufacturer struct {
	Id     uint
	Name   string
	Nation string
	Cars   []Car `gorm:"foreignKey:Brand"`
}

type Layout struct {
	Id     uint
	Name   string
	Track  string
	Length uint
	Type   models.LayoutType
	Races  []Race `gorm:"foreignKey:Layout"`
}

type Track struct {
	Id       uint
	Name     string
	Nation   string
	Location string
	Layouts  []Layout `gorm:"foreignKey:Track"`
}

type Driver struct {
	Cf        string
	Name      string
	Surname   string
	Birthdate time.Time
	Nation    string
	Sex       models.Sex
	Entries   []Entry `gorm:"many2many:driver_entries;joinForeignKey:driver"`
}

type CarClass struct {
	Id            uint
	Name          string
	Cars          []Car          `gorm:"foreignKey:Class"`
	Championships []Championship `gorm:"many2many:championship_classes;joinForeignKey:class"`
}

type Car struct {
	Id           uint
	Model        string
	Year         uint
	Brand        string
	Class        string
	Drivetrain   models.Drivetrain
	Transmission models.Transmission
}

type Entry struct {
	Id           uint
	Championship uint
	RaceNumber   uint
	Team         string
	Drivers      []Driver `gorm:"many2many:driver_entries;joinForeignKey:entry"`
}

func (Entry) TableName() string {
	return "entries"
}

type Race struct {
	Id           uint
	Name         string
	Championship string
	Layout       string
	Track        string
	Datetime     time.Time
}

type Championship struct {
	Id      uint
	Name    string
	Year    uint
	Entries []Entry    `gorm:"foreignKey:championship"`
	Races   []Race     `gorm:"foreignKey:championship"`
	Classes []CarClass `gorm:"many2many:championship_classes;joinForeignKey:championship"`
}
