package reservations

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

// AppliedScopeType enumerates the values for applied scope type.
type AppliedScopeType string

const (
	// Shared ...
	Shared AppliedScopeType = "Shared"
	// Single ...
	Single AppliedScopeType = "Single"
)

// PossibleAppliedScopeTypeValues returns an array of possible values for the AppliedScopeType const type.
func PossibleAppliedScopeTypeValues() []AppliedScopeType {
	return []AppliedScopeType{Shared, Single}
}

// CalculateExchangeOperationResultStatus enumerates the values for calculate exchange operation result status.
type CalculateExchangeOperationResultStatus string

const (
	// Cancelled ...
	Cancelled CalculateExchangeOperationResultStatus = "Cancelled"
	// Failed ...
	Failed CalculateExchangeOperationResultStatus = "Failed"
	// Pending ...
	Pending CalculateExchangeOperationResultStatus = "Pending"
	// Succeeded ...
	Succeeded CalculateExchangeOperationResultStatus = "Succeeded"
)

// PossibleCalculateExchangeOperationResultStatusValues returns an array of possible values for the CalculateExchangeOperationResultStatus const type.
func PossibleCalculateExchangeOperationResultStatusValues() []CalculateExchangeOperationResultStatus {
	return []CalculateExchangeOperationResultStatus{Cancelled, Failed, Pending, Succeeded}
}

// ErrorResponseCode enumerates the values for error response code.
type ErrorResponseCode string

