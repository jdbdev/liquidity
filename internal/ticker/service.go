package ticker

import (
	"net/http"
	"time"

	"github.com/jdbdev/liquidity/config"
)

// TickerSerivce manages the ticker schedule
type TickerService struct {

	//Core fields
	interval   time.Duration
	apiKey     string
	baseURL    string
	httpClient *http.Client
}

// NewTickerService creates a new TickerService
func NewTickerService(app *config.AppConfig) *TickerService {
	return &TickerService{}
}
