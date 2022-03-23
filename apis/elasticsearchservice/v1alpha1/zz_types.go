/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by ack-generate. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Hack to avoid import errors during build...
var (
	_ = &metav1.Time{}
)

// +kubebuilder:skipversion
type AccessPoliciesStatus struct {
	// Access policy rules for an Elasticsearch domain service endpoints. For more
	// information, see Configuring Access Policies (http://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/es-createupdatedomains.html#es-createdomain-configure-access-policies)
	// in the Amazon Elasticsearch Service Developer Guide. The maximum size of
	// a policy document is 100 KB.
	Options *string `json:"options,omitempty"`
}

// +kubebuilder:skipversion
type AdvancedOptionsStatus struct {
	// Exposes select native Elasticsearch configuration values from elasticsearch.yml.
	// Currently, the following advanced options are available:
	//
	//    * Option to allow references to indices in an HTTP request body. Must
	//    be false when configuring access to individual sub-resources. By default,
	//    the value is true. See Configuration Advanced Options for more information.
	//
	//    * Option to specify the percentage of heap space that is allocated to
	//    field data. By default, this setting is unbounded.
	//
	// For more information, see Configuring Advanced Options (http://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/es-createupdatedomains.html#es-createdomain-configure-advanced-options).
	Options map[string]*string `json:"options,omitempty"`
}

// +kubebuilder:skipversion
type AdvancedSecurityOptions struct {
	Enabled *bool `json:"enabled,omitempty"`

	InternalUserDatabaseEnabled *bool `json:"internalUserDatabaseEnabled,omitempty"`
	// Describes the SAML application configured for the domain.
	SAMLOptions *SAMLOptionsOutput `json:"sAMLOptions,omitempty"`
}

// +kubebuilder:skipversion
type AdvancedSecurityOptionsInput struct {
	Enabled *bool `json:"enabled,omitempty"`

	InternalUserDatabaseEnabled *bool `json:"internalUserDatabaseEnabled,omitempty"`
	// Credentials for the master user: username and password, ARN, or both.
	MasterUserOptions *MasterUserOptions `json:"masterUserOptions,omitempty"`
	// Specifies the SAML application configuration for the domain.
	SAMLOptions *SAMLOptionsInput `json:"sAMLOptions,omitempty"`
}

// +kubebuilder:skipversion
type AdvancedSecurityOptionsStatus struct {
	// Specifies the advanced security configuration: whether advanced security
	// is enabled, whether the internal database option is enabled.
	Options *AdvancedSecurityOptions `json:"options,omitempty"`
}

// +kubebuilder:skipversion
type AutoTuneMaintenanceSchedule struct {
	CronExpressionForRecurrence *string `json:"cronExpressionForRecurrence,omitempty"`
	// Specifies maintenance schedule duration: duration value and duration unit.
	// See the Developer Guide (https://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/auto-tune.html)
	// for more information.
	Duration *Duration `json:"duration,omitempty"`

	StartAt *metav1.Time `json:"startAt,omitempty"`
}

// +kubebuilder:skipversion
type AutoTuneOptions struct {
	// Specifies the Auto-Tune desired state. Valid values are ENABLED, DISABLED.
	DesiredState *string `json:"desiredState,omitempty"`

	MaintenanceSchedules []*AutoTuneMaintenanceSchedule `json:"maintenanceSchedules,omitempty"`
}

// +kubebuilder:skipversion
type AutoTuneOptionsInput struct {
	// Specifies the Auto-Tune desired state. Valid values are ENABLED, DISABLED.
	DesiredState *string `json:"desiredState,omitempty"`

	MaintenanceSchedules []*AutoTuneMaintenanceSchedule `json:"maintenanceSchedules,omitempty"`
}

// +kubebuilder:skipversion
type AutoTuneOptionsOutput struct {
	ErrorMessage *string `json:"errorMessage,omitempty"`
	// Specifies the Auto-Tune state for the Elasticsearch domain. For valid states
	// see the Developer Guide (https://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/auto-tune.html).
	State *string `json:"state,omitempty"`
}

// +kubebuilder:skipversion
type AutoTuneStatus struct {
	ErrorMessage *string `json:"errorMessage,omitempty"`

	PendingDeletion *bool `json:"pendingDeletion,omitempty"`
	// Specifies the Auto-Tune state for the Elasticsearch domain. For valid states
	// see the Developer Guide (https://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/auto-tune.html).
	State *string `json:"state,omitempty"`
}

