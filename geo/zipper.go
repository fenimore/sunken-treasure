// geo will help encode user's zip codes into coordinates, and
// to take user location, from say a mobile device, and find
// the human readible zip code.
// TODO: Zips are only working in the states, wtf?
// TODO: reverse geocoder returns whole address, but that's silly.
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

// Reverse returns the reverse geocoded address from a given coordinate.
// TODO: Parse to return simply the zip
func Reverse(coord [2]float64) (string, error) {
	geocoder := openstreetmap.Geocoder()
	addr, err := geocoder.ReverseGeocode(coord[0], coord[1])
	return addr, err
}
