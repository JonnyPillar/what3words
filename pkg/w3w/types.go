package w3w

import "github.com/jonnypillar/what3words/internal/api"

// Words ...
type Words [3]string

// Coordinates ...
type Coordinates struct {
	Lat float64
	Lng float64
}

// Result defines a W3W result
type Result struct {
	Country      string `json:"country"`
	Square       Square `json:"square"`
	NearestPlace string `json:"nearestPlace"`
	Coordinates  Coords `json:"coordinates"`
	Words        string `json:"words"`
	Language     string `json:"language"`
	Map          string `json:"map"`
}

// Square ...
type Square struct {
	Southwest Southwest `json:"southwest"`
	Northeast Northeast `json:"northeast"`
}

// Southwest ...
type Southwest struct {
	Lng float64 `json:"lng"`
	Lat float64 `json:"lat"`
}

// Northeast ...
type Northeast struct {
	Lng float64 `json:"lng"`
	Lat float64 `json:"lat"`
}

// Coords ...
type Coords struct {
	Lng float64 `json:"lng"`
	Lat float64 `json:"lat"`
}

func newResponse(r *api.Response) *Result {
	return &Result{
		Country: r.Country,
		Square: Square{
			Southwest: r.Square.Southwest,
			Northeast: r.Square.Northeast,
		},
		NearestPlace: r.NearestPlace,
		Coordinates:  r.Coordinates,
		Words:        r.Words,
		Language:     r.Language,
		Map:          r.Map,
	}
}