// +kubebuilder:skipversion
type CognitoOptions struct {
	Enabled *bool `json:"enabled,omitempty"`

	IdentityPoolID *string `json:"identityPoolID,omitempty"`

	RoleARN *string `json:"roleARN,omitempty"`

	UserPoolID *string `json:"userPoolID,omitempty"`
}

// +kubebuilder:skipversion
type CognitoOptionsStatus struct {
	// Options to specify the Cognito user and identity pools for Kibana authentication.
	// For more information, see Amazon Cognito Authentication for Kibana (http://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/es-cognito-auth.html).
	Options *CognitoOptions `json:"options,omitempty"`
}

// +kubebuilder:skipversion
type ColdStorageOptions struct {
	Enabled *bool `json:"enabled,omitempty"`
}

// +kubebuilder:skipversion
type CompatibleVersionsMap struct {
	SourceVersion *string `json:"sourceVersion,omitempty"`
}

// +kubebuilder:skipversion
type DomainEndpointOptions struct {
	CustomEndpoint *string `json:"customEndpoint,omitempty"`
	// The Amazon Resource Name (ARN) of the Elasticsearch domain. See Identifiers
	// for IAM Entities (http://docs.aws.amazon.com/IAM/latest/UserGuide/index.html?Using_Identifiers.html)
	// in Using AWS Identity and Access Management for more information.
	CustomEndpointCertificateARN *string `json:"customEndpointCertificateARN,omitempty"`

	CustomEndpointEnabled *bool `json:"customEndpointEnabled,omitempty"`

	EnforceHTTPS *bool `json:"enforceHTTPS,omitempty"`

	TLSSecurityPolicy *string `json:"tlsSecurityPolicy,omitempty"`
}

// +kubebuilder:skipversion
type DomainEndpointOptionsStatus struct {
	// Options to configure endpoint for the Elasticsearch domain.
	Options *DomainEndpointOptions `json:"options,omitempty"`
}

// +kubebuilder:skipversion
type DomainInfo struct {
	// The name of an Elasticsearch domain. Domain names are unique across the domains
	// owned by an account within an AWS region. Domain names start with a letter
	// or number and can contain the following characters: a-z (lowercase), 0-9,
	// and - (hyphen).
	DomainName *string `json:"domainName,omitempty"`
}

// +kubebuilder:skipversion
type DomainInformation struct {
	// The name of an Elasticsearch domain. Domain names are unique across the domains
	// owned by an account within an AWS region. Domain names start with a letter
	// or number and can contain the following characters: a-z (lowercase), 0-9,
	// and - (hyphen).
	DomainName *string `json:"domainName,omitempty"`
}

// +kubebuilder:skipversion
type DomainPackageDetails struct {
	// The name of an Elasticsearch domain. Domain names are unique across the domains
	// owned by an account within an AWS region. Domain names start with a letter
	// or number and can contain the following characters: a-z (lowercase), 0-9,
	// and - (hyphen).
	DomainName *string `json:"domainName,omitempty"`
}

// +kubebuilder:skipversion
type Duration struct {
	// Specifies the unit of a maintenance schedule duration. Valid value is HOUR.
	// See the Developer Guide (https://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/auto-tune.html)
	// for more information.
	Unit *string `json:"unit,omitempty"`
	// Integer to specify the value of a maintenance schedule duration. See the
	// Developer Guide (https://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/auto-tune.html)
	// for more information.
	Value *int64 `json:"value,omitempty"`
}

// +kubebuilder:skipversion
type EBSOptions struct {
	EBSEnabled *bool `json:"ebsEnabled,omitempty"`

	IOPS *int64 `json:"iops,omitempty"`

	VolumeSize *int64 `json:"volumeSize,omitempty"`
	// The type of EBS volume, standard, gp2, or io1. See Configuring EBS-based
	// Storage (http://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/es-createupdatedomains.html#es-createdomain-configure-ebs)for
	// more information.
	VolumeType *string `json:"volumeType,omitempty"`
}

// +kubebuilder:skipversion
type EBSOptionsStatus struct {
	// Options to enable, disable, and specify the properties of EBS storage volumes.
	// For more information, see Configuring EBS-based Storage (http://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/es-createupdatedomains.html#es-createdomain-configure-ebs).
	Options *EBSOptions `json:"options,omitempty"`
}

