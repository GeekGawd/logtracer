package service

import (
	"github.com/GeekGawd/logtracer/internal/core/domain"
	"github.com/GeekGawd/logtracer/internal/core/port"
)

type IngestionService struct {
	port port.IngestionPort
}

func NewIngestionService(port port.IngestionPort) *IngestionService {
	return &IngestionService{port: port}
}

func (a *IngestionService) Insert(data domain.LoggerData) error {
	err := a.port.Insert(data);
	if err != nil {
		return err
	}
	return nil
}

func (a *IngestionService) BulkInsert(data []domain.LoggerData) error {
	err := a.port.BulkInsert(data);
	if err != nil {
		return err
	}
	return nil
}

func (a *IngestionService) Query(query domain.LogQuery) ([]domain.LoggerData, error) {
	data, err := a.port.Query(query);
	if err != nil {
		return nil, err
	}
	return data, nil
}