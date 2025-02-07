/*

	Running this program:
		In your terminal, invoke the following command : go run main.go

------------------------------------------------------------------------------

	Resources used:
		https://www.geeksforgeeks.org/math-package-in-golang/
		https://pkg.go.dev/math
*/

package main

import (
	"fmt"
	"math"
)

// Function to calculate the Pythagorean Means
func calculateMeans(start, end float64) (float64, float64, float64) {
	// Initialize variables for sum, reciprocal sum, and product
	sum, reciprocalSum, product := 0.0, 0.0, 1.0

	// Loop through the numbers from `start` to `end`
	for n := start; n <= end; n++ {
		sum += n                 // Calculate the sum
		reciprocalSum += 1 / n   // Calculate the sum of reciprocals
		product *= n             // Calculate the product
	}

	// Calculate Arithmetic Mean (A)
	a := sum / (end - start + 1)

	// Calculate Geometric Mean (G)
	g := math.Pow(product, 1/(end-start+1))

	// Calculate Harmonic Mean (H)
	h := (end - start + 1) / reciprocalSum

	return a, g, h
}

func main() {
	// Specify the range of numbers
	start, end := 1.0, 10.0

	// Calculate the means
	a, g, h := calculateMeans(start, end)

	// Print the results
	fmt.Printf("Arithmetic Mean: %.4f \nGeometric Mean: %.4f \nHarmonic Mean: %.4f\n", a, g, h)

	// Verify and print the inequality A >= G >= H
	fmt.Println("A >= G >= H:", a >= g && g >= h)
}
