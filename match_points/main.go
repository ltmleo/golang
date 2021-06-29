package main
import (
	"fmt"
)

var newResultsString string

func main() {
	fmt.Scan(&newResultsString)

    results := []string{"w", "l", "w", "d", "w", "l", "l", "l", "d", "d", "w", "l", "w", "d"}
	results = append(results, newResultsString)
    pointsTable := map[string]int{"w": 3, "d": 1, "l": 0}
	points := 0
	for _, k := range results {
		points += pointsTable[k]
	}
	fmt.Println(points)
}
