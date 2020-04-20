package subscriptionsapi

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
	"github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2019-11-01/subscriptions"
)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result subscriptions.OperationListResultPage, err error)
	ListComplete(ctx context.Context) (result subscriptions.OperationListResultIterator, err error)
}

var _ OperationsClientAPI = (*subscriptions.OperationsClient)(nil)

// ClientAPI contains the set of methods on the Client type.
type ClientAPI interface {
	Get(ctx context.Context, subscriptionID string) (result subscriptions.Subscription, err error)
	List(ctx context.Context) (result subscriptions.ListResultPage, err error)
	ListComplete(ctx context.Context) (result subscriptions.ListResultIterator, err error)
	ListLocations(ctx context.Context, subscriptionID string) (result subscriptions.LocationListResult, err error)
}

var _ ClientAPI = (*subscriptions.Client)(nil)

// TenantsClientAPI contains the set of methods on the TenantsClient type.
type TenantsClientAPI interface {
	List(ctx context.Context) (result subscriptions.TenantListResultPage, err error)
	ListComplete(ctx context.Context) (result subscriptions.TenantListResultIterator, err error)
}

var _ TenantsClientAPI = (*subscriptions.TenantsClient)(nil)
