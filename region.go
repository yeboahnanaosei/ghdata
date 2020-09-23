package ghdata

// Region models one region
type Region struct {
	Code      string      `json:"code"`
	Name      string      `json:"name"`
	Capital   string      `json:"capital"`
	Districts []*District `json:"districts,omitempty"`
}

// RegionService exposes an interface to handle and manipulate regions
type RegionService interface {
	GetOneRegion(code string) (Region, error) // Region returns one region whose code matches code
	GetAllRegions() ([]Region, error)   // GetAllRegions returns a slice of all regions
}