const (
	// ActivateQuoteFailed ...
	ActivateQuoteFailed ErrorResponseCode = "ActivateQuoteFailed"
	// AppliedScopesNotAssociatedWithCommerceAccount ...
	AppliedScopesNotAssociatedWithCommerceAccount ErrorResponseCode = "AppliedScopesNotAssociatedWithCommerceAccount"
	// AuthorizationFailed ...
	AuthorizationFailed ErrorResponseCode = "AuthorizationFailed"
	// BadRequest ...
	BadRequest ErrorResponseCode = "BadRequest"
	// BillingCustomerInputError ...
	BillingCustomerInputError ErrorResponseCode = "BillingCustomerInputError"
	// BillingError ...
	BillingError ErrorResponseCode = "BillingError"
	// BillingPaymentInstrumentHardError ...
	BillingPaymentInstrumentHardError ErrorResponseCode = "BillingPaymentInstrumentHardError"
	// BillingPaymentInstrumentSoftError ...
	BillingPaymentInstrumentSoftError ErrorResponseCode = "BillingPaymentInstrumentSoftError"
	// BillingScopeIDCannotBeChanged ...
	BillingScopeIDCannotBeChanged ErrorResponseCode = "BillingScopeIdCannotBeChanged"
	// BillingTransientError ...
	BillingTransientError ErrorResponseCode = "BillingTransientError"
	// CalculatePriceFailed ...
	CalculatePriceFailed ErrorResponseCode = "CalculatePriceFailed"
	// CapacityUpdateScopesFailed ...
	CapacityUpdateScopesFailed ErrorResponseCode = "CapacityUpdateScopesFailed"
	// ClientCertificateThumbprintNotSet ...
	ClientCertificateThumbprintNotSet ErrorResponseCode = "ClientCertificateThumbprintNotSet"
	// CreateQuoteFailed ...
	CreateQuoteFailed ErrorResponseCode = "CreateQuoteFailed"
	// Forbidden ...
	Forbidden ErrorResponseCode = "Forbidden"
	// FulfillmentConfigurationError ...
	FulfillmentConfigurationError ErrorResponseCode = "FulfillmentConfigurationError"
	// FulfillmentError ...
	FulfillmentError ErrorResponseCode = "FulfillmentError"
	// FulfillmentOutOfStockError ...
	FulfillmentOutOfStockError ErrorResponseCode = "FulfillmentOutOfStockError"
	// FulfillmentTransientError ...
	FulfillmentTransientError ErrorResponseCode = "FulfillmentTransientError"
	// HTTPMethodNotSupported ...
	HTTPMethodNotSupported ErrorResponseCode = "HttpMethodNotSupported"
	// InternalServerError ...
	InternalServerError ErrorResponseCode = "InternalServerError"
	// InvalidAccessToken ...
	InvalidAccessToken ErrorResponseCode = "InvalidAccessToken"
	// InvalidFulfillmentRequestParameters ...
	InvalidFulfillmentRequestParameters ErrorResponseCode = "InvalidFulfillmentRequestParameters"
	// InvalidHealthCheckType ...
	InvalidHealthCheckType ErrorResponseCode = "InvalidHealthCheckType"
	// InvalidLocationID ...
	InvalidLocationID ErrorResponseCode = "InvalidLocationId"
	// InvalidRefundQuantity ...
	InvalidRefundQuantity ErrorResponseCode = "InvalidRefundQuantity"
	// InvalidRequestContent ...
	InvalidRequestContent ErrorResponseCode = "InvalidRequestContent"
	// InvalidRequestURI ...
	InvalidRequestURI ErrorResponseCode = "InvalidRequestUri"
	// InvalidReservationID ...
	InvalidReservationID ErrorResponseCode = "InvalidReservationId"
	// InvalidReservationOrderID ...
	InvalidReservationOrderID ErrorResponseCode = "InvalidReservationOrderId"
	// InvalidSingleAppliedScopesCount ...
	InvalidSingleAppliedScopesCount ErrorResponseCode = "InvalidSingleAppliedScopesCount"
	// InvalidSubscriptionID ...
	InvalidSubscriptionID ErrorResponseCode = "InvalidSubscriptionId"
	// InvalidTenantID ...
	InvalidTenantID ErrorResponseCode = "InvalidTenantId"
	// MissingAppliedScopesForSingle ...
	MissingAppliedScopesForSingle ErrorResponseCode = "MissingAppliedScopesForSingle"
	// MissingTenantID ...
	MissingTenantID ErrorResponseCode = "MissingTenantId"
	// NonsupportedAccountID ...
	NonsupportedAccountID ErrorResponseCode = "NonsupportedAccountId"
	// NotSpecified ...
	NotSpecified ErrorResponseCode = "NotSpecified"
	// NotSupportedCountry ...
	NotSupportedCountry ErrorResponseCode = "NotSupportedCountry"
	// NoValidReservationsToReRate ...
	NoValidReservationsToReRate ErrorResponseCode = "NoValidReservationsToReRate"
	// OperationCannotBePerformedInCurrentState ...
	OperationCannotBePerformedInCurrentState ErrorResponseCode = "OperationCannotBePerformedInCurrentState"
	// OperationFailed ...
	OperationFailed ErrorResponseCode = "OperationFailed"
	// PatchValuesSameAsExisting ...
	PatchValuesSameAsExisting ErrorResponseCode = "PatchValuesSameAsExisting"
	// PaymentInstrumentNotFound ...
	PaymentInstrumentNotFound ErrorResponseCode = "PaymentInstrumentNotFound"
	// PurchaseError ...
	PurchaseError ErrorResponseCode = "PurchaseError"
	// ReRateOnlyAllowedForEA ...
	ReRateOnlyAllowedForEA ErrorResponseCode = "ReRateOnlyAllowedForEA"
	// ReservationIDNotInReservationOrder ...
	ReservationIDNotInReservationOrder ErrorResponseCode = "ReservationIdNotInReservationOrder"
	// ReservationOrderCreationFailed ...
	ReservationOrderCreationFailed ErrorResponseCode = "ReservationOrderCreationFailed"
	// ReservationOrderIDAlreadyExists ...
	ReservationOrderIDAlreadyExists ErrorResponseCode = "ReservationOrderIdAlreadyExists"
	// ReservationOrderNotEnabled ...
	ReservationOrderNotEnabled ErrorResponseCode = "ReservationOrderNotEnabled"
	// ReservationOrderNotFound ...
	ReservationOrderNotFound ErrorResponseCode = "ReservationOrderNotFound"
	// RiskCheckFailed ...
	RiskCheckFailed ErrorResponseCode = "RiskCheckFailed"
	// RoleAssignmentCreationFailed ...
	RoleAssignmentCreationFailed ErrorResponseCode = "RoleAssignmentCreationFailed"
	// ServerTimeout ...
	ServerTimeout ErrorResponseCode = "ServerTimeout"
	// UnauthenticatedRequestsThrottled ...
	UnauthenticatedRequestsThrottled ErrorResponseCode = "UnauthenticatedRequestsThrottled"
	// UnsupportedReservationTerm ...
	UnsupportedReservationTerm ErrorResponseCode = "UnsupportedReservationTerm"
)

