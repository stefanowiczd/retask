package rest

import "context"

//go:generate  mockgen -destination=./mock/service_mock.go -package=mock -source=./service_interface.go

type ServicePacksManger interface {
	CalculateOptimumPacksAmount(
		ctx context.Context,
		smallPackSize int,
		mediumPackSize int,
		largePackSize int,
	) (
		smallPacksAmount int,
		mediumPacksAmount int,
		largePacksAmount int,
		err error)
}
