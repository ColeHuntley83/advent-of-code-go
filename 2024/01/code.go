package main

import (
	"math"
	"sort"
	"strconv"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

func run(part2 bool, input string) any {
	sum := 0
	simScore := 0
	leftList := []int{}
	rightList := []int{}
	leftMap := map[int]int{}
	rightMap := map[int]int{}
	for _, line := range strings.Split(strings.TrimSpace(input), "\n") {
		pair := strings.SplitN(line, " ", 2)
	
		left, right := strings.TrimSpace(pair[0]), strings.TrimSpace(pair[1])
		leftInt , _ := strconv.Atoi(left)
		rightInt , _ := strconv.Atoi(right)
		leftList = append(leftList, leftInt)
		rightList = append(rightList, rightInt)
	}
	sort.Ints(leftList)
	sort.Ints(rightList)
	
	for idx := range leftList {
		if leftVal, ok := leftMap[leftList[idx]]; ok {
			leftMap[leftList[idx]] = leftVal + 1
		} else {
			leftMap[leftList[idx]] = 1
		}
		if rightVal, ok := rightMap[rightList[idx]]; ok {
			rightMap[rightList[idx]] = rightVal + 1
		} else {
			rightMap[rightList[idx]] = 1
		}
	}


	for idx := range leftList  {

		simScore += leftList[idx] * rightMap[leftList[idx]]
		sum += int(math.Abs(float64(rightList[idx] - leftList[idx])))
	}
	//fmt.Println(rightMap)
	if part2 {
		return simScore
	}

	return sum
}
