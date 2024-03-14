package main

import (
	"container/heap"
	"fmt"
	"strconv"
	"strings"

	"github.com/pterm/pterm"
)

type IntHeap struct {
	values []int
	isMax  bool
}

func (h IntHeap) Len() int {
	return len(h.values)
}

func (h IntHeap) Less(i, j int) bool {
	if h.isMax {
		return h.values[i] > h.values[j]
	}
	return h.values[i] < h.values[j]
}

func (h IntHeap) Swap(i, j int) {
	h.values[i], h.values[j] = h.values[j], h.values[i]
}

func (h *IntHeap) Push(x any) {
	h.values = append(h.values, x.(int))
}

func (h *IntHeap) Pop() any {
	old := h.values
	n := len(old)
	x := old[n-1]
	h.values = old[0 : n-1]
	return x
}

func (h IntHeap) String() string {
	return fmt.Sprintf("%v", h.values)
}

func NewIntHeap(isMax bool, values []int) *IntHeap {
	return &IntHeap{isMax: isMax, values: values}
}

func main() {
	input, _ := pterm.DefaultInteractiveSelect.
		WithDefaultText("Select heap type").
		WithOptions([]string{"Max Heap", "Min Heap"}).
		Show()

	isMax := input == "Max Heap"

	input, _ = pterm.DefaultInteractiveTextInput.
		WithDefaultText("Input elements (separated by space or comma)").
		Show()

	// Split the input string by space or comma
	tokens := strings.FieldsFunc(input, func(r rune) bool {
		return r == ' ' || r == ','
	})

	// Convert the tokens to integers
	arr := make([]int, len(tokens))
	for i, token := range tokens {
		num, err := strconv.Atoi(token)
		if err != nil {
			fmt.Printf("Invalid input: %s\n", token)
			return
		}
		arr[i] = num
	}

	h := NewIntHeap(isMax, arr)
	heap.Init(h)

	for {
		fmt.Println("Current heap:", h)
		input, _ := pterm.DefaultInteractiveSelect.
			WithDefaultText("Select operation").
			WithOptions([]string{"Pop", "Push"}).
			Show()

		switch input {
		case "Pop":
			if h.Len() == 0 {
				fmt.Println("Heap is empty")
				continue
			}
			fmt.Println("Popped value:", heap.Pop(h))
		case "Push":
			input, _ := pterm.DefaultInteractiveTextInput.
				WithDefaultText("Enter the value to push").
				Show()

			num, err := strconv.Atoi(input)
			if err != nil {
				fmt.Printf("Invalid input: %s\n", input)
				return
			}
			heap.Push(h, num)
		}
	}
}
