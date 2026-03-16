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
	Repository: github.com/FabioLuisBoni/go-algebra

Copyright (c) 2026 Fabio Luis Boni - MIT License
*/
package test

import (
	"fmt"
	"math"
	"testing"

	algebra_equation "github.com/FabioLuisBoni/go-algebra/equation"
	algebra_expression "github.com/FabioLuisBoni/go-algebra/equation/expression"
)

func TestEvaluateEquation(t *testing.T) {
	t.Logf("testing Evaluate for Equation\n\n")
	var x float64 = 5

	var expected float64 = (math.Pow(x, 2) / 2) + (math.Sin(x) * math.Cos(x))

	var equation *algebra_equation.Equation = algebra_equation.NewEquation("f(x)").SetExpression(
		algebra_expression.Sum(
			algebra_expression.Multiply(
				algebra_expression.Pow(
					algebra_expression.Int(2),
					algebra_expression.Int(-1),
				),
				algebra_expression.Pow(
					algebra_expression.Symbol("x"),
					algebra_expression.Int(2),
				),
			),
			algebra_expression.Multiply(
				algebra_expression.Cos(
					algebra_expression.Symbol("x"),
				),
				algebra_expression.Sin(
					algebra_expression.Symbol("x"),
				),
			),
		),
	)

	var result float64 = equation.Evaluate(x)

	if result != expected {
		t.Errorf("\nexpected: %.25f\ngot     : %.25f", expected, result)
	}

	t.Logf("finished testing Evaluate for Equation\n\n")
}

func TestStringEquation(t *testing.T) {
	t.Logf("testing String for Equation\n\n")
	var signature string = "f(x)"

	var expected string = fmt.Sprintf("%s = 0.5x^2 +(cos(x) * sin(x))", signature)

	var equation *algebra_equation.Equation = algebra_equation.NewEquation(signature).SetExpression(
		algebra_expression.Sum(
			algebra_expression.Pow(
				algebra_expression.Symbol("x"),
				algebra_expression.Int(2),
			).Divide(
				algebra_expression.Int(2),
			),
			algebra_expression.Multiply(
				algebra_expression.Cos(
					algebra_expression.Symbol("x"),
				),
				algebra_expression.Sin(
					algebra_expression.Symbol("x"),
				),
			),
		),
	)

	var result string = equation.String()

	if result != expected {
		t.Errorf("\nexpected: %s\ngot     : %s", expected, result)
	}

	t.Logf("finished testing String for Equation\n\n")
}
