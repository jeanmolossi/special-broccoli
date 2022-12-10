package router

import (
	"fmt"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

type (
	ResponseOption func(*Response)

	Response struct {
		StatusCode        int               `json:"statusCode"`
		Headers           map[string]string `json:"headers"`
		MultiValueHeaders map[string]string `json:"multiValueHeaders"`
		Body              *Body             `json:"body"`
		IsBase64Encoded   bool              `json"isBase64Encoded,omitempty"`
		Cookies           []string          `json:"cookies"`
	}
)

func (r *Response) SetStatusCode(statusCode int) *Response {
	r.StatusCode = statusCode
	r.Body.StatusCode = statusCode
	r.Body.Message = http.StatusText(statusCode)

	return r
}

func (r *Response) SetData(data interface{}) *Response {
	r.Body.Data = data
	r.Body.Metadata = nil

	return r
}

func (r *Response) SetMetadata(metadata interface{}) *Response {
	r.Body.Data = nil
	r.Body.Metadata = metadata

	return r
}

func (r *Response) AddHeader(key, value string) *Response {
	r.Headers[key] = value

	return r
}

func ResponseError(request *Request, err error) (events.APIGatewayV2HTTPResponse, error) {
	if request == nil {
		return WrapResponse(nil)
	}

	cookies := make([]string, 0)
	for key, value := range request.Cookies {
		cookies = append(cookies, fmt.Sprintf("%s=%s", key, value))
	}

	response := NewResponse()
	response.Headers = request.Headers
	response.Cookies = cookies
	response.Body = &Body{}
	response.Body.StatusCode = 500
	if err != nil {
		response.Body.Message = err.Error()
	} else {
		response.Body.Message = ""
	}

	return WrapResponse(response)
}

func WrapResponse(response *Response) (events.APIGatewayV2HTTPResponse, error) {
	if response == nil {
		return events.APIGatewayV2HTTPResponse{
			StatusCode:      500,
			Headers:         make(map[string]string),
			Cookies:         make([]string, 0),
			IsBase64Encoded: false,
			Body: (&Body{
				StatusCode: 500,
				Message:    "Nil response received",
			}).ToJson(),
		}, nil
	}

	return events.APIGatewayV2HTTPResponse{
		StatusCode:      response.StatusCode,
		Headers:         response.Headers,
		Cookies:         response.Cookies,
		IsBase64Encoded: false,
		Body:            response.Body.ToJson(),
	}, nil
}

func NewResponse(options ...ResponseOption) *Response {
	defaultResponse := &Response{
		StatusCode:        200,
		Headers:           make(map[string]string),
		MultiValueHeaders: make(map[string]string),
		Body: &Body{
			StatusCode: 200,
			Message:    "OK",
		},
		IsBase64Encoded: false,
		Cookies:         make([]string, 0),
	}

	if len(options) > 0 {
		for _, option := range options {
			option(defaultResponse)
		}
	}

	return defaultResponse
}

func WithStatusCode(statusCode int) ResponseOption {
	return func(r *Response) {
		r.SetStatusCode(statusCode)
	}
}

func WithMessage(message string) ResponseOption {
	return func(r *Response) {
		if r.Body != nil {
			r.Body.Message = message
		} else {
			r.Body = &Body{
				Message: message,
			}
		}
	}
}

func WithData(data interface{}) ResponseOption {
	return func(r *Response) {
		r.SetData(data)
	}
}
