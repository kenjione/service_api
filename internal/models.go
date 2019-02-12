package internal

type Location struct {
	ID           int64  `json:"id"`
	IPaddress    string `json:"ip_address"`
	CountryCode  string `json:"country_code"`
	Country      string `json:"country"`
	City         string `json:"city"`
	Latitude     string `json:"latitude"`
	Longitude    string `json:"longitude"`
	MysteryValue string `json:"mystery_value"`
}
