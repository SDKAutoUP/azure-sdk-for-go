package cdn

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

        // CustomDomainResourceState enumerates the values for custom domain resource state.
    type CustomDomainResourceState string

    const (
                // Active ...
        Active CustomDomainResourceState = "Active"
                // Creating ...
        Creating CustomDomainResourceState = "Creating"
                // Deleting ...
        Deleting CustomDomainResourceState = "Deleting"
            )
    // PossibleCustomDomainResourceStateValues returns an array of possible values for the CustomDomainResourceState const type.
    func PossibleCustomDomainResourceStateValues() []CustomDomainResourceState {
        return []CustomDomainResourceState{Active,Creating,Deleting}
    }

        // EndpointResourceState enumerates the values for endpoint resource state.
    type EndpointResourceState string

    const (
                // EndpointResourceStateCreating ...
        EndpointResourceStateCreating EndpointResourceState = "Creating"
                // EndpointResourceStateDeleting ...
        EndpointResourceStateDeleting EndpointResourceState = "Deleting"
                // EndpointResourceStateRunning ...
        EndpointResourceStateRunning EndpointResourceState = "Running"
                // EndpointResourceStateStarting ...
        EndpointResourceStateStarting EndpointResourceState = "Starting"
                // EndpointResourceStateStopped ...
        EndpointResourceStateStopped EndpointResourceState = "Stopped"
                // EndpointResourceStateStopping ...
        EndpointResourceStateStopping EndpointResourceState = "Stopping"
            )
    // PossibleEndpointResourceStateValues returns an array of possible values for the EndpointResourceState const type.
    func PossibleEndpointResourceStateValues() []EndpointResourceState {
        return []EndpointResourceState{EndpointResourceStateCreating,EndpointResourceStateDeleting,EndpointResourceStateRunning,EndpointResourceStateStarting,EndpointResourceStateStopped,EndpointResourceStateStopping}
    }

        // OriginResourceState enumerates the values for origin resource state.
    type OriginResourceState string

    const (
                // OriginResourceStateActive ...
        OriginResourceStateActive OriginResourceState = "Active"
                // OriginResourceStateCreating ...
        OriginResourceStateCreating OriginResourceState = "Creating"
                // OriginResourceStateDeleting ...
        OriginResourceStateDeleting OriginResourceState = "Deleting"
            )
    // PossibleOriginResourceStateValues returns an array of possible values for the OriginResourceState const type.
    func PossibleOriginResourceStateValues() []OriginResourceState {
        return []OriginResourceState{OriginResourceStateActive,OriginResourceStateCreating,OriginResourceStateDeleting}
    }

        // ProfileResourceState enumerates the values for profile resource state.
    type ProfileResourceState string

    const (
                // ProfileResourceStateActive ...
        ProfileResourceStateActive ProfileResourceState = "Active"
                // ProfileResourceStateCreating ...
        ProfileResourceStateCreating ProfileResourceState = "Creating"
                // ProfileResourceStateDeleting ...
        ProfileResourceStateDeleting ProfileResourceState = "Deleting"
                // ProfileResourceStateDisabled ...
        ProfileResourceStateDisabled ProfileResourceState = "Disabled"
            )
    // PossibleProfileResourceStateValues returns an array of possible values for the ProfileResourceState const type.
    func PossibleProfileResourceStateValues() []ProfileResourceState {
        return []ProfileResourceState{ProfileResourceStateActive,ProfileResourceStateCreating,ProfileResourceStateDeleting,ProfileResourceStateDisabled}
    }

        // ProvisioningState enumerates the values for provisioning state.
    type ProvisioningState string

    const (
                // ProvisioningStateCreating ...
        ProvisioningStateCreating ProvisioningState = "Creating"
                // ProvisioningStateFailed ...
        ProvisioningStateFailed ProvisioningState = "Failed"
                // ProvisioningStateSucceeded ...
        ProvisioningStateSucceeded ProvisioningState = "Succeeded"
            )
    // PossibleProvisioningStateValues returns an array of possible values for the ProvisioningState const type.
    func PossibleProvisioningStateValues() []ProvisioningState {
        return []ProvisioningState{ProvisioningStateCreating,ProvisioningStateFailed,ProvisioningStateSucceeded}
    }

        // QueryStringCachingBehavior enumerates the values for query string caching behavior.
    type QueryStringCachingBehavior string

    const (
                // BypassCaching ...
        BypassCaching QueryStringCachingBehavior = "BypassCaching"
                // IgnoreQueryString ...
        IgnoreQueryString QueryStringCachingBehavior = "IgnoreQueryString"
                // NotSet ...
        NotSet QueryStringCachingBehavior = "NotSet"
                // UseQueryString ...
        UseQueryString QueryStringCachingBehavior = "UseQueryString"
            )
    // PossibleQueryStringCachingBehaviorValues returns an array of possible values for the QueryStringCachingBehavior const type.
    func PossibleQueryStringCachingBehaviorValues() []QueryStringCachingBehavior {
        return []QueryStringCachingBehavior{BypassCaching,IgnoreQueryString,NotSet,UseQueryString}
    }

        // ResourceType enumerates the values for resource type.
    type ResourceType string

    const (
                // MicrosoftCdnProfilesEndpoints ...
        MicrosoftCdnProfilesEndpoints ResourceType = "Microsoft.Cdn/Profiles/Endpoints"
            )
    // PossibleResourceTypeValues returns an array of possible values for the ResourceType const type.
    func PossibleResourceTypeValues() []ResourceType {
        return []ResourceType{MicrosoftCdnProfilesEndpoints}
    }

        // SkuName enumerates the values for sku name.
    type SkuName string

    const (
                // Premium ...
        Premium SkuName = "Premium"
                // Standard ...
        Standard SkuName = "Standard"
            )
    // PossibleSkuNameValues returns an array of possible values for the SkuName const type.
    func PossibleSkuNameValues() []SkuName {
        return []SkuName{Premium,Standard}
    }

