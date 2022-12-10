package router

import (
	"github.com/aws/aws-lambda-go/events"
)

type (
	Router struct {
		routes   map[string]map[Method]*Route
		basePath string
	}

	Handler func(*Request) *Response

	RouteConfig func(r *Router)
)

func (r *Router) Handle() func(event events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	return func(event events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
		request, err := NewRequestFromApiGatewayEvent(event)
		if err != nil {
			return ResponseError(request, err)
		}

		method := ToMethod(event.RequestContext.HTTP.Method)
		path := handlePath(event.RequestContext.HTTP.Path)

		routes, hasRoutes := r.routes[path]
		if !hasRoutes {
			// do default
			handler := DefaultHandler()
			return WrapResponse(handler(request))
		}

		route, hasRoute := routes[method]
		if !hasRoute {
			// do default
			handler := DefaultHandler()
			return WrapResponse(handler(request))
		}

		return WrapResponse(route.handler(request))
	}
}

func NewRouter(routes ...RouteConfig) *Router {
	r := &Router{
		basePath: "",
		routes:   make(map[string]map[Method]*Route),
	}

	if len(routes) == 0 {
		panic("you should provide at least 1 route handler")
	}

	for _, config := range routes {
		config(r)
	}

	return r
}

func AddRoute(methodStr, path string, fn Handler) RouteConfig {
	path = handleBackslashAtStart(path, true)
	method := ToMethod(methodStr)

	return func(r *Router) {
		path = handleBackslashAtEnd(r.basePath+path, false)
		_, hasRouteInPath := r.routes[path]

		if !hasRouteInPath {
			r.routes[path] = make(map[Method]*Route)
		}

		r.routes[path][method] = &Route{
			method:  method,
			path:    path,
			handler: fn,
		}
	}
}

func WithBasePath(basePath string) RouteConfig {
	return func(r *Router) {
		r.basePath = handleBackslashAtEnd(handleBackslashAtStart(basePath, true), false)
	}
}
