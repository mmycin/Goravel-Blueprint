package bootstrap

import (
	"github.com/goravel/framework/facades"

	"github.com/mmycin/goravel-test/app/providers"
)

func Boot() {
	// We boot the app providers defined in config/app.go
	app := facades.App()

	//Boot providers.
	app.BootWith(providers.NewAppServiceProvider())
}

