package router

import "fmt"

func DefaultHandler() Handler {
	return func(r *Request) *Response {
		return NewResponse(
			WithStatusCode(404),
			WithMessage(fmt.Sprintf("%s %s %s", r.Method, "Cannot perform request in ", r.Path)),
		)
	}
}
