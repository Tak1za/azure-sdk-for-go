package customerinsights

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// ImagesClient is the the Azure Customer Insights management API provides a RESTful set of web services that interact
// with Azure Customer Insights service to manage your resources. The API has entities that capture the relationship
// between an end user and the Azure Customer Insights service.
type ImagesClient struct {
	BaseClient
}

// NewImagesClient creates an instance of the ImagesClient client.
func NewImagesClient(subscriptionID string) ImagesClient {
	return NewImagesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewImagesClientWithBaseURI creates an instance of the ImagesClient client using a custom endpoint.  Use this when
// interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewImagesClientWithBaseURI(baseURI string, subscriptionID string) ImagesClient {
	return ImagesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// GetUploadURLForData gets data image upload URL.
// Parameters:
// resourceGroupName - the name of the resource group.
// hubName - the name of the hub.
// parameters - parameters supplied to the GetUploadUrlForData operation.
func (client ImagesClient) GetUploadURLForData(ctx context.Context, resourceGroupName string, hubName string, parameters GetImageUploadURLInput) (result ImageDefinition, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ImagesClient.GetUploadURLForData")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetUploadURLForDataPreparer(ctx, resourceGroupName, hubName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "customerinsights.ImagesClient", "GetUploadURLForData", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetUploadURLForDataSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "customerinsights.ImagesClient", "GetUploadURLForData", resp, "Failure sending request")
		return
	}

	result, err = client.GetUploadURLForDataResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "customerinsights.ImagesClient", "GetUploadURLForData", resp, "Failure responding to request")
	}

	return
}

// GetUploadURLForDataPreparer prepares the GetUploadURLForData request.
func (client ImagesClient) GetUploadURLForDataPreparer(ctx context.Context, resourceGroupName string, hubName string, parameters GetImageUploadURLInput) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"hubName":           autorest.Encode("path", hubName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-01-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/images/getDataImageUploadUrl", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetUploadURLForDataSender sends the GetUploadURLForData request. The method will close the
// http.Response Body if it receives an error.
func (client ImagesClient) GetUploadURLForDataSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetUploadURLForDataResponder handles the response to the GetUploadURLForData request. The method always
// closes the http.Response Body.
func (client ImagesClient) GetUploadURLForDataResponder(resp *http.Response) (result ImageDefinition, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetUploadURLForEntityType gets entity type (profile or interaction) image upload URL.
// Parameters:
// resourceGroupName - the name of the resource group.
// hubName - the name of the hub.
// parameters - parameters supplied to the GetUploadUrlForEntityType operation.
func (client ImagesClient) GetUploadURLForEntityType(ctx context.Context, resourceGroupName string, hubName string, parameters GetImageUploadURLInput) (result ImageDefinition, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ImagesClient.GetUploadURLForEntityType")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetUploadURLForEntityTypePreparer(ctx, resourceGroupName, hubName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "customerinsights.ImagesClient", "GetUploadURLForEntityType", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetUploadURLForEntityTypeSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "customerinsights.ImagesClient", "GetUploadURLForEntityType", resp, "Failure sending request")
		return
	}

	result, err = client.GetUploadURLForEntityTypeResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "customerinsights.ImagesClient", "GetUploadURLForEntityType", resp, "Failure responding to request")
	}

	return
}

// GetUploadURLForEntityTypePreparer prepares the GetUploadURLForEntityType request.
func (client ImagesClient) GetUploadURLForEntityTypePreparer(ctx context.Context, resourceGroupName string, hubName string, parameters GetImageUploadURLInput) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"hubName":           autorest.Encode("path", hubName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-01-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/images/getEntityTypeImageUploadUrl", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetUploadURLForEntityTypeSender sends the GetUploadURLForEntityType request. The method will close the
// http.Response Body if it receives an error.
func (client ImagesClient) GetUploadURLForEntityTypeSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetUploadURLForEntityTypeResponder handles the response to the GetUploadURLForEntityType request. The method always
// closes the http.Response Body.
func (client ImagesClient) GetUploadURLForEntityTypeResponder(resp *http.Response) (result ImageDefinition, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
