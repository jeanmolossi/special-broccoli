package router

import "strings"

type (
	Method string
)

const (
	GET     Method = "GET"
	HEAD    Method = "HEAD"
	OPTIONS Method = "OPTIONS"
	POST    Method = "POST"
	PUT     Method = "PUT"
	PATCH   Method = "PATCH"
	DELETE  Method = "DELETE"
	ANY     Method = "ANY"
)

func ToMethod(path string) Method {
	switch strings.ToUpper(path) {
	case "GET":
		return GET
	case "HEAD":
		return HEAD
	case "OPTIONS":
		return OPTIONS
	case "POST":
		return POST
	case "PUT":
		return PUT
	case "PATCH":
		return PATCH
	case "DELETE":
		return DELETE
	default:
		return ANY
	}
}
