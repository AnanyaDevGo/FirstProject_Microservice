package interfaces

import (
	"github.com/AnanyaDevGo/GO-GRPC-PRODUCT-SVC/pkg/domain"
	models "github.com/AnanyaDevGo/GO-GRPC-PRODUCT-SVC/pkg/utils"
)

type ProductRepository interface {
	ShowAllProducts(page int, count int) ([]models.ProductBrief, error)
	AddProducts(product models.Product) (domain.Product, error)
	ProductAlreadyExist(Name string) bool
	StockInvalid(Name string) bool
	DeleteProducts(id int) error
	CheckProduct(pid int) (bool, error)
	UpdateProduct(pid int, stock int) (models.ProductUpdateReciever, error)
	GetQuantityFromProductID(id int) (int, error)
	GetPriceOfProductFromID(prodcut_id int) (float64, error)
	ProductStockMinus(productID, stock int) error
}
