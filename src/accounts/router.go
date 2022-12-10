package accounts

import "github.com/jeanmolossi/special-broccoli/common/router"

func GetRouter() *router.Router {
	return router.NewRouter(
		router.WithBasePath("/accounts"),
		router.AddRoute("POST", "/", CreateAccount()),
	)
}
