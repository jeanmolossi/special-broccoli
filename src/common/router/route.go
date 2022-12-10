package router

type (
	Route struct {
		method  Method
		path    string
		handler Handler
	}
)
