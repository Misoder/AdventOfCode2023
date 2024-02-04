package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Misoder/AdventOfCode2023/tools"
)

func findLeadingNum(s string) rune {
	numStrings := map[string]rune{
		"one":   '1',
		"two":   '2',
		"three": '3',
		"four":  '4',
		"five":  '5',
		"six":   '6',
		"seven": '7',
		"eight": '8',
		"nine":  '9',
	}

	if s[0] >= '0' && s[0] <= '9' {
		return rune(s[0])
	}

	for k, v := range numStrings {
		if len(s) >= len(k) && s[0:len(k)] == k {
			return v
		}
	}

	return 0
}

func buildNumbersSlice(s string) []rune {
	var nums []rune

	for i := 0; i < len(s); i++ {
		if n := findLeadingNum(s[i:]); n != 0 {
			nums = append(nums, n)
		}
	}

	return nums
}

func main() {
	t := time.Now()
	defer fmt.Printf("Duration: %s\n", time.Now().Sub(t))

	lineCh, err := tools.ReadLines("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	values := make([]string, 0)

	var sb strings.Builder
	for line := range lineCh {
		sb.Reset()

		numbers := buildNumbersSlice(line)

		if len(numbers) == 0 {
			continue
		}

		sb.WriteRune(numbers[0])
		sb.WriteRune(numbers[len(numbers)-1])
		values = append(values, sb.String())
	}

	sum := 0
	for _, value := range values {
		v, err := strconv.Atoi(value)
		if err != nil {
			fmt.Printf("failed to convert %s to int: %v\n", value, err)
			return
		}
		sum += v
	}

	fmt.Println(sum)
}
