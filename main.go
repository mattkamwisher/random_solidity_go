package main

import (
	"fmt"
	"math/big"
	"time"
)

//Solidity Constants --- TODO move to another file

const (
	SECONDS = 1
	MINUTES = 60 * SECONDS
	HOURS   = 60 * MINUTES
	DAYS    = 24 * HOURS
	WEEKS   = 7 * DAYS
	YEARS   = 365 * DAYS
)

//TODO now is a function in solidity
//now (uint): current block timestamp (alias for block.timestamp)
//for now we can hard code it to make tests easier
func now() int64 {
	return time.Now().Unix() // 123
}

///

type StandardToken struct {
}

type Ownable struct {
}

//Note public variables in go have first letter upcased
type APRInflationToken struct {
	StandardToken //Embedded struct
	Ownable

	// Date control variables
	StartDate   *big.Int
	DailyAdjust *big.Int

	// Inflation controlling variables
	StartRate  *big.Int
	EndRate    *big.Int
	RateAdjust *big.Int
	Rate       *big.Int
}

/**
* @dev Avoids the daily adjust to run more than necessary
 */
//attrs: modifier
func (a *APRInflationToken) canAdjustDaily() {
	day := big.NewInt(1 * DAYS) // 1 day in seconds

	if big.NewInt(now()).Cmp(a.StartDate.Add(a.StartDate, day.Mul(day, a.DailyAdjust))) >= 0 {
		panic("no good") //TODO panic is not idiomatic go, but should translate better
		//in theory we should return an error
	}
}

// Increment the daily adjust counter to avoids repeated adjusts
// in a day, also allows to adjusts a past day if it was skipped
func (a *APRInflationToken) setDailyAdjustControl() *big.Int {
	return a.DailyAdjust.Add(big.NewInt(0), big.NewInt(1))
}

/**
* @dev Adjusts all the necessary calculations in constructor
 */
func CreateAPRInflationToken() *APRInflationToken {

	return &APRInflationToken{
		StartDate:   big.NewInt(0),
		DailyAdjust: big.NewInt(0),
		// 365 / 10%
		StartRate: big.NewInt(3650),
		// 365 / 1%
		EndRate:    big.NewInt(36500),
		RateAdjust: big.NewInt(9),
		Rate:       big.NewInt(3650),
	}
}

func main() {
	fmt.Println("hello")

	a := CreateAPRInflationToken()
	fmt.Printf("created object %v\n", a)
}
