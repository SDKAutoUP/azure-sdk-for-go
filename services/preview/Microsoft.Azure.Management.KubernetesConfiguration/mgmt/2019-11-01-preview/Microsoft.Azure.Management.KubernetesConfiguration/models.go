package microsoftazuremanagementkubernetesconfiguration

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
	"encoding/json"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/date"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// The package's fully qualified name.
const fqdn = "github.com/Azure/azure-sdk-for-go/services/preview/microsoft.azure.management.kubernetesconfiguration/mgmt/2019-11-01-preview/microsoft.azure.management.kubernetesconfiguration"

// ComplianceState enumerates the values for compliance state.
type ComplianceState string

const (
	// Compliant ...
	Compliant ComplianceState = "Compliant"
	// Noncompliant ...
	Noncompliant ComplianceState = "Noncompliant"
	// Pending ...
	Pending ComplianceState = "Pending"
)

// PossibleComplianceStateValues returns an array of possible values for the ComplianceState const type.
func PossibleComplianceStateValues() []ComplianceState {
	return []ComplianceState{Compliant, Noncompliant, Pending}
}

// EnableHelmOperator enumerates the values for enable helm operator.
type EnableHelmOperator string

const (
	// False ...
	False EnableHelmOperator = "false"
	// True ...
	True EnableHelmOperator = "true"
)

// PossibleEnableHelmOperatorValues returns an array of possible values for the EnableHelmOperator const type.
func PossibleEnableHelmOperatorValues() []EnableHelmOperator {
	return []EnableHelmOperator{False, True}
}

// MessageLevel enumerates the values for message level.
type MessageLevel string

const (
	// Error ...
	Error MessageLevel = "Error"
	// Information ...
	Information MessageLevel = "Information"
	// Warning ...
	Warning MessageLevel = "Warning"
)

// PossibleMessageLevelValues returns an array of possible values for the MessageLevel const type.
func PossibleMessageLevelValues() []MessageLevel {
	return []MessageLevel{Error, Information, Warning}
}

// OperatorScope enumerates the values for operator scope.
type OperatorScope string

const (
	// Cluster ...
	Cluster OperatorScope = "cluster"
	// Namespace ...
	Namespace OperatorScope = "namespace"
)

// PossibleOperatorScopeValues returns an array of possible values for the OperatorScope const type.
func PossibleOperatorScopeValues() []OperatorScope {
	return []OperatorScope{Cluster, Namespace}
}

// OperatorType enumerates the values for operator type.
type OperatorType string

const (
	// Flux ...
	Flux OperatorType = "Flux"
)

// PossibleOperatorTypeValues returns an array of possible values for the OperatorType const type.
func PossibleOperatorTypeValues() []OperatorType {
	return []OperatorType{Flux}
}

// ProvisioningState enumerates the values for provisioning state.
type ProvisioningState string

const (
	// Accepted ...
	Accepted ProvisioningState = "Accepted"
	// Deleting ...
	Deleting ProvisioningState = "Deleting"
	// Failed ...
	Failed ProvisioningState = "Failed"
	// Running ...
	Running ProvisioningState = "Running"
	// Succeeded ...
	Succeeded ProvisioningState = "Succeeded"
)

// PossibleProvisioningStateValues returns an array of possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{Accepted, Deleting, Failed, Running, Succeeded}
}

// ComplianceStatus compliance Status details
type ComplianceStatus struct {
	// ComplianceState - READ-ONLY; The compliance state of the configuration. Possible values include: 'Pending', 'Compliant', 'Noncompliant'
	ComplianceState ComplianceState `json:"complianceState,omitempty"`
	// LastConfigApplied - Datetime the configuration was last applied.
	LastConfigApplied *date.Time `json:"lastConfigApplied,omitempty"`
	// Message - Message from when the configuration was applied.
	Message *string `json:"message,omitempty"`
	// MessageLevel - Level of the message. Possible values include: 'Error', 'Warning', 'Information'
	MessageLevel MessageLevel `json:"messageLevel,omitempty"`
}

// ErrorDefinition error definition.
type ErrorDefinition struct {
	// Code - READ-ONLY; Service specific error code which serves as the substatus for the HTTP error code.
	Code *string `json:"code,omitempty"`
	// Message - READ-ONLY; Description of the error.
	Message *string `json:"message,omitempty"`
	// Details - READ-ONLY; Internal error details.
	Details *[]ErrorDefinition `json:"details,omitempty"`
}

// ErrorResponse error response.
type ErrorResponse struct {
	// Error - Error definition.
	Error *ErrorDefinition `json:"error,omitempty"`
}

