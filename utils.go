package main

import (
	"image/color"
	"math"
)

const (
	canvasWidth  = 640
	canvasHeight = 480

	vpWidth  float64 = 1
	vpHeight float64 = 1
	vpDist   float64 = 1
)

// CanvasToViewport takes canvas coordinates (pixel coordinates) and
//	returns their position in viweport coordinates.
func CanvasToViewport(cx, cy int) Vec {

	cxf := float64(cx)
	cyf := float64(cy)

	vx := cxf * vpWidth / canvasWidth
	vy := cyf * vpHeight / canvasHeight
	vz := vpDist

	return *NewVector(vx, vy, vz)
}

func TraceRay(O, D Vec, t_min, t_max float64, scene []Sphere) color.RGBA {
	closest_t := math.MaxFloat64
	var closestSphere *Sphere = nil

	for _, sphere := range scene {
		t1, t2 := sphere.IntersectSphere(O, D)

		if (t1 >= t_min && t1 < t_max) && (t1 < closest_t) {
			closest_t = t1
			closestSphere = &sphere
		}

		if (t2 >= t_min && t2 < t_max) && (t2 < closest_t) {
			closest_t = t2
			closestSphere = &sphere
		}

		if closestSphere != nil {
			return closestSphere.GetColor()
		}

		return color.RGBA{0, 0, 0, 0xff}
	}

	return color.RGBA{0, 0, 0, 0xff}

}