// PossibleErrorResponseCodeValues returns an array of possible values for the ErrorResponseCode const type.
func PossibleErrorResponseCodeValues() []ErrorResponseCode {
	return []ErrorResponseCode{ActivateQuoteFailed, AppliedScopesNotAssociatedWithCommerceAccount, AuthorizationFailed, BadRequest, BillingCustomerInputError, BillingError, BillingPaymentInstrumentHardError, BillingPaymentInstrumentSoftError, BillingScopeIDCannotBeChanged, BillingTransientError, CalculatePriceFailed, CapacityUpdateScopesFailed, ClientCertificateThumbprintNotSet, CreateQuoteFailed, Forbidden, FulfillmentConfigurationError, FulfillmentError, FulfillmentOutOfStockError, FulfillmentTransientError, HTTPMethodNotSupported, InternalServerError, InvalidAccessToken, InvalidFulfillmentRequestParameters, InvalidHealthCheckType, InvalidLocationID, InvalidRefundQuantity, InvalidRequestContent, InvalidRequestURI, InvalidReservationID, InvalidReservationOrderID, InvalidSingleAppliedScopesCount, InvalidSubscriptionID, InvalidTenantID, MissingAppliedScopesForSingle, MissingTenantID, NonsupportedAccountID, NotSpecified, NotSupportedCountry, NoValidReservationsToReRate, OperationCannotBePerformedInCurrentState, OperationFailed, PatchValuesSameAsExisting, PaymentInstrumentNotFound, PurchaseError, ReRateOnlyAllowedForEA, ReservationIDNotInReservationOrder, ReservationOrderCreationFailed, ReservationOrderIDAlreadyExists, ReservationOrderNotEnabled, ReservationOrderNotFound, RiskCheckFailed, RoleAssignmentCreationFailed, ServerTimeout, UnauthenticatedRequestsThrottled, UnsupportedReservationTerm}
}

// ExchangeOperationResultStatus enumerates the values for exchange operation result status.
type ExchangeOperationResultStatus string

const (
	// ExchangeOperationResultStatusCancelled ...
	ExchangeOperationResultStatusCancelled ExchangeOperationResultStatus = "Cancelled"
	// ExchangeOperationResultStatusFailed ...
	ExchangeOperationResultStatusFailed ExchangeOperationResultStatus = "Failed"
	// ExchangeOperationResultStatusPendingPurchases ...
	ExchangeOperationResultStatusPendingPurchases ExchangeOperationResultStatus = "PendingPurchases"
	// ExchangeOperationResultStatusPendingRefunds ...
	ExchangeOperationResultStatusPendingRefunds ExchangeOperationResultStatus = "PendingRefunds"
	// ExchangeOperationResultStatusSucceeded ...
	ExchangeOperationResultStatusSucceeded ExchangeOperationResultStatus = "Succeeded"
)

// PossibleExchangeOperationResultStatusValues returns an array of possible values for the ExchangeOperationResultStatus const type.
func PossibleExchangeOperationResultStatusValues() []ExchangeOperationResultStatus {
	return []ExchangeOperationResultStatus{ExchangeOperationResultStatusCancelled, ExchangeOperationResultStatusFailed, ExchangeOperationResultStatusPendingPurchases, ExchangeOperationResultStatusPendingRefunds, ExchangeOperationResultStatusSucceeded}
}

// InstanceFlexibility enumerates the values for instance flexibility.
type InstanceFlexibility string

const (
	// Off ...
	Off InstanceFlexibility = "Off"
	// On ...
	On InstanceFlexibility = "On"
)

// PossibleInstanceFlexibilityValues returns an array of possible values for the InstanceFlexibility const type.
func PossibleInstanceFlexibilityValues() []InstanceFlexibility {
	return []InstanceFlexibility{Off, On}
}

// PaymentStatus enumerates the values for payment status.
type PaymentStatus string

const (
	// PaymentStatusCancelled ...
	PaymentStatusCancelled PaymentStatus = "Cancelled"
	// PaymentStatusFailed ...
	PaymentStatusFailed PaymentStatus = "Failed"
	// PaymentStatusScheduled ...
	PaymentStatusScheduled PaymentStatus = "Scheduled"
	// PaymentStatusSucceeded ...
	PaymentStatusSucceeded PaymentStatus = "Succeeded"
)

// PossiblePaymentStatusValues returns an array of possible values for the PaymentStatus const type.
func PossiblePaymentStatusValues() []PaymentStatus {
	return []PaymentStatus{PaymentStatusCancelled, PaymentStatusFailed, PaymentStatusScheduled, PaymentStatusSucceeded}
}

// ReservationBillingPlan enumerates the values for reservation billing plan.
type ReservationBillingPlan string

const (
	// Monthly ...
	Monthly ReservationBillingPlan = "Monthly"
	// Upfront ...
	Upfront ReservationBillingPlan = "Upfront"
)

// PossibleReservationBillingPlanValues returns an array of possible values for the ReservationBillingPlan const type.
func PossibleReservationBillingPlanValues() []ReservationBillingPlan {
	return []ReservationBillingPlan{Monthly, Upfront}
}

