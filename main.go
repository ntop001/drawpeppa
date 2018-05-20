package main

import (
	"math"
	"fmt"
	"flag"
	//"strconv"
	"strconv"
)

func f(x, y float64) float64 {
	return math.Sqrt(x*x + y*y) - 0.8
}

// draw circle
func outline(x, y float64) rune {
	delta := 0.001
	switch {
	case math.Abs(f(x, y)) < .05:
		dx := f(x+delta, y) - f(x-delta, y)
		dy := f(x, y+delta) - f(x, y-delta)
		ii := (math.Atan2(dy, dx)/6.2831853072 + 0.5) * 8 + 0.5
		ru := []rune("|/=\\|/=\\|")
		return ru[int(ii)]
	case f(x, y) < 0:
		return '.'
	default:
		return ' '
	}
}

// draw peppa
func c(x, y, r float64) float64 { return math.Sqrt(x*x + y*y) -r }
func u(x, y, t float64) float64 { return x*math.Cos(t) + y*math.Sin(t)}
func v(x, y, t float64) float64 { return y*math.Cos(t)-x*math.Sin(t) }

func fa(x, y float64) float64 { return math.Min(c(x,y,0.5),c(x*0.47+0.15,y+0.25,0.3))}
func no(x, y float64) float64 { return c(x*1.2+0.97,y+0.25,0.2)}
func nh(x, y float64) float64 { return math.Min(c(x+0.9,y+0.25,0.03),c(x+0.75,y+0.25,0.03))}
func ea(x, y float64) float64 { return math.Min(c(x*1.7+0.3,y+0.7,0.15),c(u(x,y,0.25)*1.7,v(x,y,0.25)+0.65,0.15))}
func ey(x, y float64) float64 { return math.Min(c(x+0.4,y+0.35,0.1),c(x+0.15,y+0.35,0.1))}
func pu(x, y float64) float64 { return math.Min(c(x+0.38,y+0.33,0.03),c(x+0.13,y+0.33,0.03))}
func fr(x, y float64) float64 { return c(x*1.1-0.3,y+0.1,0.15)}
func mo(x, y float64) float64 { return math.Max(c(x+0.15,y-0.05,0.2),-c(x+0.15,y,0.25))}

func o(x, y float64,f func(f1, f2 float64) float64, i float64) float64 {
	r := f(x, y)
	switch {
	case math.Abs(r) < 0.02:
		return (math.Atan2(f(x,y+1e-3)-r,f(x+1e-3,y)-r)+0.3)*1.273+6.5
	case r < 0:
		return i
	default:
		return 0
	}
}

func s(x, y float64, f func(f1, f2 float64) float64, i float64) float64 {
	if r := f(x, y); r < 0 {
		return i
	} else {
		return 0
	}
}

func f2(x, y float64) float64 {
	if o(x, y, no, 1) != 0 {
		return math.Max(o(x,y,no,1),s(x,y,nh,12))
	} else {
		return math.Max(o(x,y,fa,1),math.Max(o(x,y,ey,11),math.Max(o(x,y,ea,1),math.Max(o(x,y,mo,1),math.Max(s(x,y,fr,13),s(x,y,pu,12))))))
	}
}

//// This code is used to draw circle.
//func main() {
//	for y := float64(-1); y < 1; y += 0.05 {
//		for x := float64(-1); x < 1; x+= 0.025 {
//			fmt.Print(string(outline(x, y)))
//		}
//		fmt.Print("\n")
//	}
//}
//
func main() {
	var sk float64 = 1
	if f64, err := strconv.ParseFloat(flag.Arg(0), 10); err == nil {
		sk = f64
	}
	for y := float64(-1); y < .6; y += 0.05/sk {
		for x := float64(-1); x < .6; x+= 0.025/sk {
			ii := int(f2(u(x,y,0.3),v(x,y,0.3)))
			fmt.Print(string([]rune(" .|/=\\|/=\\| @!")[ii]))
		}
		fmt.Print("\n")
	}
}

func init() {
	flag.Parse()
}