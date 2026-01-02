package routes

import (
	"github.com/goravel/framework/contracts/route"
	"github.com/goravel/framework/facades"
)

func Web() {
	facades.Route().Get("/", func(ctx route.Context) route.Response {
		return ctx.Response().Json(200, route.Json{
			"Hello": "Goravel",
		})
	})
}