// ReservationTerm enumerates the values for reservation term.
type ReservationTerm string

const (
	// P1Y ...
	P1Y ReservationTerm = "P1Y"
	// P3Y ...
	P3Y ReservationTerm = "P3Y"
)

// PossibleReservationTermValues returns an array of possible values for the ReservationTerm const type.
func PossibleReservationTermValues() []ReservationTerm {
	return []ReservationTerm{P1Y, P3Y}
}

// ReservedResourceType enumerates the values for reserved resource type.
type ReservedResourceType string

const (
	// AppService ...
	AppService ReservedResourceType = "AppService"
	// AzureDataExplorer ...
	AzureDataExplorer ReservedResourceType = "AzureDataExplorer"
	// BlockBlob ...
	BlockBlob ReservedResourceType = "BlockBlob"
	// CosmosDb ...
	CosmosDb ReservedResourceType = "CosmosDb"
	// Databricks ...
	Databricks ReservedResourceType = "Databricks"
	// DedicatedHost ...
	DedicatedHost ReservedResourceType = "DedicatedHost"
	// ManagedDisk ...
	ManagedDisk ReservedResourceType = "ManagedDisk"
	// MariaDb ...
	MariaDb ReservedResourceType = "MariaDb"
	// MySQL ...
	MySQL ReservedResourceType = "MySql"
	// PostgreSQL ...
	PostgreSQL ReservedResourceType = "PostgreSql"
	// RedHat ...
	RedHat ReservedResourceType = "RedHat"
	// RedHatOsa ...
	RedHatOsa ReservedResourceType = "RedHatOsa"
	// RedisCache ...
	RedisCache ReservedResourceType = "RedisCache"
	// SapHana ...
	SapHana ReservedResourceType = "SapHana"
	// SQLAzureHybridBenefit ...
	SQLAzureHybridBenefit ReservedResourceType = "SqlAzureHybridBenefit"
	// SQLDatabases ...
	SQLDatabases ReservedResourceType = "SqlDatabases"
	// SQLDataWarehouse ...
	SQLDataWarehouse ReservedResourceType = "SqlDataWarehouse"
	// SuseLinux ...
	SuseLinux ReservedResourceType = "SuseLinux"
	// VirtualMachines ...
	VirtualMachines ReservedResourceType = "VirtualMachines"
	// VMwareCloudSimple ...
	VMwareCloudSimple ReservedResourceType = "VMwareCloudSimple"
)

// PossibleReservedResourceTypeValues returns an array of possible values for the ReservedResourceType const type.
func PossibleReservedResourceTypeValues() []ReservedResourceType {
	return []ReservedResourceType{AppService, AzureDataExplorer, BlockBlob, CosmosDb, Databricks, DedicatedHost, ManagedDisk, MariaDb, MySQL, PostgreSQL, RedHat, RedHatOsa, RedisCache, SapHana, SQLAzureHybridBenefit, SQLDatabases, SQLDataWarehouse, SuseLinux, VirtualMachines, VMwareCloudSimple}
}

// Scope enumerates the values for scope.
type Scope string

const (
	// Reservation ...
	Reservation Scope = "Reservation"
)

// PossibleScopeValues returns an array of possible values for the Scope const type.
func PossibleScopeValues() []Scope {
	return []Scope{Reservation}
}

// StatusCode enumerates the values for status code.
type StatusCode string

const (
	// StatusCodeActive ...
	StatusCodeActive StatusCode = "Active"
	// StatusCodeExpired ...
	StatusCodeExpired StatusCode = "Expired"
	// StatusCodeMerged ...
	StatusCodeMerged StatusCode = "Merged"
	// StatusCodeNone ...
	StatusCodeNone StatusCode = "None"
	// StatusCodePaymentInstrumentError ...
	StatusCodePaymentInstrumentError StatusCode = "PaymentInstrumentError"
	// StatusCodePending ...
	StatusCodePending StatusCode = "Pending"
	// StatusCodePurchaseError ...
	StatusCodePurchaseError StatusCode = "PurchaseError"
	// StatusCodeSplit ...
	StatusCodeSplit StatusCode = "Split"
	// StatusCodeSucceeded ...
	StatusCodeSucceeded StatusCode = "Succeeded"
)

// PossibleStatusCodeValues returns an array of possible values for the StatusCode const type.
func PossibleStatusCodeValues() []StatusCode {
	return []StatusCode{StatusCodeActive, StatusCodeExpired, StatusCodeMerged, StatusCodeNone, StatusCodePaymentInstrumentError, StatusCodePending, StatusCodePurchaseError, StatusCodeSplit, StatusCodeSucceeded}
}
