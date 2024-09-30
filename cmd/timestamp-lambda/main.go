package main

import (
	"fmt"
	"log"
	"context"
	"encoding/json"
	"github.com/aws/aws-lambda-go/lambda"
)

type CredlabsTimestampRequest struct {
	Name *string `json:"name"`
}

type CredlabsTimestampResponse struct {
	Status *string `json:"status"`
}

func main() {
	lambda.Start(HandleRequest)
}

func HandleRequest(ctx context.Context, request CredlabsTimestampRequest) (CredlabsTimestampResponse, error) {
	log.Println("Hello World!")

	json_msg, err := json.Marshal(request)

	if err == nil {
			panic(err)
	}//if

	str := string(json_msg)
	fmt.Println(str)

	status := "success"

	response := CredlabsTimestampResponse{
			Status: &status,
	}

	return response, nil
}
