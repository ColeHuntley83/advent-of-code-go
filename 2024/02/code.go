package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

// on code change, run will be executed 4 times:
// 1. with: false (part1), and example input
// 2. with: true (part2), and example input
// 3. with: false (part1), and user input
// 4. with: true (part2), and user input
// the return value of each run is printed to stdout
func run(part2 bool, input string) any {
	// when you're ready to do part 2, remove this "not implemented" block
	safeCount := 0
	if part2 {
		return "not implemented"
	}
	// solve part 1 here
	for _, line := range strings.Split(strings.TrimSpace(input), "\n") {
		pair := strings.SplitN(line, " ", 20)
		levels := make([]int, len(pair))
		for i, level := range pair {
			n, err := strconv.Atoi(level)
			if err != nil {
				fmt.Println(err)
				return nil
			}
			levels[i] = n
		}

		fmt.Println(pair)
		safe := checkSafe(levels)
		if safe {
			safeCount += 1
		} else {
			remove1Safe := removeAndCheck(levels)
			if remove1Safe {
				safeCount += 1
			}
		}

		fmt.Println(safe)
	}
	return safeCount
}


func removeAndCheck(levels []int)bool {
	length := len(levels)

	for idx := range length {
		newLevels := []int{}
		for i, val := range levels {
			if  i == idx {
				continue
			}
			newLevels = append(newLevels, val)
		}
		fmt.Println(newLevels)
		safe := checkSafe(newLevels)
		fmt.Println("here: ", safe)
		if safe {
			return true
		}
	}
	return false


}

func checkSafe(levels []int) bool {

	decreasingOrder := false

	//check if in decreasing order
	if levels[0] <= levels[1] {
		decreasingOrder = true
	}
	for i := 0; i < len(levels)-1; i++ {
		if levels[i] == levels[i+1] || int(math.Abs(float64(levels[i] - levels[i + 1]))) > 3  {
			return false
		}
		if decreasingOrder && levels[i] > levels[i+1] {
			return false
		}
		if !decreasingOrder && levels[i] < levels[i+1] {
			return false
		}
	}

	return true

}
