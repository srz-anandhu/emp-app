package repository

import (
	"emp-app/app/domain"

	"gorm.io/gorm"
)

type AdminRepo interface {
	FindAdminByEmail(email string) (*domain.Admin, error)
}

type AdminRepoImpl struct {
	db *gorm.DB
}

func NewAdminRepo(db *gorm.DB) AdminRepo {
	return &AdminRepoImpl{
		db: db,
	}
}

// For checking the implementation of AdminRepo interface
var _ AdminRepo = (*AdminRepoImpl)(nil)

func (r *AdminRepoImpl) FindAdminByEmail(email string) (*domain.Admin, error) {
	admin := &domain.Admin{}
	result := r.db.Table("admins").Where("email = ?", email).First(admin)
	if result.Error != nil {
		return nil, result.Error
	}
	return admin, nil
}
