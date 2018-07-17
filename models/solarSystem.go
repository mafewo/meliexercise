package models

// Sun contains the coordinates of that star
type Sun struct {
	Coordinates
}

// Vulcan contains the coordinates of that planet
type Vulcan struct {
	Coordinates
}

// Betazoid contains the coordinates of that planet
type Betazoid struct {
	Coordinates
}

// Ferengi contains the coordinates of that planet
type Ferengi struct {
	Coordinates
}

// SolarSystem is the structure that contains the planets with their coordinates
type SolarSystem struct {
	Sun
	Vulcan
	Betazoid
	Ferengi
}

// Coordinates of each element of the system
type Coordinates struct {
	x int
	y int
}

// TriangleContainSun check if the function contains the sun
func (ss *SolarSystem) TriangleContainSun() bool {
	var a = Sign(ss.Sun.Coordinates.x, ss.Sun.Coordinates.y, ss.Ferengi.Coordinates.x, ss.Ferengi.Coordinates.y, ss.Betazoid.Coordinates.x, ss.Betazoid.Coordinates.y) < 0
	var b = Sign(ss.Sun.Coordinates.x, ss.Sun.Coordinates.y, ss.Betazoid.Coordinates.x, ss.Betazoid.Coordinates.y, ss.Vulcan.Coordinates.x, ss.Vulcan.Coordinates.y) < 0
	var c = Sign(ss.Sun.Coordinates.x, ss.Sun.Coordinates.y, ss.Vulcan.Coordinates.x, ss.Vulcan.Coordinates.y, ss.Betazoid.Coordinates.x, ss.Betazoid.Coordinates.y) < 0
	return a == b && b == c
}

func Sign(x1, y1, x2, y2, x3, y3 int) int {
	return (x1-x3)*(y2-y3) - (x2-x3)*(y1-y3)
}
