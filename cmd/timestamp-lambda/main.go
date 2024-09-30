package main

import (
	"fmt"
	"log"
	"encoding/json"
	"github.com/aws/aws-lambda-go/lambda"
)

type CredlabsTimestampRequest struct {
	Name *string `json:"name"`
}

type CredlabsTimestampResponse struct {
	Status *string `json:"status"`
}

func HandleRequest(request *CredlabsTimestampRequest) (*CredlabsTimestampResponse, error) {
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

	return &response, nil
}

func main() {
	log.Println("Version: 1")
	lambda.Start(HandleRequest)
}