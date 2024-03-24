package interfaces

import (
	"github.com/AnanyaDevGo/pkg/domain"
	"github.com/AnanyaDevGo/pkg/utils/models"
)

type OrderUseCase interface {
	OrderItemsFromCart(orderFromCart models.OrderFromCart, userID int) (domain.OrderSuccessResponse, error)
	GetOrderDetails(userId int, page int, count int) ([]models.FullOrderDetails, error)
}
