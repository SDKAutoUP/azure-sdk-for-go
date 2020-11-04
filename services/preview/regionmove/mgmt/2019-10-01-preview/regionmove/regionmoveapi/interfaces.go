package regionmoveapi

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
	"github.com/Azure/azure-sdk-for-go/services/preview/regionmove/mgmt/2019-10-01-preview/regionmove"
)

// MoveCollectionsClientAPI contains the set of methods on the MoveCollectionsClient type.
type MoveCollectionsClientAPI interface {
	BulkRemove(ctx context.Context, resourceGroupName string, moveCollectionName string, body *regionmove.BulkRemoveRequest) (result regionmove.MoveCollectionsBulkRemoveFuture, err error)
	Commit(ctx context.Context, resourceGroupName string, moveCollectionName string, body *regionmove.CommitRequest) (result regionmove.MoveCollectionsCommitFuture, err error)
	Create(ctx context.Context, resourceGroupName string, moveCollectionName string, body *regionmove.MoveCollection) (result regionmove.MoveCollection, err error)
	Delete(ctx context.Context, resourceGroupName string, moveCollectionName string) (result regionmove.MoveCollectionsDeleteFuture, err error)
	Discard(ctx context.Context, resourceGroupName string, moveCollectionName string, body *regionmove.DiscardRequest) (result regionmove.MoveCollectionsDiscardFuture, err error)
	Get(ctx context.Context, resourceGroupName string, moveCollectionName string) (result regionmove.MoveCollection, err error)
	InitiateMove(ctx context.Context, resourceGroupName string, moveCollectionName string, body *regionmove.ResourceMoveRequest) (result regionmove.MoveCollectionsInitiateMoveFuture, err error)
	ListMoveCollectionsByResourceGroup(ctx context.Context, resourceGroupName string) (result regionmove.MoveCollectionResultListPage, err error)
	ListMoveCollectionsByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result regionmove.MoveCollectionResultListIterator, err error)
	ListMoveCollectionsBySubscription(ctx context.Context) (result regionmove.MoveCollectionResultListPage, err error)
	ListMoveCollectionsBySubscriptionComplete(ctx context.Context) (result regionmove.MoveCollectionResultListIterator, err error)
	Prepare(ctx context.Context, resourceGroupName string, moveCollectionName string, body *regionmove.PrepareRequest) (result regionmove.MoveCollectionsPrepareFuture, err error)
	ResolveDependencies(ctx context.Context, resourceGroupName string, moveCollectionName string) (result regionmove.MoveCollectionsResolveDependenciesFuture, err error)
	Update(ctx context.Context, resourceGroupName string, moveCollectionName string, body *regionmove.UpdateMoveCollectionRequest) (result regionmove.MoveCollection, err error)
}

var _ MoveCollectionsClientAPI = (*regionmove.MoveCollectionsClient)(nil)

// MoveResourcesClientAPI contains the set of methods on the MoveResourcesClient type.
type MoveResourcesClientAPI interface {
	Create(ctx context.Context, resourceGroupName string, moveCollectionName string, moveResourceName string, body *regionmove.MoveResource) (result regionmove.MoveResourcesCreateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, moveCollectionName string, moveResourceName string) (result regionmove.MoveResourcesDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, moveCollectionName string, moveResourceName string) (result regionmove.MoveResource, err error)
	List(ctx context.Context, resourceGroupName string, moveCollectionName string, filter string) (result regionmove.MoveResourceCollectionPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, moveCollectionName string, filter string) (result regionmove.MoveResourceCollectionIterator, err error)
}

var _ MoveResourcesClientAPI = (*regionmove.MoveResourcesClient)(nil)

// UnresolvedDependenciesClientAPI contains the set of methods on the UnresolvedDependenciesClient type.
type UnresolvedDependenciesClientAPI interface {
	Get(ctx context.Context, resourceGroupName string, moveCollectionName string) (result regionmove.UnresolvedDependencyCollection, err error)
}

var _ UnresolvedDependenciesClientAPI = (*regionmove.UnresolvedDependenciesClient)(nil)

// OperationsDiscoveryClientAPI contains the set of methods on the OperationsDiscoveryClient type.
type OperationsDiscoveryClientAPI interface {
	Get(ctx context.Context) (result regionmove.OperationsDiscoveryCollection, err error)
}

var _ OperationsDiscoveryClientAPI = (*regionmove.OperationsDiscoveryClient)(nil)
