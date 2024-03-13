package main

import (
	"fmt"
	"sort"

	"github.com/pterm/pterm"
)

type Item struct {
	weight  float64
	benefit float64
}

func fractionalKnapsack(items []Item, capacity float64) (fractions []float64, benefit float64) {
	// Create a slice of indices to store the original order of items
	indices := make([]int, len(items))
	for i := range indices {
		indices[i] = i
	}

	// Sort indices based on benefit/weight ratio in descending order
	sort.Slice(indices, func(i, j int) bool {
		return (items[indices[i]].benefit / items[indices[i]].weight) > (items[indices[j]].benefit / items[indices[j]].weight)
	})

	var totalWeight float64
	fractions = make([]float64, len(items))

	for _, index := range indices {
		if totalWeight+items[index].weight <= capacity {
			// Take the whole item
			fractions[index] = 1
			totalWeight += items[index].weight
			benefit += items[index].benefit
		} else {
			// Take a fraction of the item
			remainingCapacity := capacity - totalWeight
			fractions[index] = remainingCapacity / items[index].weight
			benefit += fractions[index] * items[index].benefit
			break
		}
	}

	return fractions, benefit
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

	fractions, benefit := fractionalKnapsack(items, capacity)
	table, _ := renderTable(items, fractions)
	area.Update(table)
	pterm.DefaultBasicText.Printf("Total Benefit: %.2f\n", benefit)
}

func renderTable(items []Item, fractions []float64) (string, error) {
	table := pterm.DefaultTable.
		WithHasHeader().
		WithBoxed().
		WithData([][]string{
			{"Benefit", "Weight", "Fraction"},
		})

	for i, item := range items {
		fraction := "N/A"
		if fractions != nil && i < len(fractions) {
			fraction = fmt.Sprintf("%.2f", fractions[i])
		}

		table.Data = append(table.Data, []string{
			fmt.Sprintf("%.2f", item.benefit),
			fmt.Sprintf("%.2f", item.weight),
			fraction,
		})
	}

	return table.Srender()
}
