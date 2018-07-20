package models

import "math"

// Sun contains the coordinates of that star
type Sun struct {
	//Coordinates
	X int
	Y int
}

// Vulcan contains the properties of that planet
type Vulcan struct {
	Raduis int
	Speed  int
	//Coordinates Coordinates
	X int
	Y int
}

// Betazoid contains the properties of that planet
type Betazoid struct {
	Raduis int
	Speed  int
	// Coordinates Coordinates
	X int
	Y int
}

// Ferengi contains the properties of that planet
type Ferengi struct {
	Raduis int
	Speed  int
	// Coordinates Coordinates
	X int
	Y int
}

// SolarSystems is a slice to SolarSystem
type SolarSystems []SolarSystem

// SolarSystem is the structure that contains the planets with their coordinates
type SolarSystem struct {
	Sun      Sun
	Vulcan   Vulcan
	Betazoid Betazoid
	Ferengi  Ferengi
}

// Coordinates of each element of the system
type Coordinates struct {
	X int
	Y int
}

var theta float64

const (
	// PI is a contant with number of PI
	PI = 3.14159265358979323846264338327950288419716939937510582097494459

	// Tolerance to diff
	Tolerance = 0.00001

	// Decimals to round
	Decimals = 1
)

// Perimiter calculate perimeter to triangle
func (ss *SolarSystem) Perimiter() float64 {
	a := float64(ss.Ferengi.X - ss.Betazoid.X)
	b := float64(ss.Betazoid.Y - ss.Ferengi.Y)
	c := float64(ss.Betazoid.X - ss.Vulcan.X)
	d := float64(ss.Vulcan.Y - ss.Betazoid.Y)
	e := float64(ss.Vulcan.X - ss.Ferengi.X)
	f := float64(ss.Ferengi.Y - ss.Vulcan.Y)
	return math.Sqrt(math.Pow(a, 2)+math.Pow(b, 2)) + math.Sqrt(math.Pow(c, 2)+math.Pow(d, 2)) + math.Sqrt(math.Pow(e, 2)+math.Pow(f, 2))
}

// TriangleContainSun check if the triangle contains the sun
func (ss *SolarSystem) TriangleContainSun() bool {
	var a = Sign(ss.Sun.X, ss.Sun.Y, ss.Ferengi.X, ss.Ferengi.Y, ss.Betazoid.X, ss.Betazoid.Y) < 0
	var b = Sign(ss.Sun.X, ss.Sun.Y, ss.Betazoid.X, ss.Betazoid.Y, ss.Vulcan.X, ss.Vulcan.Y) < 0
	var c = Sign(ss.Sun.X, ss.Sun.Y, ss.Vulcan.X, ss.Vulcan.Y, ss.Ferengi.X, ss.Ferengi.Y) < 0
	return a == b && b == c
}

// Sign calculate them
func Sign(x1, y1, x2, y2, x3, y3 int) int {
	return (x1-x3)*(y2-y3) - (x2-x3)*(y1-y3)
}

// TheyAreOnAxes check if they are on axes
func (ss *SolarSystem) TheyAreOnAxes() bool {
	if math.Abs(float64(ss.Ferengi.X)) < Tolerance && math.Abs(float64(ss.Betazoid.X)) < Tolerance && math.Abs(float64(ss.Vulcan.X)) < Tolerance {
		return true
	}
	return math.Abs(float64(ss.Ferengi.Y)) < Tolerance && math.Abs(float64(ss.Betazoid.Y)) < Tolerance && math.Abs(float64(ss.Vulcan.Y)) < Tolerance
}

// TheyAreParallels check if they are parallels
func (ss *SolarSystem) TheyAreParallels() bool {
	a := float64(ss.Betazoid.X - ss.Ferengi.X)
	b := float64(ss.Vulcan.X - ss.Betazoid.X)
	c := float64(ss.Betazoid.Y - ss.Ferengi.Y)
	d := float64(ss.Vulcan.Y - ss.Betazoid.Y)
	return math.Abs(math.Round((a)/(b))-math.Round((c)/(d))) < Tolerance
}

// TheyPassThroughTheSun check if they pass through
func (ss *SolarSystem) TheyPassThroughTheSun() bool {
	a := float64(ss.Sun.X - ss.Ferengi.X)
	b := float64(ss.Betazoid.X - ss.Ferengi.X)
	c := float64(ss.Sun.Y - ss.Ferengi.Y)
	d := float64(ss.Betazoid.Y - ss.Ferengi.Y)
	return math.Abs(math.Round((a)/(b))-math.Round((c)/(d))) < Tolerance
}

// Movement return coodinates of Vulcan
func (v *Vulcan) Movement(interval int) error {

	theta = float64(v.Speed) * PI / 180
	ti := theta * float64(interval)
	point, err := AntiClockwisePoint(v.Raduis, ti)
	if err != nil {
		return err
	}
	v.X = point.X
	v.Y = point.Y
	return nil
}

// Movement return coodinates of Betazoid
func (b *Betazoid) Movement(interval int) error {

	theta = float64(b.Speed) * PI / 180
	ti := theta * float64(interval)
	point, err := ClockwisePoint(b.Raduis, ti)
	if err != nil {
		return err
	}
	b.X = point.X
	b.Y = point.Y
	return nil
}

// Movement return coodinates of Ferengi
func (f *Ferengi) Movement(interval int) error {

	theta = float64(f.Speed) * PI / 180
	ti := theta * float64(interval)
	point, err := ClockwisePoint(f.Raduis, ti)
	if err != nil {
		return err
	}
	f.X = point.X
	f.Y = point.Y
	return nil

}

// ClockwisePoint Calculate coordinates in clock wise point
func ClockwisePoint(radius int, theta float64) (Coordinates, error) {
	X := math.Round(float64(radius) * math.Cos(theta))
	Y := math.Round(float64(radius) * math.Sin(theta))
	c := Coordinates{X: int(X), Y: int(Y)}
	return c, nil
}

// AntiClockwisePoint Calculate coordinates in anticlock wise point
func AntiClockwisePoint(radius int, theta float64) (Coordinates, error) {
	t := theta + PI/2
	X := math.Round(float64(radius) * math.Cos(t))
	Y := math.Round(float64(radius) * math.Sin(t))
	c := Coordinates{X: int(X), Y: int(Y)}
	return c, nil
}
