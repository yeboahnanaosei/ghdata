package sqlite

import (
	"database/sql"
	"fmt"
	"strings"

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

// GetDistrictsByLevels returns all districts whose level matches any one of the
// levels in the levels parameter
func (d *DistrictService) GetDistrictsByLevels(levels []string) ([]ghdata.District, error) {
	var districts []ghdata.District

	query := "SELECT name, capital, level FROM districts WHERE level IN ("
	for _, level := range levels {
		query += fmt.Sprintf(`'%s',`, level)
	}
	query = strings.TrimSuffix(query, ",")
	query += ")"

	rows, err := d.DB.Query(query)
	if err != nil {
		return districts, nil
	}
	defer rows.Close()

	for rows.Next() {
		d := ghdata.District{}
		rows.Scan(&d.Name, &d.Capital, &d.Level)
		districts = append(districts, d)
	}

	return districts, nil
}

// SearchDistrict returns ghdata.District whose name or code matches keyword
func (d *DistrictService) SearchDistrict(keyword string) ([]ghdata.District, error) {
	var districts []ghdata.District

	rows, err := d.DB.Query(
		"SELECT name, capital, level, region FROM districts WHERE name LIKE ? OR capital LIKE ? ORDER BY name",
		"%"+keyword+"%",
		"%"+keyword+"%",
	)

	if err != nil {
		return districts, err
	}
	defer rows.Close()

	for rows.Next() {
		d := ghdata.District{}
		err := rows.Scan(&d.Name, &d.Capital, &d.Level, &d.RegionCode)
		if err != nil {
			return districts, err
		}
		districts = append(districts, d)
	}

	return districts, nil
}
