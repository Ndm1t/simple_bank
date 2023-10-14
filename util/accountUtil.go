package util

import (
	"math"
)

const EURToUSD = 1.05

func ConvertBalance(currFrom string, balance float64) (convertedBalance float64) {
	if currFrom == "USD" {
		convertedBalance = balance / EURToUSD
	} else {
		convertedBalance = balance * EURToUSD
	}
	return math.Round(convertedBalance*100) / 100
}
