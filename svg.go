package clock

import (
	"encoding/xml"
	"fmt"
	"io"
	"time"
)

const svgStart = `<?xml version="1.0" encoding="UTF-8" standalone="no"?>
<!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1//EN" "http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd">
<svg xmlns="http://www.w3.org/2000/svg"
     width="100%"
     height="100%"
     viewBox="0 0 300 300"
     version="2.0">`
const bezel = `<circle cx="150" cy="150" r="100" style="fill:#fff;stroke:#000;stroke-width:5px;"/>`
const svgEnd = `</svg>`
const hand = `<line x1="150" y1="150" x2="%.3f" y2="%.3f" style="fill:none;stroke:%s;stroke-width:3px;"/>`

const redColor = "#f00"
const blackColor = "#000"

// SVG is SVG
type SVG struct {
	XMLName xml.Name `xml:"svg"`
	Xmlns   string   `xml:"xmlns,attr"`
	Width   string   `xml:"width,attr"`
	Height  string   `xml:"height,attr"`
	ViewBox string   `xml:"viewBox,attr"`
	Version string   `xml:"version,attr"`
	Circle  Circle   `xml:"circle"`
	Line    []Line   `xml:"line"`
}

// Circle is Circle
type Circle struct {
	Cx string `xml:"cx,attr"`
	Cy string `xml:"cy,attr"`
	R  string `xml:"r,attr"`
}

// Line is Line
type Line struct {
	X1 string `xml:"x1,attr"`
	Y1 string `xml:"y1,attr"`
	X2 string `xml:"x2,attr"`
	Y2 string `xml:"y2,attr"`
}

// SVGWriter creates the SVG for the clock
func SVGWriter(w io.Writer, t time.Time) {
	io.WriteString(w, svgStart)
	io.WriteString(w, bezel)
	secondsInSVG(w, t)
	minutesInSVG(w, t)
	hoursInSVG(w, t)
	io.WriteString(w, svgEnd)
}

func secondsInSVG(w io.Writer, t time.Time) {
	drawHand(w, secondHandPoint(t), secondLength, redColor)
}

func minutesInSVG(w io.Writer, t time.Time) {
	drawHand(w, minuteHandPoint(t), minuteLength, blackColor)
}

func hoursInSVG(w io.Writer, t time.Time) {
	drawHand(w, hourHandPoint(t), hourLength, blackColor)
}

func drawHand(w io.Writer, p Point, length float64, color string) {
	p = Point{p.X * length, p.Y * length}
	p = Point{p.X, -p.Y}
	p = Point{p.X + clockCentreX, p.Y + clockCentreY}
	fmt.Fprintf(w, hand, p.X, p.Y, color)
}

func containsLine(a Line, b []Line) bool {
	for _, l := range b {
		if l == a {
			return true
		}
	}
	return false
}
