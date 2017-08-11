package earth

import "math"

// Earth earth
type Earth struct{}

// NewEarth new earth
func NewEarth() *Earth {
	return &Earth{}
}

// 计算公式
// C = sin(LatA*Pi/180)*sin(LatB*Pi/180) + cos(LatA*Pi/180)*cos(LatB*Pi/180)*cos((MLonA-MLonB)*Pi/180)
// Distance = R*Arccos(C)*Pi/180

// Distance distance
func (e Earth) Distance(lat1, lng1, lat2, lng2 float64) float64 {
	radius := 6378.137
	rad := math.Pi / 180.0

	lat1 = lat1 * rad
	lng1 = lng1 * rad
	lat2 = lat2 * rad
	lng2 = lng2 * rad
	theta := lng2 - lng1

	dist := math.Acos(math.Sin(lat1)*math.Sin(lat2) + math.Cos(lat1)*math.Cos(lat2)*math.Cos(theta))
	return radius * dist
}
