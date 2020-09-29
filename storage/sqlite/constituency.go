package sqlite

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/yeboahnanaosei/ghdata"
)

// ConstituencyService handles constituencies
type ConstituencyService struct {
	DB *sql.DB
}

// GetAllConstituencies returns all constituencies in the system
func (s *ConstituencyService) GetAllConstituencies() ([]ghdata.Constituency, error) {
	constituencies := []ghdata.Constituency{}

	rows, err := s.DB.Query("SELECT * FROM constituencies ORDER BY region")
	if err == sql.ErrNoRows {
		return constituencies, nil
	}
	defer rows.Close()

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		c := ghdata.Constituency{}
		err := rows.Scan(&c.ID, &c.Name, &c.RegionCode)
		if err != nil {
			return constituencies, err
		}

		constituencies = append(constituencies, c)
	}

	return constituencies, nil
}

// GetConstituenciesByRegion returns all constituencies in a particular region
func (s *ConstituencyService) GetConstituenciesByRegion(regionCode string) ([]ghdata.Constituency, error) {
	constituencies := []ghdata.Constituency{}

	rows, err := s.DB.Query("SELECT * FROM constituencies WHERE region = ?", regionCode)
	if err == sql.ErrNoRows {
		return constituencies, nil
	}
	defer rows.Close()

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		c := ghdata.Constituency{}
		err := rows.Scan(&c.ID, &c.Name, &c.RegionCode)
		if err != nil {
			return nil, err
		}
		constituencies = append(constituencies, c)
	}

	return constituencies, nil
}

// CustomQuery allows a user to pass their own custom query
func (s *ConstituencyService) CustomQuery(query string) (*sql.Rows, error) {
	return s.DB.Query(query)
}

// GetConstituenciesByRegions takes a slice of region codes and returns constituencies
// who belong to the regions in the regionCodes slice
func (s *ConstituencyService) GetConstituenciesByRegions(regionCodes []string) ([]ghdata.Constituency, error) {
	// Prepare the list of regions for query
	if len(regionCodes) < 1 || regionCodes == nil {
		return nil, errors.New("empty slice supplied")
	}

	var queryString string
	for _, regionCode := range regionCodes {
		queryString += fmt.Sprintf(`'%s',`, regionCode)
	}
	queryString = strings.TrimSuffix(queryString, ",")
	finalQuery := fmt.Sprintf("SELECT * FROM constituencies WHERE region IN (%s) ORDER BY region", queryString)

	rows, err := s.DB.Query(finalQuery)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	constituencies := []ghdata.Constituency{}
	for rows.Next() {
		c := ghdata.Constituency{}
		err := rows.Scan(&c.ID, &c.Name, &c.RegionCode)
		if err != nil {
			return nil, err
		}
		constituencies = append(constituencies, c)
	}

	return constituencies, nil
}

// SearchConstituencies searches for a constituency whose name matches the supplied
// keyword
func (s *ConstituencyService) SearchConstituencies(keyword string) ([]ghdata.Constituency, error) {
	if len(keyword) == 0 {
		return nil, errors.New("empty keyword supplied")
	}

	rows, err := s.DB.Query("SELECT * FROM constituencies WHERE name LIKE ?", "%"+keyword+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	constituencies := []ghdata.Constituency{}
	for rows.Next() {
		c := ghdata.Constituency{}
		err := rows.Scan(&c.ID, &c.Name, &c.RegionCode)
		if err != nil {
			return nil, err
		}
		constituencies = append(constituencies, c)
	}

	return constituencies, nil
}
