package subscription

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

// OfferType enumerates the values for offer type.
type OfferType string

const (
	// MSAZR0017P ...
	MSAZR0017P OfferType = "MS-AZR-0017P"
	// MSAZR0148P ...
	MSAZR0148P OfferType = "MS-AZR-0148P"
)

// PossibleOfferTypeValues returns an array of possible values for the OfferType const type.
func PossibleOfferTypeValues() []OfferType {
	return []OfferType{MSAZR0017P, MSAZR0148P}
}

// SpendingLimit enumerates the values for spending limit.
type SpendingLimit string

const (
	// CurrentPeriodOff ...
	CurrentPeriodOff SpendingLimit = "CurrentPeriodOff"
	// Off ...
	Off SpendingLimit = "Off"
	// On ...
	On SpendingLimit = "On"
)

// PossibleSpendingLimitValues returns an array of possible values for the SpendingLimit const type.
func PossibleSpendingLimitValues() []SpendingLimit {
	return []SpendingLimit{CurrentPeriodOff, Off, On}
}

// State enumerates the values for state.
type State string

const (
	// Deleted ...
	Deleted State = "Deleted"
	// Disabled ...
	Disabled State = "Disabled"
	// Enabled ...
	Enabled State = "Enabled"
	// PastDue ...
	PastDue State = "PastDue"
	// Warned ...
	Warned State = "Warned"
)

// PossibleStateValues returns an array of possible values for the State const type.
func PossibleStateValues() []State {
	return []State{Deleted, Disabled, Enabled, PastDue, Warned}
}