// +kubebuilder:skipversion
type ElasticsearchClusterConfig struct {
	// Specifies the configuration for cold storage options such as enabled
	ColdStorageOptions *ColdStorageOptions `json:"coldStorageOptions,omitempty"`

	DedicatedMasterCount *int64 `json:"dedicatedMasterCount,omitempty"`

	DedicatedMasterEnabled *bool `json:"dedicatedMasterEnabled,omitempty"`

	DedicatedMasterType *string `json:"dedicatedMasterType,omitempty"`

	InstanceCount *int64 `json:"instanceCount,omitempty"`

	InstanceType *string `json:"instanceType,omitempty"`

	WarmCount *int64 `json:"warmCount,omitempty"`

	WarmEnabled *bool `json:"warmEnabled,omitempty"`

	WarmType *string `json:"warmType,omitempty"`
	// Specifies the zone awareness configuration for the domain cluster, such as
	// the number of availability zones.
	ZoneAwarenessConfig *ZoneAwarenessConfig `json:"zoneAwarenessConfig,omitempty"`

	ZoneAwarenessEnabled *bool `json:"zoneAwarenessEnabled,omitempty"`
}

// +kubebuilder:skipversion
type ElasticsearchClusterConfigStatus struct {
	// Specifies the configuration for the domain cluster, such as the type and
	// number of instances.
	Options *ElasticsearchClusterConfig `json:"options,omitempty"`
}

// +kubebuilder:skipversion
type ElasticsearchDomainStatus_SDK struct {
	// The Amazon Resource Name (ARN) of the Elasticsearch domain. See Identifiers
	// for IAM Entities (http://docs.aws.amazon.com/IAM/latest/UserGuide/index.html?Using_Identifiers.html)
	// in Using AWS Identity and Access Management for more information.
	ARN *string `json:"arn,omitempty"`
	// Access policy rules for an Elasticsearch domain service endpoints. For more
	// information, see Configuring Access Policies (http://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/es-createupdatedomains.html#es-createdomain-configure-access-policies)
	// in the Amazon Elasticsearch Service Developer Guide. The maximum size of
	// a policy document is 100 KB.
	AccessPolicies *string `json:"accessPolicies,omitempty"`
	// Exposes select native Elasticsearch configuration values from elasticsearch.yml.
	// Currently, the following advanced options are available:
	//
	//    * Option to allow references to indices in an HTTP request body. Must
	//    be false when configuring access to individual sub-resources. By default,
	//    the value is true. See Configuration Advanced Options for more information.
	//
	//    * Option to specify the percentage of heap space that is allocated to
	//    field data. By default, this setting is unbounded.
	//
	// For more information, see Configuring Advanced Options (http://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/es-createupdatedomains.html#es-createdomain-configure-advanced-options).
	AdvancedOptions map[string]*string `json:"advancedOptions,omitempty"`
	// Specifies the advanced security configuration: whether advanced security
	// is enabled, whether the internal database option is enabled.
	AdvancedSecurityOptions *AdvancedSecurityOptions `json:"advancedSecurityOptions,omitempty"`
	// Specifies the Auto-Tune options: the Auto-Tune desired state for the domain
	// and list of maintenance schedules.
	AutoTuneOptions *AutoTuneOptionsOutput `json:"autoTuneOptions,omitempty"`
	// Options to specify the Cognito user and identity pools for Kibana authentication.
	// For more information, see Amazon Cognito Authentication for Kibana (http://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/es-cognito-auth.html).
	CognitoOptions *CognitoOptions `json:"cognitoOptions,omitempty"`

	Created *bool `json:"created,omitempty"`

	Deleted *bool `json:"deleted,omitempty"`
	// Options to configure endpoint for the Elasticsearch domain.
	DomainEndpointOptions *DomainEndpointOptions `json:"domainEndpointOptions,omitempty"`
	// Unique identifier for an Elasticsearch domain.
	DomainID *string `json:"domainID,omitempty"`
	// The name of an Elasticsearch domain. Domain names are unique across the domains
	// owned by an account within an AWS region. Domain names start with a letter
	// or number and can contain the following characters: a-z (lowercase), 0-9,
	// and - (hyphen).
	DomainName *string `json:"domainName,omitempty"`
	// Options to enable, disable, and specify the properties of EBS storage volumes.
	// For more information, see Configuring EBS-based Storage (http://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/es-createupdatedomains.html#es-createdomain-configure-ebs).
	EBSOptions *EBSOptions `json:"ebsOptions,omitempty"`
	// Specifies the configuration for the domain cluster, such as the type and
	// number of instances.
	ElasticsearchClusterConfig *ElasticsearchClusterConfig `json:"elasticsearchClusterConfig,omitempty"`

	ElasticsearchVersion *string `json:"elasticsearchVersion,omitempty"`
	// Specifies the Encryption At Rest Options.
	EncryptionAtRestOptions *EncryptionAtRestOptions `json:"encryptionAtRestOptions,omitempty"`
	// The endpoint to which service requests are submitted. For example, search-imdb-movies-oopcnjfn6ugofer3zx5iadxxca.eu-west-1.es.amazonaws.com
	// or doc-imdb-movies-oopcnjfn6ugofer3zx5iadxxca.eu-west-1.es.amazonaws.com.
	Endpoint *string `json:"endpoint,omitempty"`

	Endpoints map[string]*string `json:"endpoints,omitempty"`

	LogPublishingOptions map[string]*LogPublishingOption `json:"logPublishingOptions,omitempty"`
	// Specifies the node-to-node encryption options.
	NodeToNodeEncryptionOptions *NodeToNodeEncryptionOptions `json:"nodeToNodeEncryptionOptions,omitempty"`

	Processing *bool `json:"processing,omitempty"`
	// The current options of an Elasticsearch domain service software options.
	ServiceSoftwareOptions *ServiceSoftwareOptions `json:"serviceSoftwareOptions,omitempty"`
	// Specifies the time, in UTC format, when the service takes a daily automated
	// snapshot of the specified Elasticsearch domain. Default value is 0 hours.
	SnapshotOptions *SnapshotOptions `json:"snapshotOptions,omitempty"`

	UpgradeProcessing *bool `json:"upgradeProcessing,omitempty"`
	// Options to specify the subnets and security groups for VPC endpoint. For
	// more information, see VPC Endpoints for Amazon Elasticsearch Service Domains
	// (http://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/es-vpc.html).
	VPCOptions *VPCDerivedInfo `json:"vpcOptions,omitempty"`
}