// HelmOperatorProperties properties for Helm operator.
type HelmOperatorProperties struct {
	// ChartVersion - Version of the operator Helm chart.
	ChartVersion *string `json:"chartVersion,omitempty"`
	// ChartValues - Values override for the operator Helm chart.
	ChartValues *string `json:"chartValues,omitempty"`
}

// ProxyResource ARM proxy resource.
type ProxyResource struct {
	// ID - READ-ONLY; Resource Id
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; Resource name
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; Resource type
	Type *string `json:"type,omitempty"`
}

// Resource the Resource model definition.
type Resource struct {
	// ID - READ-ONLY; Resource Id
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; Resource name
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; Resource type
	Type *string `json:"type,omitempty"`
}

// ResourceProviderOperation supported operation of this resource provider.
type ResourceProviderOperation struct {
	// Name - Operation name, in format of {provider}/{resource}/{operation}
	Name *string `json:"name,omitempty"`
	// Display - Display metadata associated with the operation.
	Display *ResourceProviderOperationDisplay `json:"display,omitempty"`
}

// ResourceProviderOperationDisplay display metadata associated with the operation.
type ResourceProviderOperationDisplay struct {
	// Provider - Resource provider: Microsoft KubernetesConfiguration.
	Provider *string `json:"provider,omitempty"`
	// Resource - Resource on which the operation is performed.
	Resource *string `json:"resource,omitempty"`
	// Operation - Type of operation: get, read, delete, etc.
	Operation *string `json:"operation,omitempty"`
	// Description - Description of this operation.
	Description *string `json:"description,omitempty"`
}

// ResourceProviderOperationList result of the request to list operations.
type ResourceProviderOperationList struct {
	autorest.Response `json:"-"`
	// Value - List of operations supported by this resource provider.
	Value *[]ResourceProviderOperation `json:"value,omitempty"`
	// NextLink - READ-ONLY; URL to the next set of results, if any.
	NextLink *string `json:"nextLink,omitempty"`
}

// ResourceProviderOperationListIterator provides access to a complete listing of ResourceProviderOperation
// values.
type ResourceProviderOperationListIterator struct {
	i    int
	page ResourceProviderOperationListPage
}

// NextWithContext advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *ResourceProviderOperationListIterator) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ResourceProviderOperationListIterator.NextWithContext")
		defer func() {
			sc := -1
			if iter.Response().Response.Response != nil {
				sc = iter.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	iter.i++
	if iter.i < len(iter.page.Values()) {
		return nil
	}
	err = iter.page.NextWithContext(ctx)
	if err != nil {
		iter.i--
		return err
	}
	iter.i = 0
	return nil
}

// Next advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (iter *ResourceProviderOperationListIterator) Next() error {
	return iter.NextWithContext(context.Background())
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter ResourceProviderOperationListIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter ResourceProviderOperationListIterator) Response() ResourceProviderOperationList {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter ResourceProviderOperationListIterator) Value() ResourceProviderOperation {
	if !iter.page.NotDone() {
		return ResourceProviderOperation{}
	}
	return iter.page.Values()[iter.i]
}

// Creates a new instance of the ResourceProviderOperationListIterator type.
func NewResourceProviderOperationListIterator(page ResourceProviderOperationListPage) ResourceProviderOperationListIterator {
	return ResourceProviderOperationListIterator{page: page}
}

// IsEmpty returns true if the ListResult contains no values.
func (rpol ResourceProviderOperationList) IsEmpty() bool {
	return rpol.Value == nil || len(*rpol.Value) == 0
}

// resourceProviderOperationListPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (rpol ResourceProviderOperationList) resourceProviderOperationListPreparer(ctx context.Context) (*http.Request, error) {
	if rpol.NextLink == nil || len(to.String(rpol.NextLink)) < 1 {
		return nil, nil
	}
	return autorest.Prepare((&http.Request{}).WithContext(ctx),
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(rpol.NextLink)))
}

// ResourceProviderOperationListPage contains a page of ResourceProviderOperation values.
type ResourceProviderOperationListPage struct {
	fn   func(context.Context, ResourceProviderOperationList) (ResourceProviderOperationList, error)
	rpol ResourceProviderOperationList
}

