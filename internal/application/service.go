package application

import "context"

type ServicePackageManager struct{}

func NewServicePackageManager() *ServicePackageManager { return &ServicePackageManager{} }

func (s *ServicePackageManager) CalculateOptimumPackagesNumber(
	ctx context.Context,
	smallPackageSize, mediumPackageSize, largePackageSize int,
) (int, int, int, error) {
	smallPackageAmount, mediumPackageAmount, largePackageAmount := -1, -1, -1

	return smallPackageAmount, mediumPackageAmount, largePackageAmount, nil
}
