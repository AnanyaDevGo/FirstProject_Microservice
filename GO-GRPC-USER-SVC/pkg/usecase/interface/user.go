package interfaces

import (
	"github.com/AnanyaDevGo/pkg/domain"
	"github.com/AnanyaDevGo/pkg/utils/models"
)

type UserUseCase interface {
	UsersSignUp(user models.UserSignUp) (domain.TokenUser, error)
	UsersLogin(user models.UserLogin) (domain.TokenUser, error)
}
