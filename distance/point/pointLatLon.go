package point

type PoinLatLon struct {
	lat float64
	lon float64
}

func (p PoinLatLon) Lat() float64 {
	return p.lat
}

func (p PoinLatLon) Lon() float64 {
	return p.lon
}

func NewPoinLatLon(lat float64, lon float64) *PoinLatLon {
	return &PoinLatLon{lat: lat, lon: lon}
}
