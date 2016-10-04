package stuff

import (
	"fmt"
	"time"
)

type Stuff struct {
	Id      int64     `json:"id"`
	Title   string    `json:"title"`
	Zip     string    `json:"zip"`
	Lat     float64   `json:"latitude"`
	Lon     float64   `json:"longitude"`
	Contact string    `json:"contact"`
	Date    time.Time `json:"date"`
	Expired bool      `json:"expired"`
	// User should eventually beome it's own struct
	// User    string `json:"user"`
}

func (s *Stuff) String() string {
	return fmt.Sprintf("Stuff %i: %s\nAt %f %f",
		s.Id, s.Title, s.Lat, s.Lon)
}
