package documentdb

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

// CompositePathSortOrder enumerates the values for composite path sort order.
type CompositePathSortOrder string

const (
	// Ascending ...
	Ascending CompositePathSortOrder = "Ascending"
	// Descending ...
	Descending CompositePathSortOrder = "Descending"
)

// PossibleCompositePathSortOrderValues returns an array of possible values for the CompositePathSortOrder const type.
func PossibleCompositePathSortOrderValues() []CompositePathSortOrder {
	return []CompositePathSortOrder{Ascending, Descending}
}

// ConflictResolutionMode enumerates the values for conflict resolution mode.
type ConflictResolutionMode string

const (
	// Custom ...
	Custom ConflictResolutionMode = "Custom"
	// LastWriterWins ...
	LastWriterWins ConflictResolutionMode = "LastWriterWins"
)

// PossibleConflictResolutionModeValues returns an array of possible values for the ConflictResolutionMode const type.
func PossibleConflictResolutionModeValues() []ConflictResolutionMode {
	return []ConflictResolutionMode{Custom, LastWriterWins}
}

// ConnectorOffer enumerates the values for connector offer.
type ConnectorOffer string

const (
	// Small ...
	Small ConnectorOffer = "Small"
)

// PossibleConnectorOfferValues returns an array of possible values for the ConnectorOffer const type.
func PossibleConnectorOfferValues() []ConnectorOffer {
	return []ConnectorOffer{Small}
}

// CreatedByType enumerates the values for created by type.
type CreatedByType string

const (
	// Application ...
	Application CreatedByType = "Application"
	// Key ...
	Key CreatedByType = "Key"
	// ManagedIdentity ...
	ManagedIdentity CreatedByType = "ManagedIdentity"
	// User ...
	User CreatedByType = "User"
)

// PossibleCreatedByTypeValues returns an array of possible values for the CreatedByType const type.
func PossibleCreatedByTypeValues() []CreatedByType {
	return []CreatedByType{Application, Key, ManagedIdentity, User}
}

// CreateMode enumerates the values for create mode.
type CreateMode string

const (
	// Default ...
	Default CreateMode = "Default"
	// Restore ...
	Restore CreateMode = "Restore"
)

// PossibleCreateModeValues returns an array of possible values for the CreateMode const type.
func PossibleCreateModeValues() []CreateMode {
	return []CreateMode{Default, Restore}
}

// CreateModeBasicDatabaseAccountCreateUpdateProperties enumerates the values for create mode basic database
// account create update properties.
type CreateModeBasicDatabaseAccountCreateUpdateProperties string

const (
	// CreateModeDatabaseAccountCreateUpdateProperties ...
	CreateModeDatabaseAccountCreateUpdateProperties CreateModeBasicDatabaseAccountCreateUpdateProperties = "DatabaseAccountCreateUpdateProperties"
	// CreateModeDefault ...
	CreateModeDefault CreateModeBasicDatabaseAccountCreateUpdateProperties = "Default"
	// CreateModeRestore ...
	CreateModeRestore CreateModeBasicDatabaseAccountCreateUpdateProperties = "Restore"
)

// PossibleCreateModeBasicDatabaseAccountCreateUpdatePropertiesValues returns an array of possible values for the CreateModeBasicDatabaseAccountCreateUpdateProperties const type.
func PossibleCreateModeBasicDatabaseAccountCreateUpdatePropertiesValues() []CreateModeBasicDatabaseAccountCreateUpdateProperties {
	return []CreateModeBasicDatabaseAccountCreateUpdateProperties{CreateModeDatabaseAccountCreateUpdateProperties, CreateModeDefault, CreateModeRestore}
}

// DatabaseAccountKind enumerates the values for database account kind.
type DatabaseAccountKind string

const (
	// GlobalDocumentDB ...
	GlobalDocumentDB DatabaseAccountKind = "GlobalDocumentDB"
	// MongoDB ...
	MongoDB DatabaseAccountKind = "MongoDB"
	// Parse ...
	Parse DatabaseAccountKind = "Parse"
)

