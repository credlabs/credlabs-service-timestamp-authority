package main

import (
	"fmt"
)

type CredlabsTimestampRequest struct {
	PublicKeyType string `json:"publicKeyType"`
	PublicKey string `json:"publicKey"`
	PrivateKeyType string `json:"privateKeyType"`
	PrivateKey string `json:"privateKey"`
	CertificateChainType string `json:"certificateChainType"`
	CertificateChain string `json:"certificateChain"`
	Request string `json:"request"`
}

type CredlabsTimestampResponse struct {
	Status string `json:"status"`
	Result string `json:"result"`
}

func validateTimestampRequest(event *CredlabsTimestampRequest) (error) {
	if event == nil { return fmt.Errorf("Missing Credlabs Timestamp Request") }
	if event.PublicKeyType == "" { return fmt.Errorf("Missing Public Key Type") }
	if event.PublicKey == "" { return fmt.Errorf("Missing Public Key") }
	if event.PrivateKeyType == "" { return fmt.Errorf("Missing Private Key") }
	if event.PrivateKey == "" { return fmt.Errorf("Missing Private Key Type") }
	if event.CertificateChainType == "" { return fmt.Errorf("Missing Certificate Chain Type") }
	if event.CertificateChain == "" { return fmt.Errorf("Missing Certificate Chain") }
	if event.Request == "" { return fmt.Errorf("Missing Timestamp Request") }
	return nil;
}
