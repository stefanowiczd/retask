package rest

//go:generate  mockgen -destination=./mock/service_mock.go -package=mock -source=./service_interface.go

type ServicePackageManger interface {
	CalculateOptimumPackagesNumber(
		smallPackageSize int,
		mediumPackageSize int,
		largePackageSize int,
	) (
		smallPackageAmount int,
		mediumPackageAmount int,
		largePackageAmount int,
		err error)
}
