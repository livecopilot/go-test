//go:build wireinject
// +build wireinject

package main

import (
	"go-test/study_wire/service"

	"github.com/google/wire"
)

func InitializeBar() *service.Bar {
	wire.Build(service.ProviderSet)
	return &service.Bar{}
}
