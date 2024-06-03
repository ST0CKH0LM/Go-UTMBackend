package usecase

import (
	"gitlab.com/Std217/test/repositories"
	"gitlab.com/Std217/test/serializers"
)

type UsersUsecase struct {
	usersRepo repositories.UsersRepository
}

func NewUsersUsecase(repo repositories.UsersRepository) *UsersUsecase {
	return &UsersUsecase{usersRepo: repo}
}

func (uc *UsersUsecase) GetAllUser() ([]serializers.UsersResponse, error) {
	return uc.usersRepo.GetAllUser()
}
