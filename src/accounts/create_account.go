package accounts

import "github.com/jeanmolossi/special-broccoli/common/router"

func CreateAccount() router.Handler {
	return func(r *router.Request) *router.Response {
		return router.NewResponse(
			router.WithStatusCode(200),
		)
	}
}
