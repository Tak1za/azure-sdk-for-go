// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azmonitor

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
)

// Operations contains the methods for the Operations group.
type Operations interface {
	// List - Lists all of the available operations from Microsoft.Insights provider.
	List(ctx context.Context) (*OperationListResultResponse, error)
}

// operations implements the Operations interface.
type operations struct {
	*Client
}

// List - Lists all of the available operations from Microsoft.Insights provider.
func (client *operations) List(ctx context.Context) (*OperationListResultResponse, error) {
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
	urlPath := "/providers/microsoft.insights/operations"
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("api-version", "2015-04-01")
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodGet, *u)
	return req, nil
}

// listHandleResponse handles the List response.
func (client *operations) listHandleResponse(resp *azcore.Response) (*OperationListResultResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.listHandleError(resp)
	}
	result := OperationListResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.OperationListResult)
}

// listHandleError handles the List error response.
func (client *operations) listHandleError(resp *azcore.Response) error {
	return errors.New(resp.Status)
}