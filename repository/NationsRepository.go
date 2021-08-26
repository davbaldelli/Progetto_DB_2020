package repository

import "gorm.io/gorm"

type NationsRepository struct {
	Db *gorm.DB
}

func (n NationsRepository) GetAllNations() ([]string, error) {
	var dbNations []Nation
	if err := n.Db.Find(&dbNations).Error; err != nil {
		return nil, err
	}

	var nations []string

	for _, nation := range dbNations {
		nations = append(nations, nation.Name)
	}

	return nations, nil
}

