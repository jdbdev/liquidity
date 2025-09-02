package models

import "time"

// LiquidityPosition holds all the data for a single liquidity position at entry point.
type LiquidityPosition struct {
	BaseAsset     string
	QuoteAsset    string
	Chain         string
	Exchange      string
	BaseBalance   float64
	QuoteBalance  float64
	QuotePrice    float64 // price of base asset in quote asset
	Range         float64 // 0-100% range, half below, half above.
	UpperBound    float64 // upper bound of the range
	LowerBound    float64 // lower bound of the range
	StakingStatus bool    // whether the position is in a staking pool
}

// LiquidityPositions holds all the LP positions
type LiquidityPositions struct {
	Chain     string
	Positions []LiquidityPosition
}

type PriceData struct {
	CMCID     int
	Price     float64
	Timestamp time.Time
}

type Alert struct {
	ID        int
	Message   string
	Severity  string
	CreatedAt time.Time
}