// +kubebuilder:skipversion
type ElasticsearchVersionStatus struct {
	Options *string `json:"options,omitempty"`
}

// +kubebuilder:skipversion
type EncryptionAtRestOptions struct {
	Enabled *bool `json:"enabled,omitempty"`

	KMSKeyID *string `json:"kmsKeyID,omitempty"`
}

// +kubebuilder:skipversion
type EncryptionAtRestOptionsStatus struct {
	// Specifies the Encryption At Rest Options.
	Options *EncryptionAtRestOptions `json:"options,omitempty"`
}

// +kubebuilder:skipversion
type LogPublishingOption struct {
	// ARN of the Cloudwatch log group to which log needs to be published.
	CloudWatchLogsLogGroupARN *string `json:"cloudWatchLogsLogGroupARN,omitempty"`

	Enabled *bool `json:"enabled,omitempty"`
}

// +kubebuilder:skipversion
type LogPublishingOptionsStatus struct {
	Options map[string]*LogPublishingOption `json:"options,omitempty"`
}

// +kubebuilder:skipversion
type MasterUserOptions struct {
	// The Amazon Resource Name (ARN) of the Elasticsearch domain. See Identifiers
	// for IAM Entities (http://docs.aws.amazon.com/IAM/latest/UserGuide/index.html?Using_Identifiers.html)
	// in Using AWS Identity and Access Management for more information.
	MasterUserARN *string `json:"masterUserARN,omitempty"`

	MasterUserName *string `json:"masterUserName,omitempty"`

	MasterUserPassword *string `json:"masterUserPassword,omitempty"`
}

// +kubebuilder:skipversion
type NodeToNodeEncryptionOptions struct {
	Enabled *bool `json:"enabled,omitempty"`
}

// +kubebuilder:skipversion
type NodeToNodeEncryptionOptionsStatus struct {
	// Specifies the node-to-node encryption options.
	Options *NodeToNodeEncryptionOptions `json:"options,omitempty"`
}

// +kubebuilder:skipversion
type OptionStatus struct {
	PendingDeletion *bool `json:"pendingDeletion,omitempty"`
}

// +kubebuilder:skipversion
type RecurringCharge struct {
	RecurringChargeFrequency *string `json:"recurringChargeFrequency,omitempty"`
}

