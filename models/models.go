package models

import "time"

type Transmission string
type Drivetrain string
type LayoutType string
type Sex string

const (
	Male   Sex = "M"
	Female Sex = "F"
)

const (
	AWD Drivetrain = "AWD"
	RWD Drivetrain = "RWD"
	FWD Drivetrain = "FWD"
)

const (
	Sequential Transmission = "Sequential"
	Manual     Transmission = "Manual"
)

const (
	Oval       LayoutType = "Oval"
	RoadCourse LayoutType = "RoadCourse"
)

type Brand struct {
	Name   string
	Nation string
}

type Car struct {
	Model        string
	Year         uint
	Brand        Brand
	Class        string
	Drivetrain   Drivetrain
	Transmission Transmission
}

type Track struct {
	Name     string
	Nation   string
	Location string
	Layouts  []Layout
}

type Layout struct {
	Name      string
	Length    uint
	TrackType LayoutType
}

type Team struct {
	Name   string
	Nation string
}

type Driver struct {
	Name      string
	Surname   string
	CF        string
	Sex       Sex
	Birthdate time.Time
}
type Race struct {
	Name             string
	Date             time.Time
	Track            Track
	LayoutName       string
	ChampionshipName string
}

type Entry struct {
	Car        Car
	RaceNumber uint
	Drivers    []Driver
	Team       Team
}

type Championship struct {
	Name      string
	Year      uint
	EntryList []Entry
	Races     []Race
	Classes   []CarClass
}

type CarUsage struct {
	Model        string
	Year         uint
	Brand        string
	Class        string
	Drivetrain   Drivetrain
	Transmission Transmission
	Usage        uint
}

type LayoutUsage struct {
	Track  string
	Name   string
	Length uint
	Type   LayoutType
	Usage  uint
}

type CarClass struct {
	Name string
}
