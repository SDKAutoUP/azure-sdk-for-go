package cdnapi

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
    "github.com/Azure/azure-sdk-for-go/services/cdn/mgmt/2017-04-02/cdn"
)

        // BaseClientAPI contains the set of methods on the BaseClient type.
        type BaseClientAPI interface {
            CheckNameAvailability(ctx context.Context, checkNameAvailabilityInput cdn.CheckNameAvailabilityInput) (result cdn.CheckNameAvailabilityOutput, err error)
            ValidateProbe(ctx context.Context, validateProbeInput cdn.ValidateProbeInput) (result cdn.ValidateProbeOutput, err error)
        }

        var _ BaseClientAPI = (*cdn.BaseClient)(nil)
        // ProfilesClientAPI contains the set of methods on the ProfilesClient type.
        type ProfilesClientAPI interface {
            Create(ctx context.Context, resourceGroupName string, profileName string, profile cdn.Profile) (result cdn.ProfilesCreateFuture, err error)
            Delete(ctx context.Context, resourceGroupName string, profileName string) (result cdn.ProfilesDeleteFuture, err error)
            GenerateSsoURI(ctx context.Context, resourceGroupName string, profileName string) (result cdn.SsoURI, err error)
            Get(ctx context.Context, resourceGroupName string, profileName string) (result cdn.Profile, err error)
            List(ctx context.Context) (result cdn.ProfileListResultPage, err error)
                ListComplete(ctx context.Context) (result cdn.ProfileListResultIterator, err error)
            ListByResourceGroup(ctx context.Context, resourceGroupName string) (result cdn.ProfileListResultPage, err error)
                ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result cdn.ProfileListResultIterator, err error)
            ListResourceUsage(ctx context.Context, resourceGroupName string, profileName string) (result cdn.ResourceUsageListResultPage, err error)
                ListResourceUsageComplete(ctx context.Context, resourceGroupName string, profileName string) (result cdn.ResourceUsageListResultIterator, err error)
            ListSupportedOptimizationTypes(ctx context.Context, resourceGroupName string, profileName string) (result cdn.SupportedOptimizationTypesListResult, err error)
            Update(ctx context.Context, resourceGroupName string, profileName string, profileUpdateParameters cdn.ProfileUpdateParameters) (result cdn.ProfilesUpdateFuture, err error)
        }

        var _ ProfilesClientAPI = (*cdn.ProfilesClient)(nil)
        // EndpointsClientAPI contains the set of methods on the EndpointsClient type.
        type EndpointsClientAPI interface {
            Create(ctx context.Context, resourceGroupName string, profileName string, endpointName string, endpoint cdn.Endpoint) (result cdn.EndpointsCreateFuture, err error)
            Delete(ctx context.Context, resourceGroupName string, profileName string, endpointName string) (result cdn.EndpointsDeleteFuture, err error)
            Get(ctx context.Context, resourceGroupName string, profileName string, endpointName string) (result cdn.Endpoint, err error)
            ListByProfile(ctx context.Context, resourceGroupName string, profileName string) (result cdn.EndpointListResultPage, err error)
                ListByProfileComplete(ctx context.Context, resourceGroupName string, profileName string) (result cdn.EndpointListResultIterator, err error)
            ListResourceUsage(ctx context.Context, resourceGroupName string, profileName string, endpointName string) (result cdn.ResourceUsageListResultPage, err error)
                ListResourceUsageComplete(ctx context.Context, resourceGroupName string, profileName string, endpointName string) (result cdn.ResourceUsageListResultIterator, err error)
            LoadContent(ctx context.Context, resourceGroupName string, profileName string, endpointName string, contentFilePaths cdn.LoadParameters) (result cdn.EndpointsLoadContentFuture, err error)
            PurgeContent(ctx context.Context, resourceGroupName string, profileName string, endpointName string, contentFilePaths cdn.PurgeParameters) (result cdn.EndpointsPurgeContentFuture, err error)
            Start(ctx context.Context, resourceGroupName string, profileName string, endpointName string) (result cdn.EndpointsStartFuture, err error)
            Stop(ctx context.Context, resourceGroupName string, profileName string, endpointName string) (result cdn.EndpointsStopFuture, err error)
            Update(ctx context.Context, resourceGroupName string, profileName string, endpointName string, endpointUpdateProperties cdn.EndpointUpdateParameters) (result cdn.EndpointsUpdateFuture, err error)
            ValidateCustomDomain(ctx context.Context, resourceGroupName string, profileName string, endpointName string, customDomainProperties cdn.ValidateCustomDomainInput) (result cdn.ValidateCustomDomainOutput, err error)
        }

        var _ EndpointsClientAPI = (*cdn.EndpointsClient)(nil)
        // OriginsClientAPI contains the set of methods on the OriginsClient type.
        type OriginsClientAPI interface {
            Get(ctx context.Context, resourceGroupName string, profileName string, endpointName string, originName string) (result cdn.Origin, err error)
            ListByEndpoint(ctx context.Context, resourceGroupName string, profileName string, endpointName string) (result cdn.OriginListResultPage, err error)
                ListByEndpointComplete(ctx context.Context, resourceGroupName string, profileName string, endpointName string) (result cdn.OriginListResultIterator, err error)
            Update(ctx context.Context, resourceGroupName string, profileName string, endpointName string, originName string, originUpdateProperties cdn.OriginUpdateParameters) (result cdn.OriginsUpdateFuture, err error)
        }

        var _ OriginsClientAPI = (*cdn.OriginsClient)(nil)
        // CustomDomainsClientAPI contains the set of methods on the CustomDomainsClient type.
        type CustomDomainsClientAPI interface {
            Create(ctx context.Context, resourceGroupName string, profileName string, endpointName string, customDomainName string, customDomainProperties cdn.CustomDomainParameters) (result cdn.CustomDomainsCreateFuture, err error)
            Delete(ctx context.Context, resourceGroupName string, profileName string, endpointName string, customDomainName string) (result cdn.CustomDomainsDeleteFuture, err error)
            DisableCustomHTTPS(ctx context.Context, resourceGroupName string, profileName string, endpointName string, customDomainName string) (result cdn.CustomDomain, err error)
            EnableCustomHTTPS(ctx context.Context, resourceGroupName string, profileName string, endpointName string, customDomainName string) (result cdn.CustomDomain, err error)
            Get(ctx context.Context, resourceGroupName string, profileName string, endpointName string, customDomainName string) (result cdn.CustomDomain, err error)
            ListByEndpoint(ctx context.Context, resourceGroupName string, profileName string, endpointName string) (result cdn.CustomDomainListResultPage, err error)
                ListByEndpointComplete(ctx context.Context, resourceGroupName string, profileName string, endpointName string) (result cdn.CustomDomainListResultIterator, err error)
        }

        var _ CustomDomainsClientAPI = (*cdn.CustomDomainsClient)(nil)
        // ResourceUsageClientAPI contains the set of methods on the ResourceUsageClient type.
        type ResourceUsageClientAPI interface {
            List(ctx context.Context) (result cdn.ResourceUsageListResultPage, err error)
                ListComplete(ctx context.Context) (result cdn.ResourceUsageListResultIterator, err error)
        }

        var _ ResourceUsageClientAPI = (*cdn.ResourceUsageClient)(nil)
        // OperationsClientAPI contains the set of methods on the OperationsClient type.
        type OperationsClientAPI interface {
            List(ctx context.Context) (result cdn.OperationsListResultPage, err error)
                ListComplete(ctx context.Context) (result cdn.OperationsListResultIterator, err error)
        }

        var _ OperationsClientAPI = (*cdn.OperationsClient)(nil)
        // EdgeNodesClientAPI contains the set of methods on the EdgeNodesClient type.
        type EdgeNodesClientAPI interface {
            List(ctx context.Context) (result cdn.EdgenodeResultPage, err error)
                ListComplete(ctx context.Context) (result cdn.EdgenodeResultIterator, err error)
        }

        var _ EdgeNodesClientAPI = (*cdn.EdgeNodesClient)(nil)