// +kubebuilder:skipversion
type ReservedElasticsearchInstance struct {
	CurrencyCode *string `json:"currencyCode,omitempty"`

	ElasticsearchInstanceType *string `json:"elasticsearchInstanceType,omitempty"`

	ReservedElasticsearchInstanceOfferingID *string `json:"reservedElasticsearchInstanceOfferingID,omitempty"`

	State *string `json:"state,omitempty"`
}

// +kubebuilder:skipversion
type ReservedElasticsearchInstanceOffering struct {
	CurrencyCode *string `json:"currencyCode,omitempty"`

	ElasticsearchInstanceType *string `json:"elasticsearchInstanceType,omitempty"`
}

// +kubebuilder:skipversion
type SAMLIDp struct {
	EntityID *string `json:"entityID,omitempty"`

	MetadataContent *string `json:"metadataContent,omitempty"`
}

// +kubebuilder:skipversion
type SAMLOptionsInput struct {
	Enabled *bool `json:"enabled,omitempty"`
	// Specifies the SAML Identity Provider's information.
	IDp *SAMLIDp `json:"idp,omitempty"`

	MasterBackendRole *string `json:"masterBackendRole,omitempty"`

	MasterUserName *string `json:"masterUserName,omitempty"`

	RolesKey *string `json:"rolesKey,omitempty"`

	SessionTimeoutMinutes *int64 `json:"sessionTimeoutMinutes,omitempty"`

	SubjectKey *string `json:"subjectKey,omitempty"`
}

// +kubebuilder:skipversion
type SAMLOptionsOutput struct {
	Enabled *bool `json:"enabled,omitempty"`
	// Specifies the SAML Identity Provider's information.
	IDp *SAMLIDp `json:"idp,omitempty"`

	RolesKey *string `json:"rolesKey,omitempty"`

	SessionTimeoutMinutes *int64 `json:"sessionTimeoutMinutes,omitempty"`

	SubjectKey *string `json:"subjectKey,omitempty"`
}

// +kubebuilder:skipversion
type ServiceSoftwareOptions struct {
	AutomatedUpdateDate *metav1.Time `json:"automatedUpdateDate,omitempty"`

	Cancellable *bool `json:"cancellable,omitempty"`

	CurrentVersion *string `json:"currentVersion,omitempty"`

	Description *string `json:"description,omitempty"`

	NewVersion *string `json:"newVersion,omitempty"`

	OptionalDeployment *bool `json:"optionalDeployment,omitempty"`

	UpdateAvailable *bool `json:"updateAvailable,omitempty"`

	UpdateStatus *string `json:"updateStatus,omitempty"`
}

// +kubebuilder:skipversion
type SnapshotOptions struct {
	AutomatedSnapshotStartHour *int64 `json:"automatedSnapshotStartHour,omitempty"`
}

// +kubebuilder:skipversion
type SnapshotOptionsStatus struct {
	// Specifies the time, in UTC format, when the service takes a daily automated
	// snapshot of the specified Elasticsearch domain. Default value is 0 hours.
	Options *SnapshotOptions `json:"options,omitempty"`
}

// +kubebuilder:skipversion
type Tag struct {
	// A string of length from 1 to 128 characters that specifies the key for a
	// Tag. Tag keys must be unique for the Elasticsearch domain to which they are
	// attached.
	Key *string `json:"key,omitempty"`
	// A string of length from 0 to 256 characters that specifies the value for
	// a Tag. Tag values can be null and do not have to be unique in a tag set.
	Value *string `json:"value,omitempty"`
}

// +kubebuilder:skipversion
type VPCDerivedInfo struct {
	AvailabilityZones []*string `json:"availabilityZones,omitempty"`

	SecurityGroupIDs []*string `json:"securityGroupIDs,omitempty"`

	SubnetIDs []*string `json:"subnetIDs,omitempty"`

	VPCID *string `json:"vpcID,omitempty"`
}

// +kubebuilder:skipversion
type VPCDerivedInfoStatus struct {
	// Options to specify the subnets and security groups for VPC endpoint. For
	// more information, see VPC Endpoints for Amazon Elasticsearch Service Domains
	// (http://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/es-vpc.html).
	Options *VPCDerivedInfo `json:"options,omitempty"`
}

// +kubebuilder:skipversion
type VPCOptions struct {
	SecurityGroupIDs []*string `json:"securityGroupIDs,omitempty"`

	SubnetIDs []*string `json:"subnetIDs,omitempty"`
}

// +kubebuilder:skipversion
type ZoneAwarenessConfig struct {
	AvailabilityZoneCount *int64 `json:"availabilityZoneCount,omitempty"`
}
