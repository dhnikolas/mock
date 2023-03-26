package logrequest

import (
	"gorm.io/gorm"
	"mock/internal/dto"
)

type Repository struct {
	db *gorm.DB
}
func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Add (mock *dto.LogRequest) error {
	result := r.db.Create(mock)
	return result.Error
}

func (r *Repository) GetByMockId (mockId string) (*dto.LogRequests, error) {
	logs := &dto.LogRequests{}
	result := r.db.Order("created_at desc").Find(logs, "mock_id = ?", mockId)
	return logs, result.Error
}