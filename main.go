package main

import (
	testlib "github.com/GlebAltshifter/go-mod-versioning-test-lib/v2"
	"github.com/shopspring/decimal"
)

func main() {
	print(testlib.Sum(decimal.New(10, 1), decimal.New(11, 1)).String())
}
