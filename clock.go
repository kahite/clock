package clock

import (
	"math"
	"time"
)

const (
	secondLength       = 90
	minuteLength       = 80
	hourLength         = 50
	clockCentreX       = 150
	clockCentreY       = 150
	secondsInHalfClock = 30
	secondsInClock     = 2 * secondsInHalfClock
	minutesInHalfClock = 30
	minutesInClock     = 2 * minutesInHalfClock
	hoursInHalfClock   = 6
	hoursInClock       = 2 * hoursInHalfClock
)

// Point is a point
type Point struct {
	X float64
	Y float64
}

// Seconds get the Point for the seconds
func Seconds(t time.Time) Point {
	return adjustPointToClock(secondHandPoint(t))
}

// Minutes get the Point for the minutes
func Minutes(t time.Time) Point {
	return adjustPointToClock(minuteHandPoint(t))
}

// Hours get the Point for the hours
func Hours(t time.Time) Point {
	return adjustPointToClock(hourHandPoint(t))
}

func adjustPointToClock(p Point) Point {
	p = Point{p.X * secondLength, p.Y * secondLength}
	p = Point{p.X, -p.Y}
	p = Point{p.X + clockCentreX, p.Y + clockCentreY}
	return p
}

func secondsInRadians(t time.Time) float64 {
	return math.Pi / (secondsInHalfClock / float64(t.Second()))
}

func minutesInRadians(t time.Time) float64 {
	return math.Pi / (minutesInHalfClock / float64(t.Minute()))
}

func hoursInRadians(t time.Time) float64 {
	return math.Pi / (hoursInHalfClock / float64(t.Hour()))
}

func secondHandPoint(t time.Time) Point {
	return angleToPoint(secondsInRadians(t))
}

func minuteHandPoint(t time.Time) Point {
	return angleToPoint(minutesInRadians(t))
}

func hourHandPoint(t time.Time) Point {
	return angleToPoint(hoursInRadians(t))
}

func angleToPoint(angle float64) Point {
	return Point{math.Sin(angle), math.Cos(angle)}
}

func simpleTime(hours, minutes, seconds int) time.Time {
	return time.Date(1337, time.December, 12, hours, minutes, seconds, 0, time.UTC)
}

func roughlyEqualFloat64(a, b float64) bool {
	const thresold = 1e-7
	return math.Abs(a-b) < thresold
}

func roughlyEqualPoint(a, b Point) bool {
	return roughlyEqualFloat64(a.X, b.X) && roughlyEqualFloat64(a.Y, b.Y)
}
