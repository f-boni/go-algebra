/*
	/\\\\\\\\\\\\\\\                  /\\\\\\\\\\\\\         /\\\\\       /\\\\\     /\\\  /\\\\\\\\\\\
	\/\\\///////////                  \/\\\/////////\\\     /\\\///\\\    \/\\\\\\   \/\\\ \/////\\\///
	 \/\\\                             \/\\\       \/\\\   /\\\/  \///\\\  \/\\\/\\\  \/\\\     \/\\\
	  \/\\\\\\\\\\\                     \/\\\\\\\\\\\\\\   /\\\      \//\\\ \/\\\//\\\ \/\\\     \/\\\
	   \/\\\///////                      \/\\\/////////\\\ \/\\\       \/\\\ \/\\\\//\\\\/\\\     \/\\\
	    \/\\\                             \/\\\       \/\\\ \//\\\      /\\\  \/\\\ \//\\\/\\\     \/\\\
	     \/\\\                             \/\\\       \/\\\  \///\\\  /\\\    \/\\\  \//\\\\\\     \/\\\
	      \/\\\              /\\\           \/\\\\\\\\\\\\\/     \///\\\\\/     \/\\\   \//\\\\\  /\\\\\\\\\\\
	       \///              \///            \/////////////         \/////       \///     \/////  \///////////

	Created:    12 mar 2026
	Author:     F. Boni    Email:      fabioboni96@hotmail.com
	Repository: github.com/f-boni/go-algebra

Copyright (c) 2026 Fabio Luis Boni - MIT License
*/
package algebra_frame

import (
	"fmt"
	"sort"
	"strings"
)

type SortingMethod string

const (
	X_HIGHER SortingMethod = "x_higher"
	Y_HIGHER SortingMethod = "y_higher"
	X_LOWER  SortingMethod = "x_lower"
	Y_LOWER  SortingMethod = "y_lower"
)

type Frame struct {
	Name  string
	Point []Point
}

type Point struct {
	X float64
	Y float64
}

func (frame *Frame) AddPoint(x float64, y float64) {
	frame.Point = append(
		frame.Point,
		Point{
			X: x,
			Y: y,
		},
	)
}

func (frame *Frame) Sort(method SortingMethod) {
	sort.Slice(
		frame.Point, func(i, j int) bool {
			switch method {
			case X_LOWER:
				return frame.Point[i].X < frame.Point[j].X

			case X_HIGHER:
				return frame.Point[i].X > frame.Point[j].X

			case Y_LOWER:
				return frame.Point[i].Y < frame.Point[j].Y

			case Y_HIGHER:
				return frame.Point[i].Y > frame.Point[j].Y

			default:
				return frame.Point[i].X < frame.Point[j].X
			}
		},
	)
}

func (frame Frame) String() string {
	var result strings.Builder
	if frame.Name != "" {
		fmt.Fprintf(&result, "%s:\n", frame.Name)
	}

	if len(frame.Point) == 0 {
		result.WriteString("[empty]")
		return result.String()
	}

	maxXWidth := 0
	maxYWidth := 0

	formatted := make([][2]string, len(frame.Point))

	for i, point := range frame.Point {
		xStr := fmt.Sprintf("%.2f", point.X)
		yStr := fmt.Sprintf("%.2f", point.Y)

		formatted[i] = [2]string{xStr, yStr}

		if len(xStr) > maxXWidth {
			maxXWidth = len(xStr)
		}
		if len(yStr) > maxYWidth {
			maxYWidth = len(yStr)
		}
	}

	for _, pair := range formatted {
		fmt.Fprintf(
			&result,
			"(%*s, %*s)\n",
			maxXWidth, pair[0],
			maxYWidth, pair[1],
		)
	}

	return result.String()
}

func (frame Frame) CSV() string {
	var result strings.Builder

	if frame.Name != "" {
		fmt.Fprintf(
			&result,
			"%s\n",
			frame.Name,
		)
	}
	result.WriteString("x;y\n")

	for _, point := range frame.Point {
		fmt.Fprintf(
			&result,
			"%g;%g\n",
			point.X,
			point.Y,
		)
	}

	return result.String()
}
