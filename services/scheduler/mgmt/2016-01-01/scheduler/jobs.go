package scheduler

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
	"net/http"
)

// Deprecated: Please use github.com/Azure/azure-sdk-for-go/services/logic/mgmt/2016-06-01/logic instead.
// JobsClient is the client for the Jobs methods of the Scheduler service.
type JobsClient struct {
	BaseClient
}

// Deprecated: Please use github.com/Azure/azure-sdk-for-go/services/logic/mgmt/2016-06-01/logic instead.
// NewJobsClient creates an instance of the JobsClient client.
func NewJobsClient(subscriptionID string) JobsClient {
	return NewJobsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// Deprecated: Please use github.com/Azure/azure-sdk-for-go/services/logic/mgmt/2016-06-01/logic instead.
// NewJobsClientWithBaseURI creates an instance of the JobsClient client.
func NewJobsClientWithBaseURI(baseURI string, subscriptionID string) JobsClient {
	return JobsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Deprecated: Please use github.com/Azure/azure-sdk-for-go/services/logic/mgmt/2016-06-01/logic instead.
// CreateOrUpdate provisions a new job or updates an existing job.
// Parameters:
// resourceGroupName - the resource group name.
// jobCollectionName - the job collection name.
// jobName - the job name.
// job - the job definition.
func (client JobsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, jobCollectionName string, jobName string, job JobDefinition) (result JobDefinition, err error) {
	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, jobCollectionName, jobName, job)
	if err != nil {
		err = autorest.NewErrorWithError(err, "scheduler.JobsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "scheduler.JobsClient", "CreateOrUpdate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "scheduler.JobsClient", "CreateOrUpdate", resp, "Failure responding to request")
	}

	return
}

// Deprecated: Please use github.com/Azure/azure-sdk-for-go/services/logic/mgmt/2016-06-01/logic instead.
// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client JobsClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, jobCollectionName string, jobName string, job JobDefinition) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"jobCollectionName": autorest.Encode("path", jobCollectionName),
		"jobName":           autorest.Encode("path", jobName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-01-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Scheduler/jobCollections/{jobCollectionName}/jobs/{jobName}", pathParameters),
		autorest.WithJSON(job),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// Deprecated: Please use github.com/Azure/azure-sdk-for-go/services/logic/mgmt/2016-06-01/logic instead.
// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client JobsClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// Deprecated: Please use github.com/Azure/azure-sdk-for-go/services/logic/mgmt/2016-06-01/logic instead.
// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client JobsClient) CreateOrUpdateResponder(resp *http.Response) (result JobDefinition, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Deprecated: Please use github.com/Azure/azure-sdk-for-go/services/logic/mgmt/2016-06-01/logic instead.
// Delete deletes a job.
// Parameters:
// resourceGroupName - the resource group name.
// jobCollectionName - the job collection name.
// jobName - the job name.
func (client JobsClient) Delete(ctx context.Context, resourceGroupName string, jobCollectionName string, jobName string) (result autorest.Response, err error) {
	req, err := client.DeletePreparer(ctx, resourceGroupName, jobCollectionName, jobName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "scheduler.JobsClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "scheduler.JobsClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "scheduler.JobsClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// Deprecated: Please use github.com/Azure/azure-sdk-for-go/services/logic/mgmt/2016-06-01/logic instead.
// DeletePreparer prepares the Delete request.
func (client JobsClient) DeletePreparer(ctx context.Context, resourceGroupName string, jobCollectionName string, jobName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"jobCollectionName": autorest.Encode("path", jobCollectionName),
		"jobName":           autorest.Encode("path", jobName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-01-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Scheduler/jobCollections/{jobCollectionName}/jobs/{jobName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// Deprecated: Please use github.com/Azure/azure-sdk-for-go/services/logic/mgmt/2016-06-01/logic instead.
// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client JobsClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// Deprecated: Please use github.com/Azure/azure-sdk-for-go/services/logic/mgmt/2016-06-01/logic instead.
// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client JobsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Deprecated: Please use github.com/Azure/azure-sdk-for-go/services/logic/mgmt/2016-06-01/logic instead.
// Get gets a job.
// Parameters:
// resourceGroupName - the resource group name.
// jobCollectionName - the job collection name.
// jobName - the job name.
func (client JobsClient) Get(ctx context.Context, resourceGroupName string, jobCollectionName string, jobName string) (result JobDefinition, err error) {
	req, err := client.GetPreparer(ctx, resourceGroupName, jobCollectionName, jobName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "scheduler.JobsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "scheduler.JobsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "scheduler.JobsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// Deprecated: Please use github.com/Azure/azure-sdk-for-go/services/logic/mgmt/2016-06-01/logic instead.
// GetPreparer prepares the Get request.
func (client JobsClient) GetPreparer(ctx context.Context, resourceGroupName string, jobCollectionName string, jobName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"jobCollectionName": autorest.Encode("path", jobCollectionName),
		"jobName":           autorest.Encode("path", jobName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-01-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Scheduler/jobCollections/{jobCollectionName}/jobs/{jobName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// Deprecated: Please use github.com/Azure/azure-sdk-for-go/services/logic/mgmt/2016-06-01/logic instead.
// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client JobsClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// Deprecated: Please use github.com/Azure/azure-sdk-for-go/services/logic/mgmt/2016-06-01/logic instead.
// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client JobsClient) GetResponder(resp *http.Response) (result JobDefinition, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Deprecated: Please use github.com/Azure/azure-sdk-for-go/services/logic/mgmt/2016-06-01/logic instead.
// List lists all jobs under the specified job collection.
// Parameters:
// resourceGroupName - the resource group name.
// jobCollectionName - the job collection name.
// top - the number of jobs to request, in the of range [1..100].
// skip - the (0-based) index of the job history list from which to begin requesting entries.
// filter - the filter to apply on the job state.
func (client JobsClient) List(ctx context.Context, resourceGroupName string, jobCollectionName string, top *int32, skip *int32, filter string) (result JobListResultPage, err error) {
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, resourceGroupName, jobCollectionName, top, skip, filter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "scheduler.JobsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.jlr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "scheduler.JobsClient", "List", resp, "Failure sending request")
		return
	}

	result.jlr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "scheduler.JobsClient", "List", resp, "Failure responding to request")
	}

	return
}

// Deprecated: Please use github.com/Azure/azure-sdk-for-go/services/logic/mgmt/2016-06-01/logic instead.
// ListPreparer prepares the List request.
func (client JobsClient) ListPreparer(ctx context.Context, resourceGroupName string, jobCollectionName string, top *int32, skip *int32, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"jobCollectionName": autorest.Encode("path", jobCollectionName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-01-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}
	if skip != nil {
		queryParameters["$skip"] = autorest.Encode("query", *skip)
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Scheduler/jobCollections/{jobCollectionName}/jobs", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// Deprecated: Please use github.com/Azure/azure-sdk-for-go/services/logic/mgmt/2016-06-01/logic instead.
// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client JobsClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// Deprecated: Please use github.com/Azure/azure-sdk-for-go/services/logic/mgmt/2016-06-01/logic instead.
// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client JobsClient) ListResponder(resp *http.Response) (result JobListResult, err error) {
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
func (client JobsClient) listNextResults(lastResults JobListResult) (result JobListResult, err error) {
	req, err := lastResults.jobListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "scheduler.JobsClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "scheduler.JobsClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "scheduler.JobsClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// Deprecated: Please use github.com/Azure/azure-sdk-for-go/services/logic/mgmt/2016-06-01/logic instead.
// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client JobsClient) ListComplete(ctx context.Context, resourceGroupName string, jobCollectionName string, top *int32, skip *int32, filter string) (result JobListResultIterator, err error) {
	result.page, err = client.List(ctx, resourceGroupName, jobCollectionName, top, skip, filter)
	return
}

// Deprecated: Please use github.com/Azure/azure-sdk-for-go/services/logic/mgmt/2016-06-01/logic instead.
// ListJobHistory lists job history.
// Parameters:
// resourceGroupName - the resource group name.
// jobCollectionName - the job collection name.
// jobName - the job name.
// top - the number of job history to request, in the of range [1..100].
// skip - the (0-based) index of the job history list from which to begin requesting entries.
// filter - the filter to apply on the job state.
func (client JobsClient) ListJobHistory(ctx context.Context, resourceGroupName string, jobCollectionName string, jobName string, top *int32, skip *int32, filter string) (result JobHistoryListResultPage, err error) {
	result.fn = client.listJobHistoryNextResults
	req, err := client.ListJobHistoryPreparer(ctx, resourceGroupName, jobCollectionName, jobName, top, skip, filter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "scheduler.JobsClient", "ListJobHistory", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListJobHistorySender(req)
	if err != nil {
		result.jhlr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "scheduler.JobsClient", "ListJobHistory", resp, "Failure sending request")
		return
	}

	result.jhlr, err = client.ListJobHistoryResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "scheduler.JobsClient", "ListJobHistory", resp, "Failure responding to request")
	}

	return
}

// Deprecated: Please use github.com/Azure/azure-sdk-for-go/services/logic/mgmt/2016-06-01/logic instead.
// ListJobHistoryPreparer prepares the ListJobHistory request.
func (client JobsClient) ListJobHistoryPreparer(ctx context.Context, resourceGroupName string, jobCollectionName string, jobName string, top *int32, skip *int32, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"jobCollectionName": autorest.Encode("path", jobCollectionName),
		"jobName":           autorest.Encode("path", jobName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-01-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}
	if skip != nil {
		queryParameters["$skip"] = autorest.Encode("query", *skip)
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Scheduler/jobCollections/{jobCollectionName}/jobs/{jobName}/history", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// Deprecated: Please use github.com/Azure/azure-sdk-for-go/services/logic/mgmt/2016-06-01/logic instead.
// ListJobHistorySender sends the ListJobHistory request. The method will close the
// http.Response Body if it receives an error.
func (client JobsClient) ListJobHistorySender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// Deprecated: Please use github.com/Azure/azure-sdk-for-go/services/logic/mgmt/2016-06-01/logic instead.
// ListJobHistoryResponder handles the response to the ListJobHistory request. The method always
// closes the http.Response Body.
func (client JobsClient) ListJobHistoryResponder(resp *http.Response) (result JobHistoryListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listJobHistoryNextResults retrieves the next set of results, if any.
func (client JobsClient) listJobHistoryNextResults(lastResults JobHistoryListResult) (result JobHistoryListResult, err error) {
	req, err := lastResults.jobHistoryListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "scheduler.JobsClient", "listJobHistoryNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListJobHistorySender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "scheduler.JobsClient", "listJobHistoryNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListJobHistoryResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "scheduler.JobsClient", "listJobHistoryNextResults", resp, "Failure responding to next results request")
	}
	return
}

// Deprecated: Please use github.com/Azure/azure-sdk-for-go/services/logic/mgmt/2016-06-01/logic instead.
// ListJobHistoryComplete enumerates all values, automatically crossing page boundaries as required.
func (client JobsClient) ListJobHistoryComplete(ctx context.Context, resourceGroupName string, jobCollectionName string, jobName string, top *int32, skip *int32, filter string) (result JobHistoryListResultIterator, err error) {
	result.page, err = client.ListJobHistory(ctx, resourceGroupName, jobCollectionName, jobName, top, skip, filter)
	return
}

// Deprecated: Please use github.com/Azure/azure-sdk-for-go/services/logic/mgmt/2016-06-01/logic instead.
// Patch patches an existing job.
// Parameters:
// resourceGroupName - the resource group name.
// jobCollectionName - the job collection name.
// jobName - the job name.
// job - the job definition.
func (client JobsClient) Patch(ctx context.Context, resourceGroupName string, jobCollectionName string, jobName string, job JobDefinition) (result JobDefinition, err error) {
	req, err := client.PatchPreparer(ctx, resourceGroupName, jobCollectionName, jobName, job)
	if err != nil {
		err = autorest.NewErrorWithError(err, "scheduler.JobsClient", "Patch", nil, "Failure preparing request")
		return
	}

	resp, err := client.PatchSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "scheduler.JobsClient", "Patch", resp, "Failure sending request")
		return
	}

	result, err = client.PatchResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "scheduler.JobsClient", "Patch", resp, "Failure responding to request")
	}

	return
}

// Deprecated: Please use github.com/Azure/azure-sdk-for-go/services/logic/mgmt/2016-06-01/logic instead.
// PatchPreparer prepares the Patch request.
func (client JobsClient) PatchPreparer(ctx context.Context, resourceGroupName string, jobCollectionName string, jobName string, job JobDefinition) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"jobCollectionName": autorest.Encode("path", jobCollectionName),
		"jobName":           autorest.Encode("path", jobName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-01-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Scheduler/jobCollections/{jobCollectionName}/jobs/{jobName}", pathParameters),
		autorest.WithJSON(job),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// Deprecated: Please use github.com/Azure/azure-sdk-for-go/services/logic/mgmt/2016-06-01/logic instead.
// PatchSender sends the Patch request. The method will close the
// http.Response Body if it receives an error.
func (client JobsClient) PatchSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// Deprecated: Please use github.com/Azure/azure-sdk-for-go/services/logic/mgmt/2016-06-01/logic instead.
// PatchResponder handles the response to the Patch request. The method always
// closes the http.Response Body.
func (client JobsClient) PatchResponder(resp *http.Response) (result JobDefinition, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Deprecated: Please use github.com/Azure/azure-sdk-for-go/services/logic/mgmt/2016-06-01/logic instead.
// Run runs a job.
// Parameters:
// resourceGroupName - the resource group name.
// jobCollectionName - the job collection name.
// jobName - the job name.
func (client JobsClient) Run(ctx context.Context, resourceGroupName string, jobCollectionName string, jobName string) (result autorest.Response, err error) {
	req, err := client.RunPreparer(ctx, resourceGroupName, jobCollectionName, jobName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "scheduler.JobsClient", "Run", nil, "Failure preparing request")
		return
	}

	resp, err := client.RunSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "scheduler.JobsClient", "Run", resp, "Failure sending request")
		return
	}

	result, err = client.RunResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "scheduler.JobsClient", "Run", resp, "Failure responding to request")
	}

	return
}

// Deprecated: Please use github.com/Azure/azure-sdk-for-go/services/logic/mgmt/2016-06-01/logic instead.
// RunPreparer prepares the Run request.
func (client JobsClient) RunPreparer(ctx context.Context, resourceGroupName string, jobCollectionName string, jobName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"jobCollectionName": autorest.Encode("path", jobCollectionName),
		"jobName":           autorest.Encode("path", jobName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-01-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Scheduler/jobCollections/{jobCollectionName}/jobs/{jobName}/run", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// Deprecated: Please use github.com/Azure/azure-sdk-for-go/services/logic/mgmt/2016-06-01/logic instead.
// RunSender sends the Run request. The method will close the
// http.Response Body if it receives an error.
func (client JobsClient) RunSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// Deprecated: Please use github.com/Azure/azure-sdk-for-go/services/logic/mgmt/2016-06-01/logic instead.
// RunResponder handles the response to the Run request. The method always
// closes the http.Response Body.
func (client JobsClient) RunResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}
