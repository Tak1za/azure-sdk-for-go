package iotspaces

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

// NameUnavailabilityReason enumerates the values for name unavailability reason.
type NameUnavailabilityReason string

const (
	// AlreadyExists ...
	AlreadyExists NameUnavailabilityReason = "AlreadyExists"
	// Invalid ...
	Invalid NameUnavailabilityReason = "Invalid"
)

// PossibleNameUnavailabilityReasonValues returns an array of possible values for the NameUnavailabilityReason const type.
func PossibleNameUnavailabilityReasonValues() []NameUnavailabilityReason {
	return []NameUnavailabilityReason{AlreadyExists, Invalid}
}

// ProvisioningState enumerates the values for provisioning state.
type ProvisioningState string

const (
	// Canceled ...
	Canceled ProvisioningState = "Canceled"
	// Deleting ...
	Deleting ProvisioningState = "Deleting"
	// Failed ...
	Failed ProvisioningState = "Failed"
	// Provisioning ...
	Provisioning ProvisioningState = "Provisioning"
	// Succeeded ...
	Succeeded ProvisioningState = "Succeeded"
)

// PossibleProvisioningStateValues returns an array of possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{Canceled, Deleting, Failed, Provisioning, Succeeded}
}

// Sku enumerates the values for sku.
type Sku string

const (
	// F1 ...
	F1 Sku = "F1"
	// S1 ...
	S1 Sku = "S1"
	// S2 ...
	S2 Sku = "S2"
	// S3 ...
	S3 Sku = "S3"
)

// PossibleSkuValues returns an array of possible values for the Sku const type.
func PossibleSkuValues() []Sku {
	return []Sku{F1, S1, S2, S3}
}