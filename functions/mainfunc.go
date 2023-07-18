package maski

import (
	"fmt"
	"math"
	"sort"
	"strconv"
)

func MainFunc(s string) []float64 {
	numbers, numStr, check := []float64{}, "", false
	for i := 0; i < len(s); i++ {
		if s[i] != '\n' {
			if (s[i] >= '0' && s[i] <= '9') || s[i] == 45 {
				numStr += string(s[i])
			} else if s[i] >= 0 && s[i] <= 32 {
				continue
			}
		} else {
			num, err := strconv.Atoi(numStr)
			if err != nil {
				check = true
			} else {
				numbers = append(numbers, float64(num))
				numStr = ""
			}
		}
	}
	if len(numStr) > 0 {
		num, err := strconv.Atoi(numStr)
		if err != nil {
			fmt.Println(err)
		}
		numbers = append(numbers, float64(num))
	}
	if check {
		fmt.Println(("Error! Some items not number!"))
	}
	if len(numbers) == 0 {
		fmt.Println("Error! No numbers detected!")
		return nil
	}
	sort.Float64s(numbers)
	PrintResult(numbers)
	return numbers
}

func PrintResult(numbers []float64) {
	average, median, variance := int(math.Round(Average(numbers))), int(math.Round(Median(numbers))), int(math.Round(Variance(numbers)))
	standardDeviation := int(math.Round(math.Sqrt(Variance(numbers))))
	fmt.Print("Average: ", average, "\nMedian: ", median, "\nVariance: ", variance, "\nStandard Deviation: ", standardDeviation, "\n")
}

func Average(numbers []float64) float64 {
	var sum float64
	for i := 0; i < len(numbers); i++ {
		sum += numbers[i]
	}
	return sum / float64(len(numbers))
}

func Median(numbers []float64) float64 {
	if len(numbers)%2 == 0 {
		first := numbers[len(numbers)/2-1]
		second := numbers[len(numbers)/2]
		return (first + second) / 2
	} else {
		return numbers[len(numbers)/2]
	}
}

func Variance(numbers []float64) float64 {
	average := Average(numbers)
	var d float64
	for i := 0; i < len(numbers); i++ {
		d += (numbers[i] - average) * (numbers[i] - average)
	}
	return float64(d / float64(len(numbers)))
}

func Atoi(s string) float64 {
	res := 0
	for i := 0; i < len(s); i++ {
		res *= 10
		res += int(s[i] - 48)
	}
	return float64(res)
}
