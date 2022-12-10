package accounts

import (
	"github.com/jeanmolossi/special-broccoli/common/router"
)

func GetAccount() router.Handler {
	return func(r *router.Request) *router.Response {
		return router.NewResponse(
			router.WithStatusCode(200),
		)
	}
}

func PostAccount() router.Handler {
	return func(r *router.Request) *router.Response {
		return router.NewResponse(
			router.WithMessage("POST Here"),
		)
	}
}
