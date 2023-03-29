package logrequest

import (
	"github.com/dhnikolas/mock/internal/dto"
	"gorm.io/gorm"
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

func (r *Repository) DeleteByMockId (mockId string) (int64, error) {
	result := r.db.Where("mock_id = ?", mockId).Delete(&dto.LogRequest{})
	return result.RowsAffected , result.Error
}

func (r *Repository) DeleteLog (mockId, logId string) (bool, error) {
	result := r.db.Where("mock_id = ? AND id = ?", mockId, logId).Delete(&dto.LogRequest{})
	return result.RowsAffected > 0, result.Error
}