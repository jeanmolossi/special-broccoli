package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/jeanmolossi/special-broccoli/accounts"
	"github.com/jeanmolossi/special-broccoli/common/router"
)

func main() {
	lambda.Start(Handler)
}

func Handler(request events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	handler := router.NewRouter(router.WithBasePath("/")).Handle()

	switch request.RequestContext.HTTP.Path {
	case "/accounts":
		handler = accounts.GetRouter().Handle()
	}

	return handler(request)
}
