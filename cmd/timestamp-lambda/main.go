package main

import (
	"log"
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(HandleRequest)
}

func HandleRequest(ctx context.Context) (events.APIGatewayProxyResponse, error) {
	log.Println("Hello World!")

	response := events.APIGatewayProxyResponse{
			StatusCode: 200,
	}

	return response, nil
}
