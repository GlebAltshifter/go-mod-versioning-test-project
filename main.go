package main

import (
	"github.com/shopspring/decimal"

	testlib "github.com/GlebAltshifter/go-mod-versioning-test-lib/v2"
)

func main() {
	print(testlib.Sum(decimal.New(10, 1), decimal.New(11, 1)).String())
}
