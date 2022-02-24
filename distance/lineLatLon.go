package distance

import (
	"Beautiful_Code/distance/point"
	"math"
)

type LineLatLon struct {
	point1 point.PoinLatLon
	point2 point.PoinLatLon
}

func (l LineLatLon) Point1() point.PoinLatLon {
	return l.point1
}

func (l LineLatLon) Point2() point.PoinLatLon {
	return l.point2
}

func NewLineLatLon(point1 point.PoinLatLon, point2 point.PoinLatLon) *LineLatLon {
	return &LineLatLon{point1: point1, point2: point2}
}

func (l LineLatLon) Distance() (distance float64, err error) {
	distance = 2 * 6471 * math.Asin(math.Sqrt(math.Pow(math.Sin((DegToRad(l.point2.Lat())-DegToRad(l.point1.Lat()))/2), 2)+math.Cos(DegToRad(l.point1.Lat()))*math.Cos(DegToRad(l.point2.Lat()))*math.Pow(math.Sin((DegToRad(l.point2.Lon())-DegToRad(l.point1.Lon()))/2), 2)))
	return
}
func DegToRad(x float64) float64 {
	return x * math.Pi / 180
}
