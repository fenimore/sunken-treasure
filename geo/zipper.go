package geo

import "github.com/codingsince1985/geo-golang/openstreetmap"

// Resolve returns the float coordinates of an address
// or zip code passed to it. It does not, however, resolve
// Canadian zip codes?
func Resolve(zip string) ([2]float64, error) {
	var coord [2]float64
	geocoder := openstreetmap.Geocoder()
	loc, err := geocoder.Geocode(zip)
	if err != nil {
		return coord, err
	}
	coord[0], coord[1] = loc.Lat, loc.Lng
	return coord, err
}
