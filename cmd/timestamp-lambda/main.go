package main

import (
	"fmt"
	"log"
	"context"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(HandleRequest)
}

func HandleRequest(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Println("Hello World!")

	json_msg, err := json.Marshal(event)

	if err == nil {
			panic(err)
	}//if

	str := string(json_msg)
	fmt.Println(str)

	response := events.APIGatewayProxyResponse{
			StatusCode: 200,
	}

	return response, nil
}
