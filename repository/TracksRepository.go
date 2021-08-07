package repository

import (
	"ProgettoDB/models"
	"gorm.io/gorm"
)

type TracksRepository struct {
	Db *gorm.DB
}

func (t TracksRepository) GetAllTracks() ([]models.Track, error) {
	var dbTracks []Track

	if err := t.Db.Debug().Preload("Layouts").Find(&dbTracks).Error; err != nil {
		return nil, err
	}

	var tracks []models.Track
	for _, track := range dbTracks {
		var layouts []models.Layout

		for _, layout := range track.Layouts {
			layouts = append(layouts, models.Layout{
				Name:      layout.Name,
				Length:    layout.Length,
				TrackType: layout.Type,
			})
		}

		tracks = append(tracks, models.Track{
			Name:     track.Name,
			Nation:   track.Nation,
			Location: track.Location,
			Layouts:  layouts,
		})
	}

	return tracks, nil
}
