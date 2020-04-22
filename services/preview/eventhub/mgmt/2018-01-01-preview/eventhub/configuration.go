package eventhub

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

// ConfigurationClient is the azure Event Hubs client for managing Event Hubs Cluster, IPFilter Rules and
// VirtualNetworkRules resources.
type ConfigurationClient struct {
	BaseClient
}

// NewConfigurationClient creates an instance of the ConfigurationClient client.
func NewConfigurationClient(subscriptionID string) ConfigurationClient {
	return NewConfigurationClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewConfigurationClientWithBaseURI creates an instance of the ConfigurationClient client using a custom endpoint.
// Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewConfigurationClientWithBaseURI(baseURI string, subscriptionID string) ConfigurationClient {
	return ConfigurationClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Get get all Event Hubs Cluster settings - a collection of key/value pairs which represent the quotas and settings
// imposed on the cluster.
// Parameters:
// resourceGroupName - name of the resource group within the azure subscription.
// clusterName - the name of the Event Hubs Cluster.
func (client ConfigurationClient) Get(ctx context.Context, resourceGroupName string, clusterName string) (result ClusterQuotaConfigurationProperties, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ConfigurationClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: clusterName,
			Constraints: []validation.Constraint{{Target: "clusterName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "clusterName", Name: validation.MinLength, Rule: 6, Chain: nil}}}}); err != nil {
		return result, validation.NewError("eventhub.ConfigurationClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, clusterName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "eventhub.ConfigurationClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "eventhub.ConfigurationClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "eventhub.ConfigurationClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client ConfigurationClient) GetPreparer(ctx context.Context, resourceGroupName string, clusterName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"clusterName":       autorest.Encode("path", clusterName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-01-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/clusters/{clusterName}/quotaConfiguration/default", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ConfigurationClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ConfigurationClient) GetResponder(resp *http.Response) (result ClusterQuotaConfigurationProperties, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Patch replace all specified Event Hubs Cluster settings with those contained in the request body. Leaves the
// settings not specified in the request body unmodified.
// Parameters:
// resourceGroupName - name of the resource group within the azure subscription.
// clusterName - the name of the Event Hubs Cluster.
// parameters - parameters for creating an Event Hubs Cluster resource.
func (client ConfigurationClient) Patch(ctx context.Context, resourceGroupName string, clusterName string, parameters ClusterQuotaConfigurationProperties) (result ClusterQuotaConfigurationProperties, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ConfigurationClient.Patch")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: clusterName,
			Constraints: []validation.Constraint{{Target: "clusterName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "clusterName", Name: validation.MinLength, Rule: 6, Chain: nil}}}}); err != nil {
		return result, validation.NewError("eventhub.ConfigurationClient", "Patch", err.Error())
	}

	req, err := client.PatchPreparer(ctx, resourceGroupName, clusterName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "eventhub.ConfigurationClient", "Patch", nil, "Failure preparing request")
		return
	}

	resp, err := client.PatchSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "eventhub.ConfigurationClient", "Patch", resp, "Failure sending request")
		return
	}

	result, err = client.PatchResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "eventhub.ConfigurationClient", "Patch", resp, "Failure responding to request")
	}

	return
}

// PatchPreparer prepares the Patch request.
func (client ConfigurationClient) PatchPreparer(ctx context.Context, resourceGroupName string, clusterName string, parameters ClusterQuotaConfigurationProperties) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"clusterName":       autorest.Encode("path", clusterName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-01-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/clusters/{clusterName}/quotaConfiguration/default", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// PatchSender sends the Patch request. The method will close the
// http.Response Body if it receives an error.
func (client ConfigurationClient) PatchSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// PatchResponder handles the response to the Patch request. The method always
// closes the http.Response Body.
func (client ConfigurationClient) PatchResponder(resp *http.Response) (result ClusterQuotaConfigurationProperties, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