// PossibleDatabaseAccountKindValues returns an array of possible values for the DatabaseAccountKind const type.
func PossibleDatabaseAccountKindValues() []DatabaseAccountKind {
	return []DatabaseAccountKind{GlobalDocumentDB, MongoDB, Parse}
}

// DatabaseAccountOfferType enumerates the values for database account offer type.
type DatabaseAccountOfferType string

const (
	// Standard ...
	Standard DatabaseAccountOfferType = "Standard"
)

// PossibleDatabaseAccountOfferTypeValues returns an array of possible values for the DatabaseAccountOfferType const type.
func PossibleDatabaseAccountOfferTypeValues() []DatabaseAccountOfferType {
	return []DatabaseAccountOfferType{Standard}
}

// DataType enumerates the values for data type.
type DataType string

const (
	// LineString ...
	LineString DataType = "LineString"
	// MultiPolygon ...
	MultiPolygon DataType = "MultiPolygon"
	// Number ...
	Number DataType = "Number"
	// Point ...
	Point DataType = "Point"
	// Polygon ...
	Polygon DataType = "Polygon"
	// String ...
	String DataType = "String"
)

// PossibleDataTypeValues returns an array of possible values for the DataType const type.
func PossibleDataTypeValues() []DataType {
	return []DataType{LineString, MultiPolygon, Number, Point, Polygon, String}
}

// DefaultConsistencyLevel enumerates the values for default consistency level.
type DefaultConsistencyLevel string

const (
	// BoundedStaleness ...
	BoundedStaleness DefaultConsistencyLevel = "BoundedStaleness"
	// ConsistentPrefix ...
	ConsistentPrefix DefaultConsistencyLevel = "ConsistentPrefix"
	// Eventual ...
	Eventual DefaultConsistencyLevel = "Eventual"
	// Session ...
	Session DefaultConsistencyLevel = "Session"
	// Strong ...
	Strong DefaultConsistencyLevel = "Strong"
)

// PossibleDefaultConsistencyLevelValues returns an array of possible values for the DefaultConsistencyLevel const type.
func PossibleDefaultConsistencyLevelValues() []DefaultConsistencyLevel {
	return []DefaultConsistencyLevel{BoundedStaleness, ConsistentPrefix, Eventual, Session, Strong}
}

// IndexingMode enumerates the values for indexing mode.
type IndexingMode string

const (
	// Consistent ...
	Consistent IndexingMode = "Consistent"
	// Lazy ...
	Lazy IndexingMode = "Lazy"
	// None ...
	None IndexingMode = "None"
)

// PossibleIndexingModeValues returns an array of possible values for the IndexingMode const type.
func PossibleIndexingModeValues() []IndexingMode {
	return []IndexingMode{Consistent, Lazy, None}
}

// IndexKind enumerates the values for index kind.
type IndexKind string

const (
	// Hash ...
	Hash IndexKind = "Hash"
	// Range ...
	Range IndexKind = "Range"
	// Spatial ...
	Spatial IndexKind = "Spatial"
)

// PossibleIndexKindValues returns an array of possible values for the IndexKind const type.
func PossibleIndexKindValues() []IndexKind {
	return []IndexKind{Hash, Range, Spatial}
}

// KeyKind enumerates the values for key kind.
type KeyKind string

const (
	// Primary ...
	Primary KeyKind = "primary"
	// PrimaryReadonly ...
	PrimaryReadonly KeyKind = "primaryReadonly"
	// Secondary ...
	Secondary KeyKind = "secondary"
	// SecondaryReadonly ...
	SecondaryReadonly KeyKind = "secondaryReadonly"
)

// PossibleKeyKindValues returns an array of possible values for the KeyKind const type.
func PossibleKeyKindValues() []KeyKind {
	return []KeyKind{Primary, PrimaryReadonly, Secondary, SecondaryReadonly}
}

// PartitionKind enumerates the values for partition kind.
type PartitionKind string

