package pkg

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/kthanasio/golang-server-api/internal/entity"
)

var (
	apiMaxTime = 200
)

// GetTodayCurrency will call an external API to ge data
func GetTodayCurrency(ctx context.Context, endpoint string) (*entity.Currency, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*time.Duration(apiMaxTime))
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, "GET", endpoint, nil)
	if err != nil {
		return nil, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	res, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var curr *entity.CurrencyExternal
	err = json.Unmarshal(res, &curr)
	if err != nil {
		return nil, err
	}
	valor, err := strconv.ParseFloat(curr.Usd2Brl.Bid, 64)
	if err != nil {
		panic(err)
	}
	currency := entity.NewCurrency(time.Now().Format("2006-01-02"), valor)
	return currency, nil
}
