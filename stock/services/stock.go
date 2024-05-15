package services

import (
	"context"
	"fmt"

	"github.com/ismailozdel/micro/common/proto/stock"
)

var stockList = []*stock.Stock{
	{
		Symbol: "AAPL",
		Id:     "1",
		Name:   "Apple",
		Price:  100,
	},
	{
		Symbol: "MSFT",
		Id:     "2",
		Name:   "Microsoft",
		Price:  200,
	},
}

type StockService struct {
}

func NewStockService() *StockService {
	return &StockService{}
}

func (s *StockService) GetStocks(ctx context.Context) ([]*stock.Stock, error) {
	return stockList, nil
}

func (s *StockService) CreateStock(ctx context.Context, u *stock.Stock) ([]*stock.Stock, error) {
	u.Id = fmt.Sprintf("%d", len(stockList)+1)
	stockList = append(stockList, u)
	return stockList, nil
}
func (s *StockService) GetStock(ctx context.Context, id string) ([]*stock.Stock, error) {
	for _, i := range stockList {
		if i.Id == id {
			return []*stock.Stock{i}, nil
		}
	}

	return nil, nil
}

/*
type StockServiceClient interface {
	GetStocks(context.Context) ([]*stock.Stock, error)
	CreateStock(context.Context, *stock.Stock) ([]*stock.Stock, error)
}

*/
