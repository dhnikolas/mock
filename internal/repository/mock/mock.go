package mock

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

func (r *Repository) Add (mock *dto.Mock) error {
	result := r.db.Create(mock)
	return result.Error
}

func (r *Repository) GetAll () (*dto.Mocks, error) {
	mocks := &dto.Mocks{}
	result := r.db.Find(mocks)
	return mocks, result.Error
}

func (r *Repository) GetByUrl (url string) (*dto.Mocks, error) {
	mocks := &dto.Mocks{}
	result := r.db.Find(mocks, "url = ?", url)
	return mocks, result.Error
}

func (r *Repository) GetByUrlAndMethod (url, method string) (*dto.Mock, bool, error) {
	mock := &dto.Mock{}
	result := r.db.First(mock, "url = ? AND method = ?", url, method)
	
	return mock, result.RowsAffected > 0, result.Error
}

func (r *Repository) Delete (id string) (bool, error) {
	result := r.db.Where("id = ?", id).Delete(&dto.Mock{})
	return result.RowsAffected > 0, result.Error
}

func (r *Repository) Update (mock *dto.Mock) (bool, error) {
	result := r.db.Save(mock)
	return result.RowsAffected > 0, result.Error
}




