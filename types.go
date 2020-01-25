package w3w

// Words ...
type Words [3]string

// Coordinates ...
type Coordinates struct {
	Lat float64
	Lng float64
}

// Options ...
type Options struct {
	APIURL string
	Format string
}

// Response defines the response body for Coordinates request
type Response struct {
	Country string `json:"country"`
	Square  struct {
		Southwest struct {
			Lng float64 `json:"lng"`
			Lat float64 `json:"lat"`
		} `json:"southwest"`
		Northeast struct {
			Lng float64 `json:"lng"`
			Lat float64 `json:"lat"`
		} `json:"northeast"`
	} `json:"square"`
	NearestPlace string `json:"nearestPlace"`
	Coordinates  struct {
		Lng float64 `json:"lng"`
		Lat float64 `json:"lat"`
	} `json:"coordinates"`
	Words    string `json:"words"`
	Language string `json:"language"`
	Map      string `json:"map"`
}
