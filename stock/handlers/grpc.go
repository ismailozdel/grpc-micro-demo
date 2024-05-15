package handlers

import (
	"context"

	"github.com/ismailOZdel/micro/stock/types"
	"github.com/ismailozdel/micro/common/proto/stock"
	"google.golang.org/grpc"
)

type StockGrpcHandler struct {
	stockService types.StockServiceClient
	stock.UnimplementedStockServiceServer
}

func NewGrpcStockService(grpc *grpc.Server, stockService types.StockServiceClient) {
	grpcHandler := &StockGrpcHandler{
		stockService: stockService,
	}
	stock.RegisterStockServiceServer(grpc, grpcHandler)
}

func (s *StockGrpcHandler) GetStocks(ctx context.Context, req *stock.GetStocksRequest) (*stock.GetStocksResponse, error) {
	stocks, err := s.stockService.GetStocks(ctx)
	if err != nil {
		return nil, err
	}
	return &stock.GetStocksResponse{Stocks: stocks}, nil
}

func (s *StockGrpcHandler) CreateStock(ctx context.Context, req *stock.CreateStockRequest) (*stock.CreateStockResponse, error) {
	stocks, err := s.stockService.CreateStock(ctx, &stock.Stock{Symbol: req.Symbol, Name: req.Name, Price: req.Price})
	if err != nil {
		return nil, err
	}
	return &stock.CreateStockResponse{Stock: stocks}, nil
}

func (s *StockGrpcHandler) GetStock(ctx context.Context, req *stock.GetStockRequest) (*stock.GetStockResponse, error) {
	stocks, err := s.stockService.GetStock(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	if len(stocks) == 0 {
		return &stock.GetStockResponse{Stock: nil}, nil
	}
	return &stock.GetStockResponse{Stock: stocks[0]}, nil
}
