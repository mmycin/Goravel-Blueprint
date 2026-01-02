package providers

import (
	"github.com/goravel/framework/contracts/foundation"
	"github.com/goravel/framework/facades"
)

type AppServiceProvider struct {
}

func NewAppServiceProvider() *AppServiceProvider {
	return &AppServiceProvider{}
}

func (receiver *AppServiceProvider) Register(app foundation.Application) {
	// Add your register logic here
}

func (receiver *AppServiceProvider) Boot(app foundation.Application) {
	// Add your boot logic here
	facades.Route().Get("/", func() {
		// Add your route handler here
	})
}

