package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
)

type CredlabsTimestampRequest struct {
	Name string `json:"name"`
}

type CredlabsTimestampResponse struct {
	Status string `json:"status"`
}

func HandleRequest(ctx context.Context, event *CredlabsTimestampRequest) (*CredlabsTimestampResponse, error) {
	if event == nil {
		return nil, fmt.Errorf("received nil event")
	}
	message := fmt.Sprintf("Hello %s!", event.Name)

	response := CredlabsTimestampResponse{
		Status: message,
	}

	return &response, nil
}

func main() {
	lambda.Start(HandleRequest)
}