package sqlite

import (
	"database/sql"

	"github.com/yeboahnanaosei/ghdata"
)

// DistrictService handles districts
type DistrictService struct {
	DB *sql.DB
}

// GetAllDistricts returns all districts
func (d *DistrictService) GetAllDistricts() ([]ghdata.District, error) {
	var districts []ghdata.District

	rows, err := d.DB.Query("SELECT name, capital, level, region FROM districts ORDER BY region")
	if err != nil {
		return districts, err
	}

	for rows.Next() {
		d := ghdata.District{}
		err := rows.Scan(&d.Name, &d.Capital, &d.Level, &d.RegionCode)
		if err != nil {
			return districts, err
		}
		districts = append(districts, d)
	}
	defer rows.Close()
	return districts, nil
}

// GetDistrictsByRegion returns all districts whose region = regionCode
func (d *DistrictService) GetDistrictsByRegion(regionCode string) ([]*ghdata.District, error) {
	var districts []*ghdata.District
	rows, err := d.DB.Query("SELECT name, capital, level FROM districts WHERE region = ? ORDER BY name", regionCode)
	if err != nil {
		return districts, err
	}

	for rows.Next() {
		d := ghdata.District{}
		err := rows.Scan(&d.Name, &d.Capital, &d.Level)
		if err != nil {
			return districts, err
		}
		districts = append(districts, &d)
	}
	defer rows.Close()
	return districts, nil
}