const (
	// PartitionKindHash ...
	PartitionKindHash PartitionKind = "Hash"
	// PartitionKindRange ...
	PartitionKindRange PartitionKind = "Range"
)

// PossiblePartitionKindValues returns an array of possible values for the PartitionKind const type.
func PossiblePartitionKindValues() []PartitionKind {
	return []PartitionKind{PartitionKindHash, PartitionKindRange}
}

// PrimaryAggregationType enumerates the values for primary aggregation type.
type PrimaryAggregationType string

const (
	// PrimaryAggregationTypeAverage ...
	PrimaryAggregationTypeAverage PrimaryAggregationType = "Average"
	// PrimaryAggregationTypeLast ...
	PrimaryAggregationTypeLast PrimaryAggregationType = "Last"
	// PrimaryAggregationTypeMaximum ...
	PrimaryAggregationTypeMaximum PrimaryAggregationType = "Maximum"
	// PrimaryAggregationTypeMinimum ...
	PrimaryAggregationTypeMinimum PrimaryAggregationType = "Minimum"
	// PrimaryAggregationTypeNone ...
	PrimaryAggregationTypeNone PrimaryAggregationType = "None"
	// PrimaryAggregationTypeTotal ...
	PrimaryAggregationTypeTotal PrimaryAggregationType = "Total"
)

// PossiblePrimaryAggregationTypeValues returns an array of possible values for the PrimaryAggregationType const type.
func PossiblePrimaryAggregationTypeValues() []PrimaryAggregationType {
	return []PrimaryAggregationType{PrimaryAggregationTypeAverage, PrimaryAggregationTypeLast, PrimaryAggregationTypeMaximum, PrimaryAggregationTypeMinimum, PrimaryAggregationTypeNone, PrimaryAggregationTypeTotal}
}

// PublicNetworkAccess enumerates the values for public network access.
type PublicNetworkAccess string

const (
	// Disabled ...
	Disabled PublicNetworkAccess = "Disabled"
	// Enabled ...
	Enabled PublicNetworkAccess = "Enabled"
)

// PossiblePublicNetworkAccessValues returns an array of possible values for the PublicNetworkAccess const type.
func PossiblePublicNetworkAccessValues() []PublicNetworkAccess {
	return []PublicNetworkAccess{Disabled, Enabled}
}

// ResourceIdentityType enumerates the values for resource identity type.
type ResourceIdentityType string

const (
	// ResourceIdentityTypeNone ...
	ResourceIdentityTypeNone ResourceIdentityType = "None"
	// ResourceIdentityTypeSystemAssigned ...
	ResourceIdentityTypeSystemAssigned ResourceIdentityType = "SystemAssigned"
	// ResourceIdentityTypeSystemAssignedUserAssigned ...
	ResourceIdentityTypeSystemAssignedUserAssigned ResourceIdentityType = "SystemAssigned,UserAssigned"
	// ResourceIdentityTypeUserAssigned ...
	ResourceIdentityTypeUserAssigned ResourceIdentityType = "UserAssigned"
)

// PossibleResourceIdentityTypeValues returns an array of possible values for the ResourceIdentityType const type.
func PossibleResourceIdentityTypeValues() []ResourceIdentityType {
	return []ResourceIdentityType{ResourceIdentityTypeNone, ResourceIdentityTypeSystemAssigned, ResourceIdentityTypeSystemAssignedUserAssigned, ResourceIdentityTypeUserAssigned}
}

// RestoreMode enumerates the values for restore mode.
type RestoreMode string

const (
	// PointInTime ...
	PointInTime RestoreMode = "PointInTime"
)

// PossibleRestoreModeValues returns an array of possible values for the RestoreMode const type.
func PossibleRestoreModeValues() []RestoreMode {
	return []RestoreMode{PointInTime}
}

// RoleDefinitionType enumerates the values for role definition type.
type RoleDefinitionType string

const (
	// BuiltInRole ...
	BuiltInRole RoleDefinitionType = "BuiltInRole"
	// CustomRole ...
	CustomRole RoleDefinitionType = "CustomRole"
)

