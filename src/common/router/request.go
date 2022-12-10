package router

import (
	"encoding/base64"
	"strings"

	"github.com/aws/aws-lambda-go/events"
)

type (
	Request struct {
		Method      Method
		Path        string
		Headers     map[string]string
		Cookies     map[string]string
		Body        string
		PathParams  map[string]string
		QueryParams map[string]string
	}
)

func NewRequestFromApiGatewayEvent(event events.APIGatewayV2HTTPRequest) (*Request, error) {
	method := ToMethod(event.RequestContext.HTTP.Method)
	path := event.RequestContext.HTTP.Path

	request := &Request{
		Method:      method,
		Path:        path,
		Headers:     event.Headers,
		Body:        event.Body,
		PathParams:  event.PathParameters,
		QueryParams: event.QueryStringParameters,
	}

	cookies := map[string]string{}

	for _, cookie := range event.Cookies {
		cookieSlices := strings.Split(cookie, "=")
		cookies[cookieSlices[0]] = cookieSlices[1]
	}

	request.Cookies = cookies

	content := event.Headers["Content-Type"]
	if content == "" {
		content = event.Headers["content-type"]
	}

	if content != "application/json" {
		decodedBody, err := base64.StdEncoding.DecodeString(event.Body)
		if err != nil {
			return request, err
		} else {
			event.Body = string(decodedBody)
		}
	}

	request.Headers = event.Headers
	request.Body = event.Body

	return request, nil
}
