package application

type ServicePackageManager struct{}

func NewServicePackageManager() *ServicePackageManager { return &ServicePackageManager{} }

func (s *ServicePackageManager) CalculateOptimumPackagesNumber(
	smallPackageSize, mediumPackageSize, largePackageSize int,
) (int, int, int, error) {
	smallPackageAmount, mediumPackageAmount, largePackageAmount := -1, -1, -1

	return smallPackageAmount, mediumPackageAmount, largePackageAmount, nil
}
