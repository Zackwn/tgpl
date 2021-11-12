package main

import (
	"fmt"
	"log"
	"math"
	"os"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

type CornerType int

const (
	middle CornerType = iota
	peak
	valley
)

func e3() {
	file, err := os.Create("surface.svg")
	if err != nil {
		log.Fatal(err)
	}
	file.Write([]byte(fmt.Sprintf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)))
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, ct := e3corner(i+1, j)
			bx, by, ct1 := e3corner(i, j)
			cx, cy, ct3 := e3corner(i, j+1)
			dx, dy, ct2 := e3corner(i+1, j+1)

			// Skip invalid points.
			if math.IsInf(ax, 0) ||
				math.IsInf(ay, 0) ||
				math.IsInf(bx, 0) ||
				math.IsInf(by, 0) ||
				math.IsInf(cx, 0) ||
				math.IsInf(cy, 0) ||
				math.IsInf(dx, 0) ||
				math.IsInf(dy, 0) {
				fmt.Println("Skip here")
				continue
			}

			var color string
			if ct == peak || ct1 == peak || ct2 == peak || ct3 == peak {
				color = "#f00"
			} else if ct == valley || ct1 == valley || ct2 == valley || ct3 == valley {
				color = "#00f"
			} else {
				// same as default
				color = "grey"
			}

			file.Write([]byte(fmt.Sprintf("<polygon points='%g,%g %g,%g %g,%g %g,%g' style='stroke: %s'/>\n", ax, ay, bx, by, cx, cy, dx, dy, color)))
		}
	}
	file.Write([]byte("</svg>"))
}

func e3corner(i, j int) (float64, float64, CornerType) {
	// Find point (x, y) at corner of cell (i, j).
	x := xyrange * (float64(i)/cells - .5)
	y := xyrange * (float64(j)/cells - .5)

	// Compute surface height z.
	z, ct := e3f(x, y)

	// Project (x, y, z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, ct
}

func e3f(x, y float64) (float64, CornerType) {
	r := math.Hypot(x, y) // distance from (0,0)
	ct := middle

	// f(x) = sin(x)/x, f'(x) = (x*cos(x)-sin(x))/x^2
	// f'(x) = 0 ==> x = tan(x), peak or vally
	// if f''(x) > 0, vally
	// if f''(x) < 0, peak
	// f''(x) = {2(sin(x)-x*cos(x)) - x*x*sin(x)}/x*x*x
	if math.Abs(r-math.Tan(r)) < 3 {
		ct = peak
		if 2*(math.Sin(r)-r*math.Cos(r))-r*r*math.Sin(r) > 0 {
			ct = valley
		}
	}
	return math.Sin(r) / r, ct
}