// NextWithContext advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *ResourceProviderOperationListPage) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ResourceProviderOperationListPage.NextWithContext")
		defer func() {
			sc := -1
			if page.Response().Response.Response != nil {
				sc = page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	next, err := page.fn(ctx, page.rpol)
	if err != nil {
		return err
	}
	page.rpol = next
	return nil
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (page *ResourceProviderOperationListPage) Next() error {
	return page.NextWithContext(context.Background())
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page ResourceProviderOperationListPage) NotDone() bool {
	return !page.rpol.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page ResourceProviderOperationListPage) Response() ResourceProviderOperationList {
	return page.rpol
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page ResourceProviderOperationListPage) Values() []ResourceProviderOperation {
	if page.rpol.IsEmpty() {
		return nil
	}
	return *page.rpol.Value
}

// Creates a new instance of the ResourceProviderOperationListPage type.
func NewResourceProviderOperationListPage(getNextPage func(context.Context, ResourceProviderOperationList) (ResourceProviderOperationList, error)) ResourceProviderOperationListPage {
	return ResourceProviderOperationListPage{fn: getNextPage}
}

// Result sample result definition
type Result struct {
	// SampleProperty - Sample property of type string
	SampleProperty *string `json:"sampleProperty,omitempty"`
}

// SourceControlConfiguration the SourceControl Configuration object.
type SourceControlConfiguration struct {
	autorest.Response `json:"-"`
	// SourceControlConfigurationProperties - Properties to create a Source Control Configuration resource
	*SourceControlConfigurationProperties `json:"properties,omitempty"`
	// ID - READ-ONLY; Resource Id
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; Resource name
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; Resource type
	Type *string `json:"type,omitempty"`
}

// MarshalJSON is the custom marshaler for SourceControlConfiguration.
func (scc SourceControlConfiguration) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if scc.SourceControlConfigurationProperties != nil {
		objectMap["properties"] = scc.SourceControlConfigurationProperties
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for SourceControlConfiguration struct.
func (scc *SourceControlConfiguration) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "properties":
			if v != nil {
				var sourceControlConfigurationProperties SourceControlConfigurationProperties
				err = json.Unmarshal(*v, &sourceControlConfigurationProperties)
				if err != nil {
					return err
				}
				scc.SourceControlConfigurationProperties = &sourceControlConfigurationProperties
			}
		case "id":
			if v != nil {
				var ID string
				err = json.Unmarshal(*v, &ID)
				if err != nil {
					return err
				}
				scc.ID = &ID
			}
		case "name":
			if v != nil {
				var name string
				err = json.Unmarshal(*v, &name)
				if err != nil {
					return err
				}
				scc.Name = &name
			}
		case "type":
			if v != nil {
				var typeVar string
				err = json.Unmarshal(*v, &typeVar)
				if err != nil {
					return err
				}
				scc.Type = &typeVar
			}
		}
	}

	return nil
}

// SourceControlConfigurationList result of the request to list Source Control Configurations.  It contains
// a list of SourceControlConfiguration objects and a URL link to get the next set of results.
type SourceControlConfigurationList struct {
	autorest.Response `json:"-"`
	// Value - READ-ONLY; List of Source Control Configurations within a Kubernetes cluster.
	Value *[]SourceControlConfiguration `json:"value,omitempty"`
	// NextLink - READ-ONLY; URL to get the next set of configuration objects, if any.
	NextLink *string `json:"nextLink,omitempty"`
}

// SourceControlConfigurationListIterator provides access to a complete listing of
// SourceControlConfiguration values.
type SourceControlConfigurationListIterator struct {
	i    int
	page SourceControlConfigurationListPage
}

// NextWithContext advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *SourceControlConfigurationListIterator) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SourceControlConfigurationListIterator.NextWithContext")
		defer func() {
			sc := -1
			if iter.Response().Response.Response != nil {
				sc = iter.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	iter.i++
	if iter.i < len(iter.page.Values()) {
		return nil
	}
	err = iter.page.NextWithContext(ctx)
	if err != nil {
		iter.i--
		return err
	}
	iter.i = 0
	return nil
}

// Next advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (iter *SourceControlConfigurationListIterator) Next() error {
	return iter.NextWithContext(context.Background())
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter SourceControlConfigurationListIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter SourceControlConfigurationListIterator) Response() SourceControlConfigurationList {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter SourceControlConfigurationListIterator) Value() SourceControlConfiguration {
	if !iter.page.NotDone() {
		return SourceControlConfiguration{}
	}
	return iter.page.Values()[iter.i]
}

// Creates a new instance of the SourceControlConfigurationListIterator type.
func NewSourceControlConfigurationListIterator(page SourceControlConfigurationListPage) SourceControlConfigurationListIterator {
	return SourceControlConfigurationListIterator{page: page}
}

// IsEmpty returns true if the ListResult contains no values.
func (sccl SourceControlConfigurationList) IsEmpty() bool {
	return sccl.Value == nil || len(*sccl.Value) == 0
}

// sourceControlConfigurationListPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (sccl SourceControlConfigurationList) sourceControlConfigurationListPreparer(ctx context.Context) (*http.Request, error) {
	if sccl.NextLink == nil || len(to.String(sccl.NextLink)) < 1 {
		return nil, nil
	}
	return autorest.Prepare((&http.Request{}).WithContext(ctx),
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(sccl.NextLink)))
}

// SourceControlConfigurationListPage contains a page of SourceControlConfiguration values.
type SourceControlConfigurationListPage struct {
	fn   func(context.Context, SourceControlConfigurationList) (SourceControlConfigurationList, error)
	sccl SourceControlConfigurationList
}

// NextWithContext advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *SourceControlConfigurationListPage) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SourceControlConfigurationListPage.NextWithContext")
		defer func() {
			sc := -1
			if page.Response().Response.Response != nil {
				sc = page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	next, err := page.fn(ctx, page.sccl)
	if err != nil {
		return err
	}
	page.sccl = next
	return nil
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (page *SourceControlConfigurationListPage) Next() error {
	return page.NextWithContext(context.Background())
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page SourceControlConfigurationListPage) NotDone() bool {
	return !page.sccl.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page SourceControlConfigurationListPage) Response() SourceControlConfigurationList {
	return page.sccl
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page SourceControlConfigurationListPage) Values() []SourceControlConfiguration {
	if page.sccl.IsEmpty() {
		return nil
	}
	return *page.sccl.Value
}

// Creates a new instance of the SourceControlConfigurationListPage type.
func NewSourceControlConfigurationListPage(getNextPage func(context.Context, SourceControlConfigurationList) (SourceControlConfigurationList, error)) SourceControlConfigurationListPage {
	return SourceControlConfigurationListPage{fn: getNextPage}
}

// SourceControlConfigurationProperties properties to create a Source Control Configuration resource
type SourceControlConfigurationProperties struct {
	// RepositoryURL - Url of the SourceControl Repository.
	RepositoryURL *string `json:"repositoryUrl,omitempty"`
	// OperatorNamespace - The namespace to which this operator is installed to. Maximum of 253 lower case alphanumeric characters, hyphen and period only.
	OperatorNamespace *string `json:"operatorNamespace,omitempty"`
	// OperatorInstanceName - Instance name of the operator - identifying the specific configuration.
	OperatorInstanceName *string `json:"operatorInstanceName,omitempty"`
	// OperatorType - Type of the operator. Possible values include: 'Flux'
	OperatorType OperatorType `json:"operatorType,omitempty"`
	// OperatorParams - Any Parameters for the Operator instance in string format.
	OperatorParams *string `json:"operatorParams,omitempty"`
	// OperatorScope - Scope at which the operator will be installed. Possible values include: 'Cluster', 'Namespace'
	OperatorScope OperatorScope `json:"operatorScope,omitempty"`
	// RepositoryPublicKey - READ-ONLY; Public Key associated with this SourceControl configuration (either generated within the cluster or provided by the user).
	RepositoryPublicKey *string `json:"repositoryPublicKey,omitempty"`
	// EnableHelmOperator - Option to enable Helm Operator for this git configuration. Possible values include: 'True', 'False'
	EnableHelmOperator EnableHelmOperator `json:"enableHelmOperator,omitempty"`
	// HelmOperatorProperties - Properties for Helm operator.
	HelmOperatorProperties *HelmOperatorProperties `json:"helmOperatorProperties,omitempty"`
	// ProvisioningState - READ-ONLY; The provisioning state of the resource provider. Possible values include: 'Accepted', 'Deleting', 'Running', 'Succeeded', 'Failed'
	ProvisioningState ProvisioningState `json:"provisioningState,omitempty"`
	// ComplianceStatus - READ-ONLY; Compliance Status of the Configuration
	ComplianceStatus *ComplianceStatus `json:"complianceStatus,omitempty"`
}

// SourceControlConfigurationsDeleteFuture an abstraction for monitoring and retrieving the results of a
// long-running operation.
type SourceControlConfigurationsDeleteFuture struct {
	azure.Future
}

// Result returns the result of the asynchronous operation.
// If the operation has not completed it will return an error.
func (future *SourceControlConfigurationsDeleteFuture) Result(client SourceControlConfigurationsClient) (ar autorest.Response, err error) {
	var done bool
	done, err = future.DoneWithContext(context.Background(), client)
	if err != nil {
		err = autorest.NewErrorWithError(err, "microsoftazuremanagementkubernetesconfiguration.SourceControlConfigurationsDeleteFuture", "Result", future.Response(), "Polling failure")
		return
	}
	if !done {
		err = azure.NewAsyncOpIncompleteError("microsoftazuremanagementkubernetesconfiguration.SourceControlConfigurationsDeleteFuture")
		return
	}
	ar.Response = future.Response()
	return
}
