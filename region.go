package ghdata

// Region models one region
type Region struct {
	Code      string      `json:"code"`
	Name      string      `json:"name"`
	Capital   string      `json:"capital"`
	Districts []District `json:"districts,omitempty"`
}

// RegionService exposes an interface to handle and manipulate regions
type RegionService interface {
	// Region returns one region whose code matches code
	GetOneRegion(code string) (Region, error)

	// GetAllRegions returns a slice of all regions
	GetAllRegions() ([]Region, error)

	// SearchRegions returns regions whose name or code matches keyword
	SearchRegion(keyword string) ([]Region, error)
}
