package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {

	fmt.Println("T(n) = aT(n/b) + Θ(n^k log^i n)")
	fmt.Print("Enter a, b, k, i (separated by spaces): ")
	var aStr, bStr, kStr, iStr string
	fmt.Scanln(&aStr, &bStr, &kStr, &iStr)

	a, err := strconv.ParseFloat(aStr, 64)
	if err != nil {
		fmt.Println("Invalid value for a")
		os.Exit(1)
	}

	b, err := strconv.ParseFloat(bStr, 64)
	if err != nil {
		fmt.Println("Invalid value for b")
		os.Exit(1)
	}

	k, err := strconv.ParseFloat(kStr, 64)
	if err != nil {
		fmt.Println("Invalid value for k")
		os.Exit(1)
	}

	i, err := strconv.ParseFloat(iStr, 64)
	if err != nil {
		fmt.Println("Invalid value for i")
		os.Exit(1)
	}

	if a < 0 {
		fmt.Println("a must be non-negative")
		os.Exit(1)
	}
	if b <= 1 {
		fmt.Println("b must be greater than 1")
		os.Exit(1)
	}
	if k < 0 {
		fmt.Println("k must be at least 0")
		os.Exit(1)
	}
	if i < 0 {
		fmt.Println("i must be at least 0")
		os.Exit(1)
	}

	recurrenceText := fmt.Sprintf("T(n) = %sT(n%s) + Θ(%s)",
		formatCoefficient(a), formatDivision(b), formatPolyLog(k, i))
	fmt.Println("Recurrence:", recurrenceText)

	p := math.Log(a) / math.Log(b)
	var result string
	switch {
	case floatEquals(p, k):
		result = fmt.Sprintf("Θ(%s)", formatPolyLog(k, i+1))
	case p < k:
		result = fmt.Sprintf("Θ(%s)", formatPolyLog(k, i))
	case p > k:
		if floatEquals(math.Round(p), p) {
			result = fmt.Sprintf("Θ(n^%s)", formatFloat(p))
		} else {
			result = fmt.Sprintf("Θ(n^(log_%s %s))", formatFloat(b), formatFloat(a))
		}
	default:
		result = "Arithmetic error"
	}
	fmt.Println("Solution:", result)
}

func formatCoefficient(a float64) string {
	if a != 1 {
		return formatFloat(a) + " "
	}
	return ""
}

func formatDivision(b float64) string {
	if b != 1 {
		return " / " + formatFloat(b)
	}
	return ""
}

func formatPolyLog(k, i float64) string {
	var result string
	switch {
	case k == 0 && i != 0:
		result = ""
	case k == 0 && i == 0:
		result = "1"
	case k == 0.5:
		result = "√n"
	case k == 1:
		result = "n"
	default:
		result = "n^" + formatFloat(k)
	}

	if i != 0 {
		if result != "" {
			result += " "
		}
		result += "log"
		if i != 1 {
			result += "^" + formatFloat(i)
		}
		result += " n"
	}

	return result
}

func formatFloat(value float64) string {
	return strconv.FormatFloat(value, 'f', -1, 64)
}

func floatEquals(x, y float64) bool {
	return math.Abs(x-y) < 1e-9
}
