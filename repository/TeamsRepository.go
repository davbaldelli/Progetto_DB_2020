package repository

import (
	"ProgettoDB/models"
	"gorm.io/gorm"
)

type TeamsRepository struct {
	Db *gorm.DB
}

func (t TeamsRepository) GetAllTeams() ([]models.Team, error) {
	var dbTeams []Team
	if err := t.Db.Find(&dbTeams).Error; err != nil {
		return nil, err
	}

	var teams []models.Team
	for _, team := range dbTeams {
		teams = append(teams, models.Team{
			Name:   team.Name,
			Nation: team.Nation,
		})
	}
	return teams, nil
}

func (t TeamsRepository) GetTeamsWithoutParticipationByYear(year uint) ([]models.Team, error) {
	var dbTeams []Team
	if err := t.Db.Where("teams.name NOT IN (?)",
		t.Db.Table("teams").Select("teams.name").
			Joins("JOIN entries ON entries.team = teams.name").
			Joins("JOIN championships on championships.id = entries.championship").
			Group("teams.name").Where("championships.year = ?", year)).
		Find(&dbTeams).Error; err != nil {
		return nil, err
	}

	var teams []models.Team
	for _, team := range dbTeams {
		teams = append(teams, models.Team{
			Name:   team.Name,
			Nation: team.Nation,
		})
	}
	return teams, nil
}
