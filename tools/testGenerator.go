package tools

import (
	"os"
	"strings"
)

func TestCaseGenGoodBoy(n int, goodBoy bool,invalidCase bool) {
	var sb strings.Builder
	for i := 0; i < n; i++ {
		sb.WriteString("SR")
	}
	if goodBoy != true {
		sb.WriteString("S")
	}
	if invalidCase == true {
		sb.WriteString("1กขABS456SR")
	}
	os.WriteFile("./testCase", []byte(sb.String()), 0666)
}
