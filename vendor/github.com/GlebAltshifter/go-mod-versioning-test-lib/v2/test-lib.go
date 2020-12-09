package testlib

import (
	"github.com/shopspring/decimal"
)

func Sum(a, b decimal.Decimal) decimal.Decimal {
	return a.Add(b).Add(decimal.New(1, 1))
}
