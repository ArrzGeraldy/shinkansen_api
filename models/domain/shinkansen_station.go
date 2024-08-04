package domain

type Station struct {
	Id                int     `json:"id"`
	StationName       string  `json:"station_name"`
	ShinkansenLine    string  `json:"shinkansen_line"`
	Year              int     `json:"year"`
	Prefecture        string  `json:"prefecture"`
	DistanceFromTokyo float64 `json:"distance_from_tokyo"`
	Company           string  `json:"company"`
}