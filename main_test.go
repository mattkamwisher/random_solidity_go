package main

import (
	"math/big"
	"testing"
)

func Test1(t *testing.T) {
	a := CreateAPRInflationToken()
	//a.canAdjustDaily()

	if a.DailyAdjust.Cmp(big.NewInt(132)) == 0 {
		t.Error("Invalid daily adjust")
	}
}