// PossibleRoleDefinitionTypeValues returns an array of possible values for the RoleDefinitionType const type.
func PossibleRoleDefinitionTypeValues() []RoleDefinitionType {
	return []RoleDefinitionType{BuiltInRole, CustomRole}
}

// ServerVersion enumerates the values for server version.
type ServerVersion string

const (
	// ThreeFullStopSix ...
	ThreeFullStopSix ServerVersion = "3.6"
	// ThreeFullStopTwo ...
	ThreeFullStopTwo ServerVersion = "3.2"
)

// PossibleServerVersionValues returns an array of possible values for the ServerVersion const type.
func PossibleServerVersionValues() []ServerVersion {
	return []ServerVersion{ThreeFullStopSix, ThreeFullStopTwo}
}

// SpatialType enumerates the values for spatial type.
type SpatialType string

const (
	// SpatialTypeLineString ...
	SpatialTypeLineString SpatialType = "LineString"
	// SpatialTypeMultiPolygon ...
	SpatialTypeMultiPolygon SpatialType = "MultiPolygon"
	// SpatialTypePoint ...
	SpatialTypePoint SpatialType = "Point"
	// SpatialTypePolygon ...
	SpatialTypePolygon SpatialType = "Polygon"
)

// PossibleSpatialTypeValues returns an array of possible values for the SpatialType const type.
func PossibleSpatialTypeValues() []SpatialType {
	return []SpatialType{SpatialTypeLineString, SpatialTypeMultiPolygon, SpatialTypePoint, SpatialTypePolygon}
}

// TriggerOperation enumerates the values for trigger operation.
type TriggerOperation string

const (
	// All ...
	All TriggerOperation = "All"
	// Create ...
	Create TriggerOperation = "Create"
	// Delete ...
	Delete TriggerOperation = "Delete"
	// Replace ...
	Replace TriggerOperation = "Replace"
	// Update ...
	Update TriggerOperation = "Update"
)

// PossibleTriggerOperationValues returns an array of possible values for the TriggerOperation const type.
func PossibleTriggerOperationValues() []TriggerOperation {
	return []TriggerOperation{All, Create, Delete, Replace, Update}
}

// TriggerType enumerates the values for trigger type.
type TriggerType string

const (
	// Post ...
	Post TriggerType = "Post"
	// Pre ...
	Pre TriggerType = "Pre"
)

// PossibleTriggerTypeValues returns an array of possible values for the TriggerType const type.
func PossibleTriggerTypeValues() []TriggerType {
	return []TriggerType{Post, Pre}
}

// Type enumerates the values for type.
type Type string

const (
	// TypeBackupPolicy ...
	TypeBackupPolicy Type = "BackupPolicy"
	// TypeContinuous ...
	TypeContinuous Type = "Continuous"
	// TypePeriodic ...
	TypePeriodic Type = "Periodic"
)

// PossibleTypeValues returns an array of possible values for the Type const type.
func PossibleTypeValues() []Type {
	return []Type{TypeBackupPolicy, TypeContinuous, TypePeriodic}
}

// UnitType enumerates the values for unit type.
type UnitType string

const (
	// Bytes ...
	Bytes UnitType = "Bytes"
	// BytesPerSecond ...
	BytesPerSecond UnitType = "BytesPerSecond"
	// Count ...
	Count UnitType = "Count"
	// CountPerSecond ...
	CountPerSecond UnitType = "CountPerSecond"
	// Milliseconds ...
	Milliseconds UnitType = "Milliseconds"
	// Percent ...
	Percent UnitType = "Percent"
	// Seconds ...
	Seconds UnitType = "Seconds"
)

// PossibleUnitTypeValues returns an array of possible values for the UnitType const type.
func PossibleUnitTypeValues() []UnitType {
	return []UnitType{Bytes, BytesPerSecond, Count, CountPerSecond, Milliseconds, Percent, Seconds}
}
