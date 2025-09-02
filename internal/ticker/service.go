package ticker

import "time"

// TickerSerivce manages the ticker schedule
type TickerService struct {
	interval time.Duration
}

// NewTickerService creates a new TickerService
func NewTickerService() *TickerService {
	return &TickerService{}
}
