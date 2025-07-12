package usecase

import (
	"github.com/devfullcycle/20-CleanArch/internal/entity"
)

type (
	ListOrderUseCase struct {
		OrderRepository entity.OrderRepositoryInterface
	}

	ListOrderOutputDTO struct {
		Orders []OrderOutputDTO `json:"orders"`
	}
)

func NewListOrderUseCase(
	OrderRepository entity.OrderRepositoryInterface,
) *ListOrderUseCase {
	return &ListOrderUseCase{
		OrderRepository: OrderRepository,
	}
}

func (c *ListOrderUseCase) Execute() (ListOrderOutputDTO, error) {
	orders, err := c.OrderRepository.List()
	if err != nil {
		return ListOrderOutputDTO{}, err
	}

	var listOrders []OrderOutputDTO
	for _, order := range orders {
		listOrders = append(listOrders, OrderOutputDTO{
			ID:         order.ID,
			Price:      order.Price,
			Tax:        order.Tax,
			FinalPrice: order.FinalPrice,
		})
	}

	return ListOrderOutputDTO{Orders: listOrders}, nil
}
