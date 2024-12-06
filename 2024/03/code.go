package main

import (
	"fmt"
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
	sum := 0
	if part2 {
		return "not implemented"
	}
	const searchStr = "mul("
	const searchEndStr = ")"
	// scan text and parse out only text that starts with "mul(" and end match with closing ")"
	// take substring and split on "," and then check to make sure we only have 2 substrings
	// also need to check that each substring only is a valid number with no more than 3 digits
	// if valid then multiply and add to sum.
	// else ignore
	for _, line := range strings.Split(strings.TrimSpace(input), "\n") {
		end := false
		for {
			fmt.Println("full line: ", line)
			if end {
				break
			}
			index := strings.Index(line, searchStr)
			if index != -1 {
				//found a potential start of match
				
				endIndex := strings.Index(line, searchEndStr)
				if endIndex != -1 && endIndex > index {
					fmt.Println("start index: ", index)
					fmt.Println("end index: ", endIndex)
					// found a match
					//split into substrings and check valid 3 digit numbers
					prob := line[strings.Index(line, "mul(")+ 4 : endIndex]
					fmt.Println("full prob after start and end slice: ", prob)
					pair := strings.SplitN(prob, ",", 2)
					if len(pair) != 2 {
						line = line[index + 4:]
						continue
					}
					arg1 := pair[0]
					arg2 := pair[1]
					fmt.Println(arg1, arg2)
					valid, pairSum := checkPairMult(prob)
					if !valid {
						fmt.Println("invalid operation found: \n", prob)
						fmt.Println("line before slice: \n", line)
						fmt.Println("index: ", index)
						line = line[index + 4:]
						fmt.Println("line now set to: \n", line)

					} else {
					fmt.Println("sum after check: ", pairSum)
					sum += pairSum
					line = line[endIndex + 1:]
					fmt.Println("line before ")
					fmt.Println("line after start and end found: \n",  line)
					//pair := strings.SplitN(line[index:], " ", 20)
					}
				} else {
					fmt.Println("found start but could not find end or end index is smaller than start")
					line = line[index:]
					fmt.Println(line)
				}
			} else {
				end = true

			}


		}
	}

	return sum
}

func checkPairMult(prob string) (bool,  int) {
	pair := strings.SplitN(prob, ",", 2)
	arg1 := pair[0]
	arg2 := pair[1]
	fmt.Println("arg 1: ", arg1, len(arg1))
	fmt.Println("arg 2: ", arg2, len(arg2))

	if len(arg1) > 3 || len(arg2) > 3 {
		return  false, 0
	}
	 arg1Num, err := strconv.Atoi(arg1)  
	  if  err != nil {

		return false, 0
	}
	arg2Num, err := strconv.Atoi(arg2)
	if  err != nil {
		return false ,0
	}
	fmt.Println("args after str to int: ", arg1Num, arg2Num)
	return true,  arg1Num * arg2Num

}
