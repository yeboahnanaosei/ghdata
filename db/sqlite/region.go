package sqlite

import (
	"database/sql"

	"github.com/yeboahnanaosei/ghdata"
)

// RegionService implements ghdata.RegionService
type RegionService struct {
	DB *sql.DB
}

// GetOneRegion fetches one region from the storage
func (r *RegionService) GetOneRegion(code string) (ghdata.Region, error) {
	var region ghdata.Region
	row := r.DB.QueryRow("SELECT * FROM regions WHERE code = ?", code)
	err := row.Scan(&region.Code, &region.Name, &region.Capital)
	if err != nil {
		return region, err
	}
	return region, nil
}

// GetAllRegions returns all regions in the system
func (r *RegionService) GetAllRegions() ([]*ghdata.Region, error) {
	var regions []*ghdata.Region
	rows, err := r.DB.Query("SELECT * FROM regions ORDER BY code")
	if err != nil {
		return regions, err
	}
	defer rows.Close()

	for rows.Next() {
		region := ghdata.Region{}
		err := rows.Scan(&region.Code, &region.Name, &region.Capital)
		if err != nil {
			return regions, err
		}
		regions = append(regions, &region)
	}
	defer rows.Close()

	return regions, nil
}

// Search returns ghdata.Regions whose name or code matches keyword
func (r *RegionService) Search(keyword string) ([]ghdata.Region, error) {
	var regions []ghdata.Region

	rows, err := r.DB.Query(
		"SELECT * FROM regions WHERE name LIKE ? OR code LIKE ? ORDER BY name",
		"%"+keyword+"%",
		"%"+keyword+"%",
	)

	if err != nil {
		return regions, err
	}
	defer rows.Close()

	for rows.Next() {
		r := ghdata.Region{}
		err := rows.Scan(&r.Code, &r.Name, &r.Capital)
		if err != nil {
			return regions, err
		}
		regions = append(regions, r)
	}

	return regions, nil
}
