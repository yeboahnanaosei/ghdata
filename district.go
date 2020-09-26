package ghdata

import "database/sql"

// District represents one district
type District struct {
	ID         string  `json:"id,omitempty"`
	Name       string  `json:"name"`
	Capital    string  `json:"capital"`
	Level      string  `json:"level"`
	RegionCode string  `json:"regionCode,omitempty"`
	Region     *Region `json:"region,omitempty"`
}

// DistrictService exposes an interface to handle and manipulate districts
type DistrictService interface {
	// District returns one district
	GetOneDistrict(ID string) District

	// Returns a slice of all districts
	GetAllDistricts() []District

	// Returns a slice of all districts whose region code matches code
	GetDistrictsByRegion(regionCode string) ([]*District, error)

	// Get districts by level
	GetDistrictsByLevels(levels []string) ([]District, error)

	// SearchDistrict returns districts whose name matches keyword
	SearchDistrict(keyword string) ([]Region, error)

	// CustomQuery allows user land to pass their custom query
	CustomQuery(query string) (*sql.Rows, error)
}
