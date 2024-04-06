package main

// Variant-8 2x^5 + 3x^2 - 2x - 6 = 0

import (
	"fmt"
	"math"
)

func f(x float64) float64 {
	return 2*math.Pow(x, 5) + 3*math.Pow(x, 2) - 2*x - 6
}

func fd(x float64) float64 {
	return 10*math.Pow(x, 4) + 6*x - 2
}

func bisectionMethod(a float64, b float64, epsilon float64) float64 {
	var c float64
	for math.Abs(a-b) > epsilon {
		c = (a + b) / 2.0
		f_a := f(a)
		f_c := f(c)
		if f_a*f_c <= 0 {
			b = c
			continue
		}
		a = c
	}

	return c
}

func hordeMethod(a float64, b float64, epsilon float64) float64 {
	var c, c_prev float64
	for {
		c = (a*f(b) - b*f(a)) / (f(b) - f(a))
		if math.Abs(c-c_prev) < epsilon {
			break
		}
		f_a := f(a)
		f_c := f(c)
		if f_a*f_c <= 0 {
			b = c
		} else {
			a = c
		}
		c_prev = c
	}

	return c
}

func tangentMethod(x float64, epsilon float64) float64 {
	var x_next float64 = x
	for {
		x = x_next
		x_next = x - f(x)/fd(x)
		if math.Abs(x_next-x) < epsilon {
			break
		}
	}
	return x_next
}

func main() {
	fmt.Print("Bisection method: ")
	fmt.Println(bisectionMethod(2.0/3.0, 1.2, 0.00005))
	fmt.Print("hordeMethod: ")
	fmt.Println(hordeMethod(2.0/3.0, 1.2, 0.000000005))
	fmt.Print("tangentMethod: ")
	fmt.Println(tangentMethod(1.1645, 0.1))
}
