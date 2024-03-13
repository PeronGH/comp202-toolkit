package main

import (
	"fmt"

	"github.com/pterm/pterm"
)

type Item struct {
	weight  float64
	benefit float64
}

// knapsack solves the 0/1 Knapsack problem using dynamic programming.
func knapsack(items []Item, capacity float64) (taken []bool, totalBenefit float64) {
	n := len(items)
	// Create a 2D slice to store the maximum benefit of knapsack capacity j with first i items.
	dp := make([][]float64, n+1)
	for i := range dp {
		dp[i] = make([]float64, int(capacity)+1)
	}

	// Build table dp[][] in bottom-up manner.
	for i := 1; i <= n; i++ {
		for w := 1; w <= int(capacity); w++ {
			if items[i-1].weight <= float64(w) {
				// Item i can be included.
				dp[i][w] = max(dp[i-1][w], items[i-1].benefit+dp[i-1][w-int(items[i-1].weight)])
			} else {
				// Item i cannot be included.
				dp[i][w] = dp[i-1][w]
			}
		}
	}

	// The maximum benefit is at the bottom-right corner of the matrix.
	totalBenefit = dp[n][int(capacity)]

	// Trace back to find which items to take.
	taken = make([]bool, n)
	w := int(capacity)
	for i := n; i > 0 && totalBenefit > 0; i-- {
		if totalBenefit != dp[i-1][w] {
			taken[i-1] = true
			totalBenefit -= items[i-1].benefit
			w -= int(items[i-1].weight)
		}
	}

	return taken, dp[n][int(capacity)]
}

func main() {
	var items []Item
	area, _ := pterm.DefaultArea.WithFullscreen().WithCenter().Start()
	defer area.Stop()

	for {
		var weight, benefit float64

		input, _ := pterm.DefaultInteractiveTextInput.WithDefaultText("Enter benefit and weight (separated by space) or 'q' to stop").Show()
		if input == "q" {
			break
		}

		fmt.Sscanf(input, "%f %f", &benefit, &weight)
		items = append(items, Item{weight, benefit})

		table, _ := renderTable(items, nil)
		area.Update(table)
	}

	var capacity float64
	input, _ := pterm.DefaultInteractiveTextInput.WithDefaultText("Enter the capacity of the knapsack").Show()
	fmt.Sscanf(input, "%f", &capacity)

	taken, benefit := knapsack(items, capacity)
	table, _ := renderTable(items, taken)
	area.Update(table)
	pterm.DefaultBasicText.Printf("Total Benefit: %.2f\n", benefit)
}

func renderTable(items []Item, fractions []bool) (string, error) {
	table := pterm.DefaultTable.
		WithHasHeader().
		WithBoxed().
		WithData([][]string{
			{"Benefit", "Weight", "Taken"},
		})

	for i, item := range items {
		fraction := "N"
		if fractions != nil && i < len(fractions) && fractions[i] {
			fraction = "Y"
		}

		table.Data = append(table.Data, []string{
			fmt.Sprintf("%.2f", item.benefit),
			fmt.Sprintf("%.2f", item.weight),
			fraction,
		})
	}

	return table.Srender()
}
