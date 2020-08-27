package network

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
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// VirtualHubIPConfigurationClient is the network Client
type VirtualHubIPConfigurationClient struct {
	BaseClient
}

// NewVirtualHubIPConfigurationClient creates an instance of the VirtualHubIPConfigurationClient client.
func NewVirtualHubIPConfigurationClient(subscriptionID string) VirtualHubIPConfigurationClient {
	return NewVirtualHubIPConfigurationClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewVirtualHubIPConfigurationClientWithBaseURI creates an instance of the VirtualHubIPConfigurationClient client
// using a custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign
// clouds, Azure stack).
func NewVirtualHubIPConfigurationClientWithBaseURI(baseURI string, subscriptionID string) VirtualHubIPConfigurationClient {
	return VirtualHubIPConfigurationClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates a VirtualHubIpConfiguration resource if it doesn't exist else updates the existing
// VirtualHubIpConfiguration.
// Parameters:
// resourceGroupName - the resource group name of the VirtualHub.
// virtualHubName - the name of the VirtualHub.
// IPConfigName - the name of the ipconfig.
// parameters - hub Ip Configuration parameters.
func (client VirtualHubIPConfigurationClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, virtualHubName string, IPConfigName string, parameters HubIPConfiguration) (result VirtualHubIPConfigurationCreateOrUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VirtualHubIPConfigurationClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.HubIPConfigurationPropertiesFormat", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "parameters.HubIPConfigurationPropertiesFormat.PublicIPAddress", Name: validation.Null, Rule: false,
					Chain: []validation.Constraint{{Target: "parameters.HubIPConfigurationPropertiesFormat.PublicIPAddress.ExtendedLocation", Name: validation.Null, Rule: false,
						Chain: []validation.Constraint{{Target: "parameters.HubIPConfigurationPropertiesFormat.PublicIPAddress.ExtendedLocation.Name", Name: validation.Null, Rule: true, Chain: nil},
							{Target: "parameters.HubIPConfigurationPropertiesFormat.PublicIPAddress.ExtendedLocation.Type", Name: validation.Null, Rule: true, Chain: nil},
						}},
						{Target: "parameters.HubIPConfigurationPropertiesFormat.PublicIPAddress.PublicIPAddressPropertiesFormat", Name: validation.Null, Rule: false,
							Chain: []validation.Constraint{{Target: "parameters.HubIPConfigurationPropertiesFormat.PublicIPAddress.PublicIPAddressPropertiesFormat.IPConfiguration", Name: validation.Null, Rule: false,
								Chain: []validation.Constraint{{Target: "parameters.HubIPConfigurationPropertiesFormat.PublicIPAddress.PublicIPAddressPropertiesFormat.IPConfiguration.IPConfigurationPropertiesFormat", Name: validation.Null, Rule: false,
									Chain: []validation.Constraint{{Target: "parameters.HubIPConfigurationPropertiesFormat.PublicIPAddress.PublicIPAddressPropertiesFormat.IPConfiguration.IPConfigurationPropertiesFormat.PublicIPAddress", Name: validation.Null, Rule: false, Chain: nil}}},
								}},
							}},
					}},
				}}}}}); err != nil {
		return result, validation.NewError("network.VirtualHubIPConfigurationClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, virtualHubName, IPConfigName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VirtualHubIPConfigurationClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VirtualHubIPConfigurationClient", "CreateOrUpdate", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client VirtualHubIPConfigurationClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, virtualHubName string, IPConfigName string, parameters HubIPConfiguration) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"ipConfigName":      autorest.Encode("path", IPConfigName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"virtualHubName":    autorest.Encode("path", virtualHubName),
	}

	const APIVersion = "2020-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	parameters.Etag = nil
	parameters.Type = nil
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualHubs/{virtualHubName}/ipConfigurations/{ipConfigName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualHubIPConfigurationClient) CreateOrUpdateSender(req *http.Request) (future VirtualHubIPConfigurationCreateOrUpdateFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client VirtualHubIPConfigurationClient) CreateOrUpdateResponder(resp *http.Response) (result HubIPConfiguration, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes a VirtualHubIpConfiguration.
// Parameters:
// resourceGroupName - the resource group name of the VirtualHubBgpConnection.
// virtualHubName - the name of the VirtualHub.
// IPConfigName - the name of the ipconfig.
func (client VirtualHubIPConfigurationClient) Delete(ctx context.Context, resourceGroupName string, virtualHubName string, IPConfigName string) (result VirtualHubIPConfigurationDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VirtualHubIPConfigurationClient.Delete")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceGroupName, virtualHubName, IPConfigName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VirtualHubIPConfigurationClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VirtualHubIPConfigurationClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client VirtualHubIPConfigurationClient) DeletePreparer(ctx context.Context, resourceGroupName string, virtualHubName string, IPConfigName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"ipConfigName":      autorest.Encode("path", IPConfigName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"virtualHubName":    autorest.Encode("path", virtualHubName),
	}

	const APIVersion = "2020-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualHubs/{virtualHubName}/ipConfigurations/{ipConfigName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualHubIPConfigurationClient) DeleteSender(req *http.Request) (future VirtualHubIPConfigurationDeleteFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client VirtualHubIPConfigurationClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get retrieves the details of a Virtual Hub Ip configuration.
// Parameters:
// resourceGroupName - the resource group name of the VirtualHub.
// virtualHubName - the name of the VirtualHub.
// IPConfigName - the name of the ipconfig.
func (client VirtualHubIPConfigurationClient) Get(ctx context.Context, resourceGroupName string, virtualHubName string, IPConfigName string) (result HubIPConfiguration, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VirtualHubIPConfigurationClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, virtualHubName, IPConfigName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VirtualHubIPConfigurationClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.VirtualHubIPConfigurationClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VirtualHubIPConfigurationClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client VirtualHubIPConfigurationClient) GetPreparer(ctx context.Context, resourceGroupName string, virtualHubName string, IPConfigName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"ipConfigName":      autorest.Encode("path", IPConfigName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"virtualHubName":    autorest.Encode("path", virtualHubName),
	}

	const APIVersion = "2020-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualHubs/{virtualHubName}/ipConfigurations/{ipConfigName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualHubIPConfigurationClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client VirtualHubIPConfigurationClient) GetResponder(resp *http.Response) (result HubIPConfiguration, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List retrieves the details of all VirtualHubIpConfigurations.
// Parameters:
// resourceGroupName - the resource group name of the VirtualHub.
// virtualHubName - the name of the VirtualHub.
func (client VirtualHubIPConfigurationClient) List(ctx context.Context, resourceGroupName string, virtualHubName string) (result ListVirtualHubIPConfigurationResultsPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VirtualHubIPConfigurationClient.List")
		defer func() {
			sc := -1
			if result.lvhicr.Response.Response != nil {
				sc = result.lvhicr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, resourceGroupName, virtualHubName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VirtualHubIPConfigurationClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.lvhicr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.VirtualHubIPConfigurationClient", "List", resp, "Failure sending request")
		return
	}

	result.lvhicr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VirtualHubIPConfigurationClient", "List", resp, "Failure responding to request")
	}
	if result.lvhicr.hasNextLink() && result.lvhicr.IsEmpty() {
		err = result.NextWithContext(ctx)
	}

	return
}

// ListPreparer prepares the List request.
func (client VirtualHubIPConfigurationClient) ListPreparer(ctx context.Context, resourceGroupName string, virtualHubName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"virtualHubName":    autorest.Encode("path", virtualHubName),
	}

	const APIVersion = "2020-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualHubs/{virtualHubName}/ipConfigurations", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualHubIPConfigurationClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client VirtualHubIPConfigurationClient) ListResponder(resp *http.Response) (result ListVirtualHubIPConfigurationResults, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client VirtualHubIPConfigurationClient) listNextResults(ctx context.Context, lastResults ListVirtualHubIPConfigurationResults) (result ListVirtualHubIPConfigurationResults, err error) {
	req, err := lastResults.listVirtualHubIPConfigurationResultsPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network.VirtualHubIPConfigurationClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network.VirtualHubIPConfigurationClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VirtualHubIPConfigurationClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client VirtualHubIPConfigurationClient) ListComplete(ctx context.Context, resourceGroupName string, virtualHubName string) (result ListVirtualHubIPConfigurationResultsIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VirtualHubIPConfigurationClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, resourceGroupName, virtualHubName)
	return
}
