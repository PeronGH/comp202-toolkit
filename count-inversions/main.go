package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/pterm/pterm"
)

func mergeSort(arr, tempArr []int, left, right int) ([][]int, int) {
	inversions := [][]int{}
	invCount := 0

	if left < right {
		mid := (left + right) / 2

		leftInversions, leftCount := mergeSort(arr, tempArr, left, mid)
		rightInversions, rightCount := mergeSort(arr, tempArr, mid+1, right)

		inversions = append(inversions, leftInversions...)
		inversions = append(inversions, rightInversions...)
		invCount = leftCount + rightCount

		mergeInversions, mergeCount := merge(arr, tempArr, left, mid, right)
		inversions = append(inversions, mergeInversions...)
		invCount += mergeCount
	}

	return inversions, invCount
}

func merge(arr, tempArr []int, left, mid, right int) ([][]int, int) {
	i := left
	j := mid + 1
	k := left
	invCount := 0
	inversions := [][]int{}

	for i <= mid && j <= right {
		if arr[i] <= arr[j] {
			tempArr[k] = arr[i]
			i++
		} else {
			tempArr[k] = arr[j]
			invCount += mid - i + 1
			for p := i; p <= mid; p++ {
				inversions = append(inversions, []int{arr[p], arr[j]})
			}
			j++
		}
		k++
	}

	for i <= mid {
		tempArr[k] = arr[i]
		i++
		k++
	}

	for j <= right {
		tempArr[k] = arr[j]
		j++
		k++
	}

	for i := left; i <= right; i++ {
		arr[i] = tempArr[i]
	}

	return inversions, invCount
}

func countInversions(arr []int) ([][]int, int) {
	n := len(arr)
	tempArr := make([]int, n)
	return mergeSort(arr, tempArr, 0, n-1)
}

func main() {
	input, _ := pterm.DefaultInteractiveTextInput.
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

	fmt.Println("Input:", arr)
	inversions, count := countInversions(arr)
	fmt.Println("Inversions:", inversions)
	fmt.Println("Count:", count)
}
