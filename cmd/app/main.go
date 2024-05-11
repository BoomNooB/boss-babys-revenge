package main

import (
	"fmt"
	"p1/constant"
	"strings"
)

func main() {
	testCases := []string{
		constant.TC0,
		constant.TC1,
		constant.TC2,
		constant.TC3,
		constant.TC4,
		constant.TC5,
		constant.TC6,
		constant.TC7,
		constant.TC8,
	}

	for _, testCase := range testCases {
		if isGoodBoy(testCase) {
			printGoodBoy()
		} else {
			printBadBoy()
		}
	}
}

func isNoShotHappen(tc string) bool {
	// since there's no shooting happen so good boy will be an answer
	return tc == ""
}

func isNotRevengeFirst(tc string) bool {
	// if first shot if got shot -> good boy -> true
	// if first shot if shot -> bad boy -> false
	return strings.HasPrefix(tc, constant.GotShot)
}

func isLastShotNotIsGotShot(tc string) bool {
	// if last shot is s -> bad boy -> false
	return strings.HasSuffix(tc, constant.Revenge)
}

func isEachShotWasRevenge(tc string) bool {
	notRevengeYetShot := 0
	for _, shot := range tc {
		// log.Printf("Current not revenge shot %d, next shot: %s", notRevengeYetShot, string(shot))
		switch string(shot) {
		case constant.GotShot:
			notRevengeYetShot++
		case constant.Revenge:
			if notRevengeYetShot != 0 {
				notRevengeYetShot--
			}
		}
	}
	// log.Printf("Shot that not revenge yet is: %d", notRevengeYetShot)
	return notRevengeYetShot == 0
}

func printGoodBoy() {
	fmt.Println(constant.GoodBoy)
}

func printBadBoy() {
	fmt.Println(constant.BadBoy)
}

func isGoodBoy(testCase string) bool {
	// true -> good boy
	// false -> bad boy
	if isNoShotHappen(testCase) {
		return true
	}
	if !isNotRevengeFirst(testCase) {
		return false
	}
	if !isLastShotNotIsGotShot(testCase) {
		return false
	}
	if isEachShotWasRevenge(testCase) {
		return true
	}
	return false
}
