// Package client
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// Copyright (c) 2024 Qwilt Inc.

package client

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Qwilt/terraform-provider-qwilt/qwilt/cdn/api"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

const CertificateSigningRequestsRoot = "/api/v2/certificate-signing-requests"

type CertificateSigningRequestClient struct {
	*Client
	apiEndpoint string
}

func NewCertificateSigningRequestClient(client *Client) *CertificateSigningRequestClient {
	c := CertificateSigningRequestClient{
		Client:      client,
		apiEndpoint: client.endpointBuilder.Build("cert-manager"),
	}
	return &c
}

// GetCertificateSigningRequest - Returns Certificate Signing Request details
func (c *CertificateSigningRequestClient) GetCertificateSigningRequest(id types.Int64) (*api.CertificateSigningRequest, error) {
	if id.IsNull() {
		return nil, fmt.Errorf("csr id is empty")
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s/%s", c.apiEndpoint, CertificateSigningRequestsRoot, id), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	certDetail := api.CertificateSigningRequest{}
	err = json.Unmarshal(body, &certDetail)
	if err != nil {
		return nil, err
	}

	return &certDetail, nil
}

type DomainPairs [][]string

type ChallengeDelegationMap struct {
	pairs DomainPairs
}

func (c *CertificateSigningRequestClient) GetChallengeDelegationDomainsListFromCsrId(id int64) (*ChallengeDelegationMap, error) {
	csr, err := c.GetCertificateSigningRequest(types.Int64Value(id))
	if err != nil {
		return nil, err
	}

	challengeDelegationMap := &ChallengeDelegationMap{
		pairs: make(DomainPairs, len(csr.ChallengeDelegationOfDomainsList)),
	}
	for i := range csr.ChallengeDelegationOfDomainsList {
		challengeDelegationMap.pairs[i] = []string{csr.ChallengeDelegationOfDomainsList[i].FromDomain, csr.ChallengeDelegationOfDomainsList[i].ToDomain}
	}

	return challengeDelegationMap, nil
}