package database

import (
	"fmt"
	"time"
)

type Stuff struct {
	Id      int64
	Title   string
	Zip     string
	Lat     float64
	Lon     float64
	Contact string
	Date    time.Time
}

func (s *Stuff) String() string {
	return fmt.Sprintf("Stuff %i: %s\nAt %f %f",
		s.Id, s.Title, s.Lat, s.Lon)
}
