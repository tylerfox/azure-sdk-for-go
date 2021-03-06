package subscription

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
	"github.com/satori/go.uuid"
	"net/http"
)

// DefinitionsClient is the subscription definitions client provides an interface to create, modify and retrieve azure
// subscriptions programmatically.
type DefinitionsClient struct {
	BaseClient
}

// NewDefinitionsClient creates an instance of the DefinitionsClient client.
func NewDefinitionsClient() DefinitionsClient {
	return NewDefinitionsClientWithBaseURI(DefaultBaseURI)
}

// NewDefinitionsClientWithBaseURI creates an instance of the DefinitionsClient client.
func NewDefinitionsClientWithBaseURI(baseURI string) DefinitionsClient {
	return DefinitionsClient{NewWithBaseURI(baseURI)}
}

// Create create an Azure subscription definition.
// Parameters:
// subscriptionDefinitionName - the name of the Azure subscription definition.
// body - the subscription definition creation.
func (client DefinitionsClient) Create(ctx context.Context, subscriptionDefinitionName string, body Definition) (result DefinitionsCreateFuture, err error) {
	req, err := client.CreatePreparer(ctx, subscriptionDefinitionName, body)
	if err != nil {
		err = autorest.NewErrorWithError(err, "subscription.DefinitionsClient", "Create", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "subscription.DefinitionsClient", "Create", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreatePreparer prepares the Create request.
func (client DefinitionsClient) CreatePreparer(ctx context.Context, subscriptionDefinitionName string, body Definition) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionDefinitionName": autorest.Encode("path", subscriptionDefinitionName),
	}

	const APIVersion = "2017-11-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Subscription/subscriptionDefinitions/{subscriptionDefinitionName}", pathParameters),
		autorest.WithJSON(body),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client DefinitionsClient) CreateSender(req *http.Request) (future DefinitionsCreateFuture, err error) {
	var resp *http.Response
	resp, err = autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	if err != nil {
		return
	}
	err = autorest.Respond(resp, azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted))
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client DefinitionsClient) CreateResponder(resp *http.Response) (result Definition, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Get get an Azure subscription definition.
// Parameters:
// subscriptionDefinitionName - the name of the Azure subscription definition.
func (client DefinitionsClient) Get(ctx context.Context, subscriptionDefinitionName string) (result Definition, err error) {
	req, err := client.GetPreparer(ctx, subscriptionDefinitionName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "subscription.DefinitionsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "subscription.DefinitionsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "subscription.DefinitionsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client DefinitionsClient) GetPreparer(ctx context.Context, subscriptionDefinitionName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionDefinitionName": autorest.Encode("path", subscriptionDefinitionName),
	}

	const APIVersion = "2017-11-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Subscription/subscriptionDefinitions/{subscriptionDefinitionName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client DefinitionsClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client DefinitionsClient) GetResponder(resp *http.Response) (result Definition, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetOperationStatus retrieves the status of the subscription definition PUT operation. The URI of this API is
// returned in the Location field of the response header.
// Parameters:
// operationID - the operation ID, which can be found from the Location field in the generate recommendation
// response header.
func (client DefinitionsClient) GetOperationStatus(ctx context.Context, operationID uuid.UUID) (result Definition, err error) {
	req, err := client.GetOperationStatusPreparer(ctx, operationID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "subscription.DefinitionsClient", "GetOperationStatus", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetOperationStatusSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "subscription.DefinitionsClient", "GetOperationStatus", resp, "Failure sending request")
		return
	}

	result, err = client.GetOperationStatusResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "subscription.DefinitionsClient", "GetOperationStatus", resp, "Failure responding to request")
	}

	return
}

// GetOperationStatusPreparer prepares the GetOperationStatus request.
func (client DefinitionsClient) GetOperationStatusPreparer(ctx context.Context, operationID uuid.UUID) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"operationId": autorest.Encode("path", operationID),
	}

	const APIVersion = "2017-11-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Subscription/subscriptionOperations/{operationId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetOperationStatusSender sends the GetOperationStatus request. The method will close the
// http.Response Body if it receives an error.
func (client DefinitionsClient) GetOperationStatusSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetOperationStatusResponder handles the response to the GetOperationStatus request. The method always
// closes the http.Response Body.
func (client DefinitionsClient) GetOperationStatusResponder(resp *http.Response) (result Definition, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List list an Azure subscription definition by subscriptionId.
func (client DefinitionsClient) List(ctx context.Context) (result DefinitionListPage, err error) {
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "subscription.DefinitionsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.dl.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "subscription.DefinitionsClient", "List", resp, "Failure sending request")
		return
	}

	result.dl, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "subscription.DefinitionsClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client DefinitionsClient) ListPreparer(ctx context.Context) (*http.Request, error) {
	const APIVersion = "2017-11-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/providers/Microsoft.Subscription/subscriptionDefinitions"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client DefinitionsClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client DefinitionsClient) ListResponder(resp *http.Response) (result DefinitionList, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client DefinitionsClient) listNextResults(lastResults DefinitionList) (result DefinitionList, err error) {
	req, err := lastResults.definitionListPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "subscription.DefinitionsClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "subscription.DefinitionsClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "subscription.DefinitionsClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client DefinitionsClient) ListComplete(ctx context.Context) (result DefinitionListIterator, err error) {
	result.page, err = client.List(ctx)
	return
}
