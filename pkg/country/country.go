package country

// Country for fetching data interface
type Country interface {
	Fetch(id string) (*CountryType, error)
	FetchByName(name string) (*CountryType, error)
	FetchAll() ([]*CountryType, error)
}

// CountryType is type of user
type CountryType struct {
	ID   int    `json:"id" boil:"id"`
	Code string `json:"country_code" boil:"country_code"`
	Name string `json:"name" boil:"name"`
}
