package migrate

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

// JobsClient is the discover your workloads for Azure.
type JobsClient struct {
	BaseClient
}

// NewJobsClient creates an instance of the JobsClient client.
func NewJobsClient() JobsClient {
	return NewJobsClientWithBaseURI(DefaultBaseURI)
}

// NewJobsClientWithBaseURI creates an instance of the JobsClient client using a custom endpoint.  Use this when
// interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewJobsClientWithBaseURI(baseURI string) JobsClient {
	return JobsClient{NewWithBaseURI(baseURI)}
}

// GetAllJobsInSite sends the get all jobs in site request.
// Parameters:
// subscriptionID - the ID of the target subscription.
// resourceGroupName - the name of the resource group. The name is case insensitive.
// siteName - site name.
// APIVersion - the API version to use for this operation.
func (client JobsClient) GetAllJobsInSite(ctx context.Context, subscriptionID string, resourceGroupName string, siteName string, APIVersion string) (result VMwareJobCollectionPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/JobsClient.GetAllJobsInSite")
		defer func() {
			sc := -1
			if result.vmjc.Response.Response != nil {
				sc = result.vmjc.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.getAllJobsInSiteNextResults
	req, err := client.GetAllJobsInSitePreparer(ctx, subscriptionID, resourceGroupName, siteName, APIVersion)
	if err != nil {
		err = autorest.NewErrorWithError(err, "migrate.JobsClient", "GetAllJobsInSite", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetAllJobsInSiteSender(req)
	if err != nil {
		result.vmjc.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "migrate.JobsClient", "GetAllJobsInSite", resp, "Failure sending request")
		return
	}

	result.vmjc, err = client.GetAllJobsInSiteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "migrate.JobsClient", "GetAllJobsInSite", resp, "Failure responding to request")
	}

	return
}

// GetAllJobsInSitePreparer prepares the GetAllJobsInSite request.
func (client JobsClient) GetAllJobsInSitePreparer(ctx context.Context, subscriptionID string, resourceGroupName string, siteName string, APIVersion string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"siteName":          autorest.Encode("path", siteName),
		"subscriptionId":    autorest.Encode("path", subscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OffAzure/VMwareSites/{siteName}/jobs", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetAllJobsInSiteSender sends the GetAllJobsInSite request. The method will close the
// http.Response Body if it receives an error.
func (client JobsClient) GetAllJobsInSiteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetAllJobsInSiteResponder handles the response to the GetAllJobsInSite request. The method always
// closes the http.Response Body.
func (client JobsClient) GetAllJobsInSiteResponder(resp *http.Response) (result VMwareJobCollection, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// getAllJobsInSiteNextResults retrieves the next set of results, if any.
func (client JobsClient) getAllJobsInSiteNextResults(ctx context.Context, lastResults VMwareJobCollection) (result VMwareJobCollection, err error) {
	req, err := lastResults.vMwareJobCollectionPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "migrate.JobsClient", "getAllJobsInSiteNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.GetAllJobsInSiteSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "migrate.JobsClient", "getAllJobsInSiteNextResults", resp, "Failure sending next results request")
	}
	result, err = client.GetAllJobsInSiteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "migrate.JobsClient", "getAllJobsInSiteNextResults", resp, "Failure responding to next results request")
	}
	return
}

// GetAllJobsInSiteComplete enumerates all values, automatically crossing page boundaries as required.
func (client JobsClient) GetAllJobsInSiteComplete(ctx context.Context, subscriptionID string, resourceGroupName string, siteName string, APIVersion string) (result VMwareJobCollectionIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/JobsClient.GetAllJobsInSite")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.GetAllJobsInSite(ctx, subscriptionID, resourceGroupName, siteName, APIVersion)
	return
}

// GetJob sends the get job request.
// Parameters:
// subscriptionID - the ID of the target subscription.
// resourceGroupName - the name of the resource group. The name is case insensitive.
// siteName - site name.
// jobName - job ARM name.
// APIVersion - the API version to use for this operation.
func (client JobsClient) GetJob(ctx context.Context, subscriptionID string, resourceGroupName string, siteName string, jobName string, APIVersion string) (result VMwareJob, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/JobsClient.GetJob")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetJobPreparer(ctx, subscriptionID, resourceGroupName, siteName, jobName, APIVersion)
	if err != nil {
		err = autorest.NewErrorWithError(err, "migrate.JobsClient", "GetJob", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetJobSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "migrate.JobsClient", "GetJob", resp, "Failure sending request")
		return
	}

	result, err = client.GetJobResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "migrate.JobsClient", "GetJob", resp, "Failure responding to request")
	}

	return
}

// GetJobPreparer prepares the GetJob request.
func (client JobsClient) GetJobPreparer(ctx context.Context, subscriptionID string, resourceGroupName string, siteName string, jobName string, APIVersion string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"jobName":           autorest.Encode("path", jobName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"siteName":          autorest.Encode("path", siteName),
		"subscriptionId":    autorest.Encode("path", subscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OffAzure/VMwareSites/{siteName}/jobs/{jobName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetJobSender sends the GetJob request. The method will close the
// http.Response Body if it receives an error.
func (client JobsClient) GetJobSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetJobResponder handles the response to the GetJob request. The method always
// closes the http.Response Body.
func (client JobsClient) GetJobResponder(resp *http.Response) (result VMwareJob, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
