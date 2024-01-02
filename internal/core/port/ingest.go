package port

import (
	"github.com/GeekGawd/logtracer/internal/core/domain"
)

type IngestionPort interface {
	Insert(domain.LoggerData) error
	BulkInsert([]domain.LoggerData) error
	Query(domain.LogQuery) ([]domain.LoggerData, error)
}