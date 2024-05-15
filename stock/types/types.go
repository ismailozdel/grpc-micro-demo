package types

import (
	"context"

	"github.com/ismailozdel/micro/common/proto/stock"
)

type StockServiceClient interface {
	GetStocks(context.Context) ([]*stock.Stock, error)
	CreateStock(context.Context, *stock.Stock) ([]*stock.Stock, error)
	GetStock(context.Context, string) ([]*stock.Stock, error)
}
