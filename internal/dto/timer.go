package dto

import "time"

type Tick struct {
	Ticker    *time.Ticker
	Executing bool
}
