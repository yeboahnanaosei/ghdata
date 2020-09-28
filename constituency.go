package ghdata

import "database/sql"

// A Constituency models one constituency
type Constituency struct {
	ID         int    `json:"-"`
	Name       string `json:"name"`
	RegionCode string `json:"regionCode"`
}

// ConstituencyService exposes methods to handle constituencies
type ConstituencyService interface {
	GetAllConstituencies() ([]Constituency, error)
	GetConstituenciesByRegion(regionCode string) ([]Constituency, error)
	GetConstituenciesByRegions(regionCodes []string) ([]Constituency, error)
	CustomQuery(query string) (*sql.Rows, error)
	SearchConstituency(keyword string) ([]Constituency, error)
}
