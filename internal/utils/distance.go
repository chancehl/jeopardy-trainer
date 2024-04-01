package utils

import "math"

// levenshtein calculates the Levenshtein distance between two strings.
func Levenshtein(correctAnswer string, userAnswer string) (int, int) {
	m, n := len(correctAnswer), len(userAnswer)

	allowed := ComputeAllowableDistance(correctAnswer)

	if m == 0 {
		return n, allowed
	}

	if n == 0 {
		return m, allowed
	}

	// Create a matrix.
	d := make([][]int, m+1)
	for i := range d {
		d[i] = make([]int, n+1)
	}

	// Initialize the first row and column of the matrix.
	for i := 0; i <= m; i++ {
		d[i][0] = i
	}
	for j := 0; j <= n; j++ {
		d[0][j] = j
	}

	// Populate the matrix.
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			cost := 0
			if correctAnswer[i-1] != userAnswer[j-1] {
				cost = 1
			}
			d[i][j] = min(
				d[i-1][j]+1,      // Deletion
				d[i][j-1]+1,      // Insertion
				d[i-1][j-1]+cost, // Substitution
			)
		}
	}

	return d[m][n], allowed
}

func ComputeAllowableDistance(s string) int {
	return int(math.Round(float64(len(s)) * 0.25))
}
