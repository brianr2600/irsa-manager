package oidc

import (
	"encoding/json"
	"fmt"

	"github.com/kkb0318/irsa-manager/internal/selfhosted"
)

type IdPDiscoveryContents struct {
	jwk            *selfhosted.JWK
	issuerHostPath string
	jwksFileName   string
}

type oidcDiscoveryConfiguration struct {
	Issuer                           string   `json:"issuer"`
	JWKSURI                          string   `json:"jwks_uri"`
	AuthorizationEndpoint            string   `json:"authorization_endpoint"`
	ResponseTypesSupported           []string `json:"response_types_supported"`
	SubjectTypesSupported            []string `json:"subject_types_supported"`
	IDTokenSigningAlgValuesSupported []string `json:"id_token_signing_alg_values_supported"`
	ClaimsSupported                  []string `json:"claims_supported"`
}

func (p *IdPDiscoveryContents) Discovery() ([]byte, error) {
	oidcConfig := oidcDiscoveryConfiguration{
		Issuer:                           fmt.Sprintf("https://%s/", p.issuerHostPath),
		JWKSURI:                          fmt.Sprintf("https://%s/%s", p.issuerHostPath, p.jwksFileName),
		AuthorizationEndpoint:            "urn:kubernetes:programmatic_authorization",
		ResponseTypesSupported:           []string{"id_token"},
		SubjectTypesSupported:            []string{"public"},
		IDTokenSigningAlgValuesSupported: []string{"RS256"},
		ClaimsSupported:                  []string{"sub", "iss"},
	}
	jsonData, err := json.MarshalIndent(oidcConfig, "", "  ")
	if err != nil {
		return nil, err
	}
	return jsonData, nil
}

func (p *IdPDiscoveryContents) JWK() ([]byte, error) {
	jsonData, err := json.MarshalIndent(p.jwk.GetKeys(), "", "  ")
	if err != nil {
		return nil, err
	}
	return jsonData, nil
}

func (p *IdPDiscoveryContents) JWKsFileName() string {
	return p.jwksFileName
}

func NewIdPDiscoveryContents(jwk *selfhosted.JWK, issuerHostPath, jwksFileName string) *IdPDiscoveryContents {
	return &IdPDiscoveryContents{jwk, issuerHostPath, jwksFileName}
}
