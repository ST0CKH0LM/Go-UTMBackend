package repositories

import (
	"gitlab.com/Std217/test/model"
	"gitlab.com/Std217/test/serializers"
	"gorm.io/gorm"
)

type UsersRepository struct {
	DB *gorm.DB
}

func NewUsersRepository(db *gorm.DB) *UsersRepository {
	return &UsersRepository{DB: db}
}

func (repo *UsersRepository) GetAllUser() ([]serializers.UsersResponse, error) {
	var users []model.Users
	if err := repo.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	var serializedUsers []serializers.UsersResponse
	for _, user := range users {
		serializedUsers = append(serializedUsers, serializers.NewUsersResponse(user))
	}
	return serializedUsers, nil
}
