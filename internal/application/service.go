package application

import "context"

type ServicePacksManager struct{}

func NewServicePacksManager() *ServicePacksManager { return &ServicePacksManager{} }

func (s *ServicePacksManager) CalculateOptimumPacksAmount(
	ctx context.Context,
	smallPackSize, mediumPackSize, largePackSize int,
) (int, int, int, error) {
	smallPacksAmount, mediumPacksAmount, largePacksAmount := -1, -1, -1

	return smallPacksAmount, mediumPacksAmount, largePacksAmount, nil
}
