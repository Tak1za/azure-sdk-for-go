// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azcompute

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
)

// Operations contains the methods for the Operations group.
type Operations interface {
	// List - Gets a list of compute operations.
	List(ctx context.Context) (*ComputeOperationListResultResponse, error)
}

// operations implements the Operations interface.
type operations struct {
	*Client
}

// List - Gets a list of compute operations.
func (client *operations) List(ctx context.Context) (*ComputeOperationListResultResponse, error) {
	req, err := client.listCreateRequest()
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.listHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// listCreateRequest creates the List request.
func (client *operations) listCreateRequest() (*azcore.Request, error) {
	urlPath := "/providers/Microsoft.Compute/operations"
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("api-version", "2019-12-01")
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodGet, *u)
	return req, nil
}

// listHandleResponse handles the List response.
func (client *operations) listHandleResponse(resp *azcore.Response) (*ComputeOperationListResultResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.listHandleError(resp)
	}
	result := ComputeOperationListResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.ComputeOperationListResult)
}

// listHandleError handles the List error response.
func (client *operations) listHandleError(resp *azcore.Response) error {
	return errors.New(resp.Status)
}