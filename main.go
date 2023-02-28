package main

import (
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/edsonjuniordev/go-serverless/infra"
)

type Response events.APIGatewayProxyResponse

func Handler(request events.APIGatewayProxyRequest) (Response, error) {
	sqlClient := infra.NewSqlClient()
	defer sqlClient.Close()
	return Response{
		StatusCode: http.StatusOK,
		Body:       "{hello: world}",
	}, nil
}

func main() {
	lambda.Start(Handler)
}
