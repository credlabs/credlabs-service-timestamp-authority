package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
)

type CredlabsTimestampRequest struct {
	PublicKeyUri string `json:"publicKeyUri"`
	PrivateKeyUri string `json:"privateKeyUri"`
	CertificateChainUri string `json:"certificateChainUri"`
	Message string `json:"message"`
}

type CredlabsTimestampResponse struct {
	Status string `json:"status"`
	Result string `json:"result"`
}

func HandleRequest(ctx context.Context, event *CredlabsTimestampRequest) (*CredlabsTimestampResponse, error) {
	if event == nil { return nil, fmt.Errorf("Missing Credlabs Timestamp Request") }
	if event.PublicKeyUri == "" { return nil, fmt.Errorf("Missing Public Key Reference") }
	if event.PrivateKeyUri == "" { return nil, fmt.Errorf("Missing Private Key Reference") }
	if event.CertificateChainUri == "" { return nil, fmt.Errorf("Missing Certificate Chain Reference") }
	if event.Message == "" { return nil, fmt.Errorf("Missing Timestamp Message") }

	message := fmt.Sprintf("Hello %s!", event.PublicKeyUri)

	response := CredlabsTimestampResponse{
		Status: message,
	}

	return &response, nil
}

func main() {
	lambda.Start(HandleRequest)
}