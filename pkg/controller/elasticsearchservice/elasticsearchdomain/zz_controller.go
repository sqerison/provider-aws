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

package elasticsearchdomain

import (
	"context"

	svcapi "github.com/aws/aws-sdk-go/service/elasticsearchservice"
	svcsdk "github.com/aws/aws-sdk-go/service/elasticsearchservice"
	svcsdkapi "github.com/aws/aws-sdk-go/service/elasticsearchservice/elasticsearchserviceiface"
	"github.com/google/go-cmp/cmp"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	"github.com/crossplane/crossplane-runtime/pkg/meta"
	"github.com/crossplane/crossplane-runtime/pkg/reconciler/managed"
	cpresource "github.com/crossplane/crossplane-runtime/pkg/resource"

	svcapitypes "github.com/crossplane/provider-aws/apis/elasticsearchservice/v1alpha1"
	awsclient "github.com/crossplane/provider-aws/pkg/clients"
)

const (
	errUnexpectedObject = "managed resource is not an ElasticsearchDomain resource"

	errCreateSession = "cannot create a new session"
	errCreate        = "cannot create ElasticsearchDomain in AWS"
	errUpdate        = "cannot update ElasticsearchDomain in AWS"
	errDescribe      = "failed to describe ElasticsearchDomain"
	errDelete        = "failed to delete ElasticsearchDomain"
)

type connector struct {
	kube client.Client
	opts []option
}

func (c *connector) Connect(ctx context.Context, mg cpresource.Managed) (managed.ExternalClient, error) {
	cr, ok := mg.(*svcapitypes.ElasticsearchDomain)
	if !ok {
		return nil, errors.New(errUnexpectedObject)
	}
	sess, err := awsclient.GetConfigV1(ctx, c.kube, mg, cr.Spec.ForProvider.Region)
	if err != nil {
		return nil, errors.Wrap(err, errCreateSession)
	}
	return newExternal(c.kube, svcapi.New(sess), c.opts), nil
}

func (e *external) Observe(ctx context.Context, mg cpresource.Managed) (managed.ExternalObservation, error) {
	cr, ok := mg.(*svcapitypes.ElasticsearchDomain)
	if !ok {
		return managed.ExternalObservation{}, errors.New(errUnexpectedObject)
	}
	if meta.GetExternalName(cr) == "" {
		return managed.ExternalObservation{
			ResourceExists: false,
		}, nil
	}
	input := GenerateDescribeElasticsearchDomainInput(cr)
	if err := e.preObserve(ctx, cr, input); err != nil {
		return managed.ExternalObservation{}, errors.Wrap(err, "pre-observe failed")
	}
	resp, err := e.client.DescribeElasticsearchDomainWithContext(ctx, input)
	if err != nil {
		return managed.ExternalObservation{ResourceExists: false}, awsclient.Wrap(cpresource.Ignore(IsNotFound, err), errDescribe)
	}
	currentSpec := cr.Spec.ForProvider.DeepCopy()
	if err := e.lateInitialize(&cr.Spec.ForProvider, resp); err != nil {
		return managed.ExternalObservation{}, errors.Wrap(err, "late-init failed")
	}
	GenerateElasticsearchDomain(resp).Status.AtProvider.DeepCopyInto(&cr.Status.AtProvider)

	upToDate, err := e.isUpToDate(cr, resp)
	if err != nil {
		return managed.ExternalObservation{}, errors.Wrap(err, "isUpToDate check failed")
	}
	return e.postObserve(ctx, cr, resp, managed.ExternalObservation{
		ResourceExists:          true,
		ResourceUpToDate:        upToDate,
		ResourceLateInitialized: !cmp.Equal(&cr.Spec.ForProvider, currentSpec),
	}, nil)
}

func (e *external) Create(ctx context.Context, mg cpresource.Managed) (managed.ExternalCreation, error) {
	cr, ok := mg.(*svcapitypes.ElasticsearchDomain)
	if !ok {
		return managed.ExternalCreation{}, errors.New(errUnexpectedObject)
	}
	cr.Status.SetConditions(xpv1.Creating())
	input := GenerateCreateElasticsearchDomainInput(cr)
	if err := e.preCreate(ctx, cr, input); err != nil {
		return managed.ExternalCreation{}, errors.Wrap(err, "pre-create failed")
	}
	resp, err := e.client.CreateElasticsearchDomainWithContext(ctx, input)
	if err != nil {
		return managed.ExternalCreation{}, awsclient.Wrap(err, errCreate)
	}

	if resp.DomainStatus.ARN != nil {
		cr.Status.AtProvider.ARN = resp.DomainStatus.ARN
	} else {
		cr.Status.AtProvider.ARN = nil
	}
	if resp.DomainStatus.AccessPolicies != nil {
		cr.Spec.ForProvider.AccessPolicies = resp.DomainStatus.AccessPolicies
	} else {
		cr.Spec.ForProvider.AccessPolicies = nil
	}
	if resp.DomainStatus.AdvancedOptions != nil {
		f2 := map[string]*string{}
		for f2key, f2valiter := range resp.DomainStatus.AdvancedOptions {
			var f2val string
			f2val = *f2valiter
			f2[f2key] = &f2val
		}
		cr.Spec.ForProvider.AdvancedOptions = f2
	} else {
		cr.Spec.ForProvider.AdvancedOptions = nil
	}
	if resp.DomainStatus.AdvancedSecurityOptions != nil {
		f3 := &svcapitypes.AdvancedSecurityOptionsInput{}
		if resp.DomainStatus.AdvancedSecurityOptions.Enabled != nil {
			f3.Enabled = resp.DomainStatus.AdvancedSecurityOptions.Enabled
		}
		if resp.DomainStatus.AdvancedSecurityOptions.InternalUserDatabaseEnabled != nil {
			f3.InternalUserDatabaseEnabled = resp.DomainStatus.AdvancedSecurityOptions.InternalUserDatabaseEnabled
		}
		if resp.DomainStatus.AdvancedSecurityOptions.SAMLOptions != nil {
			f3f2 := &svcapitypes.SAMLOptionsInput{}
			if resp.DomainStatus.AdvancedSecurityOptions.SAMLOptions.Enabled != nil {
				f3f2.Enabled = resp.DomainStatus.AdvancedSecurityOptions.SAMLOptions.Enabled
			}
			if resp.DomainStatus.AdvancedSecurityOptions.SAMLOptions.Idp != nil {
				f3f2f1 := &svcapitypes.SAMLIDp{}
				if resp.DomainStatus.AdvancedSecurityOptions.SAMLOptions.Idp.EntityId != nil {
					f3f2f1.EntityID = resp.DomainStatus.AdvancedSecurityOptions.SAMLOptions.Idp.EntityId
				}
				if resp.DomainStatus.AdvancedSecurityOptions.SAMLOptions.Idp.MetadataContent != nil {
					f3f2f1.MetadataContent = resp.DomainStatus.AdvancedSecurityOptions.SAMLOptions.Idp.MetadataContent
				}
				f3f2.IDp = f3f2f1
			}
			if resp.DomainStatus.AdvancedSecurityOptions.SAMLOptions.RolesKey != nil {
				f3f2.RolesKey = resp.DomainStatus.AdvancedSecurityOptions.SAMLOptions.RolesKey
			}
			if resp.DomainStatus.AdvancedSecurityOptions.SAMLOptions.SessionTimeoutMinutes != nil {
				f3f2.SessionTimeoutMinutes = resp.DomainStatus.AdvancedSecurityOptions.SAMLOptions.SessionTimeoutMinutes
			}
			if resp.DomainStatus.AdvancedSecurityOptions.SAMLOptions.SubjectKey != nil {
				f3f2.SubjectKey = resp.DomainStatus.AdvancedSecurityOptions.SAMLOptions.SubjectKey
			}
			f3.SAMLOptions = f3f2
		}
		cr.Spec.ForProvider.AdvancedSecurityOptions = f3
	} else {
		cr.Spec.ForProvider.AdvancedSecurityOptions = nil
	}
	if resp.DomainStatus.AutoTuneOptions != nil {
		f4 := &svcapitypes.AutoTuneOptionsInput{}
		cr.Spec.ForProvider.AutoTuneOptions = f4
	} else {
		cr.Spec.ForProvider.AutoTuneOptions = nil
	}
	if resp.DomainStatus.CognitoOptions != nil {
		f5 := &svcapitypes.CognitoOptions{}
		if resp.DomainStatus.CognitoOptions.Enabled != nil {
			f5.Enabled = resp.DomainStatus.CognitoOptions.Enabled
		}
		if resp.DomainStatus.CognitoOptions.IdentityPoolId != nil {
			f5.IdentityPoolID = resp.DomainStatus.CognitoOptions.IdentityPoolId
		}
		if resp.DomainStatus.CognitoOptions.RoleArn != nil {
			f5.RoleARN = resp.DomainStatus.CognitoOptions.RoleArn
		}
		if resp.DomainStatus.CognitoOptions.UserPoolId != nil {
			f5.UserPoolID = resp.DomainStatus.CognitoOptions.UserPoolId
		}
		cr.Spec.ForProvider.CognitoOptions = f5
	} else {
		cr.Spec.ForProvider.CognitoOptions = nil
	}
	if resp.DomainStatus.Created != nil {
		cr.Status.AtProvider.Created = resp.DomainStatus.Created
	} else {
		cr.Status.AtProvider.Created = nil
	}
	if resp.DomainStatus.Deleted != nil {
		cr.Status.AtProvider.Deleted = resp.DomainStatus.Deleted
	} else {
		cr.Status.AtProvider.Deleted = nil
	}
	if resp.DomainStatus.DomainEndpointOptions != nil {
		f8 := &svcapitypes.DomainEndpointOptions{}
		if resp.DomainStatus.DomainEndpointOptions.CustomEndpoint != nil {
			f8.CustomEndpoint = resp.DomainStatus.DomainEndpointOptions.CustomEndpoint
		}
		if resp.DomainStatus.DomainEndpointOptions.CustomEndpointCertificateArn != nil {
			f8.CustomEndpointCertificateARN = resp.DomainStatus.DomainEndpointOptions.CustomEndpointCertificateArn
		}
		if resp.DomainStatus.DomainEndpointOptions.CustomEndpointEnabled != nil {
			f8.CustomEndpointEnabled = resp.DomainStatus.DomainEndpointOptions.CustomEndpointEnabled
		}
		if resp.DomainStatus.DomainEndpointOptions.EnforceHTTPS != nil {
			f8.EnforceHTTPS = resp.DomainStatus.DomainEndpointOptions.EnforceHTTPS
		}
		if resp.DomainStatus.DomainEndpointOptions.TLSSecurityPolicy != nil {
			f8.TLSSecurityPolicy = resp.DomainStatus.DomainEndpointOptions.TLSSecurityPolicy
		}
		cr.Spec.ForProvider.DomainEndpointOptions = f8
	} else {
		cr.Spec.ForProvider.DomainEndpointOptions = nil
	}
	if resp.DomainStatus.DomainId != nil {
		cr.Status.AtProvider.DomainID = resp.DomainStatus.DomainId
	} else {
		cr.Status.AtProvider.DomainID = nil
	}
	if resp.DomainStatus.DomainName != nil {
		cr.Spec.ForProvider.DomainName = resp.DomainStatus.DomainName
	} else {
		cr.Spec.ForProvider.DomainName = nil
	}
	if resp.DomainStatus.EBSOptions != nil {
		f11 := &svcapitypes.EBSOptions{}
		if resp.DomainStatus.EBSOptions.EBSEnabled != nil {
			f11.EBSEnabled = resp.DomainStatus.EBSOptions.EBSEnabled
		}
		if resp.DomainStatus.EBSOptions.Iops != nil {
			f11.IOPS = resp.DomainStatus.EBSOptions.Iops
		}
		if resp.DomainStatus.EBSOptions.VolumeSize != nil {
			f11.VolumeSize = resp.DomainStatus.EBSOptions.VolumeSize
		}
		if resp.DomainStatus.EBSOptions.VolumeType != nil {
			f11.VolumeType = resp.DomainStatus.EBSOptions.VolumeType
		}
		cr.Spec.ForProvider.EBSOptions = f11
	} else {
		cr.Spec.ForProvider.EBSOptions = nil
	}
	if resp.DomainStatus.ElasticsearchClusterConfig != nil {
		f12 := &svcapitypes.ElasticsearchClusterConfig{}
		if resp.DomainStatus.ElasticsearchClusterConfig.ColdStorageOptions != nil {
			f12f0 := &svcapitypes.ColdStorageOptions{}
			if resp.DomainStatus.ElasticsearchClusterConfig.ColdStorageOptions.Enabled != nil {
				f12f0.Enabled = resp.DomainStatus.ElasticsearchClusterConfig.ColdStorageOptions.Enabled
			}
			f12.ColdStorageOptions = f12f0
		}
		if resp.DomainStatus.ElasticsearchClusterConfig.DedicatedMasterCount != nil {
			f12.DedicatedMasterCount = resp.DomainStatus.ElasticsearchClusterConfig.DedicatedMasterCount
		}
		if resp.DomainStatus.ElasticsearchClusterConfig.DedicatedMasterEnabled != nil {
			f12.DedicatedMasterEnabled = resp.DomainStatus.ElasticsearchClusterConfig.DedicatedMasterEnabled
		}
		if resp.DomainStatus.ElasticsearchClusterConfig.DedicatedMasterType != nil {
			f12.DedicatedMasterType = resp.DomainStatus.ElasticsearchClusterConfig.DedicatedMasterType
		}
		if resp.DomainStatus.ElasticsearchClusterConfig.InstanceCount != nil {
			f12.InstanceCount = resp.DomainStatus.ElasticsearchClusterConfig.InstanceCount
		}
		if resp.DomainStatus.ElasticsearchClusterConfig.InstanceType != nil {
			f12.InstanceType = resp.DomainStatus.ElasticsearchClusterConfig.InstanceType
		}
		if resp.DomainStatus.ElasticsearchClusterConfig.WarmCount != nil {
			f12.WarmCount = resp.DomainStatus.ElasticsearchClusterConfig.WarmCount
		}
		if resp.DomainStatus.ElasticsearchClusterConfig.WarmEnabled != nil {
			f12.WarmEnabled = resp.DomainStatus.ElasticsearchClusterConfig.WarmEnabled
		}
		if resp.DomainStatus.ElasticsearchClusterConfig.WarmType != nil {
			f12.WarmType = resp.DomainStatus.ElasticsearchClusterConfig.WarmType
		}
		if resp.DomainStatus.ElasticsearchClusterConfig.ZoneAwarenessConfig != nil {
			f12f9 := &svcapitypes.ZoneAwarenessConfig{}
			if resp.DomainStatus.ElasticsearchClusterConfig.ZoneAwarenessConfig.AvailabilityZoneCount != nil {
				f12f9.AvailabilityZoneCount = resp.DomainStatus.ElasticsearchClusterConfig.ZoneAwarenessConfig.AvailabilityZoneCount
			}
			f12.ZoneAwarenessConfig = f12f9
		}
		if resp.DomainStatus.ElasticsearchClusterConfig.ZoneAwarenessEnabled != nil {
			f12.ZoneAwarenessEnabled = resp.DomainStatus.ElasticsearchClusterConfig.ZoneAwarenessEnabled
		}
		cr.Spec.ForProvider.ElasticsearchClusterConfig = f12
	} else {
		cr.Spec.ForProvider.ElasticsearchClusterConfig = nil
	}
	if resp.DomainStatus.ElasticsearchVersion != nil {
		cr.Spec.ForProvider.ElasticsearchVersion = resp.DomainStatus.ElasticsearchVersion
	} else {
		cr.Spec.ForProvider.ElasticsearchVersion = nil
	}
	if resp.DomainStatus.EncryptionAtRestOptions != nil {
		f14 := &svcapitypes.EncryptionAtRestOptions{}
		if resp.DomainStatus.EncryptionAtRestOptions.Enabled != nil {
			f14.Enabled = resp.DomainStatus.EncryptionAtRestOptions.Enabled
		}
		if resp.DomainStatus.EncryptionAtRestOptions.KmsKeyId != nil {
			f14.KMSKeyID = resp.DomainStatus.EncryptionAtRestOptions.KmsKeyId
		}
		cr.Spec.ForProvider.EncryptionAtRestOptions = f14
	} else {
		cr.Spec.ForProvider.EncryptionAtRestOptions = nil
	}
	if resp.DomainStatus.Endpoint != nil {
		cr.Status.AtProvider.Endpoint = resp.DomainStatus.Endpoint
	} else {
		cr.Status.AtProvider.Endpoint = nil
	}
	if resp.DomainStatus.Endpoints != nil {
		f16 := map[string]*string{}
		for f16key, f16valiter := range resp.DomainStatus.Endpoints {
			var f16val string
			f16val = *f16valiter
			f16[f16key] = &f16val
		}
		cr.Status.AtProvider.Endpoints = f16
	} else {
		cr.Status.AtProvider.Endpoints = nil
	}
	if resp.DomainStatus.LogPublishingOptions != nil {
		f17 := map[string]*svcapitypes.LogPublishingOption{}
		for f17key, f17valiter := range resp.DomainStatus.LogPublishingOptions {
			f17val := &svcapitypes.LogPublishingOption{}
			if f17valiter.CloudWatchLogsLogGroupArn != nil {
				f17val.CloudWatchLogsLogGroupARN = f17valiter.CloudWatchLogsLogGroupArn
			}
			if f17valiter.Enabled != nil {
				f17val.Enabled = f17valiter.Enabled
			}
			f17[f17key] = f17val
		}
		cr.Spec.ForProvider.LogPublishingOptions = f17
	} else {
		cr.Spec.ForProvider.LogPublishingOptions = nil
	}
	if resp.DomainStatus.NodeToNodeEncryptionOptions != nil {
		f18 := &svcapitypes.NodeToNodeEncryptionOptions{}
		if resp.DomainStatus.NodeToNodeEncryptionOptions.Enabled != nil {
			f18.Enabled = resp.DomainStatus.NodeToNodeEncryptionOptions.Enabled
		}
		cr.Spec.ForProvider.NodeToNodeEncryptionOptions = f18
	} else {
		cr.Spec.ForProvider.NodeToNodeEncryptionOptions = nil
	}
	if resp.DomainStatus.Processing != nil {
		cr.Status.AtProvider.Processing = resp.DomainStatus.Processing
	} else {
		cr.Status.AtProvider.Processing = nil
	}
	if resp.DomainStatus.ServiceSoftwareOptions != nil {
		f20 := &svcapitypes.ServiceSoftwareOptions{}
		if resp.DomainStatus.ServiceSoftwareOptions.AutomatedUpdateDate != nil {
			f20.AutomatedUpdateDate = &metav1.Time{*resp.DomainStatus.ServiceSoftwareOptions.AutomatedUpdateDate}
		}
		if resp.DomainStatus.ServiceSoftwareOptions.Cancellable != nil {
			f20.Cancellable = resp.DomainStatus.ServiceSoftwareOptions.Cancellable
		}
		if resp.DomainStatus.ServiceSoftwareOptions.CurrentVersion != nil {
			f20.CurrentVersion = resp.DomainStatus.ServiceSoftwareOptions.CurrentVersion
		}
		if resp.DomainStatus.ServiceSoftwareOptions.Description != nil {
			f20.Description = resp.DomainStatus.ServiceSoftwareOptions.Description
		}
		if resp.DomainStatus.ServiceSoftwareOptions.NewVersion != nil {
			f20.NewVersion = resp.DomainStatus.ServiceSoftwareOptions.NewVersion
		}
		if resp.DomainStatus.ServiceSoftwareOptions.OptionalDeployment != nil {
			f20.OptionalDeployment = resp.DomainStatus.ServiceSoftwareOptions.OptionalDeployment
		}
		if resp.DomainStatus.ServiceSoftwareOptions.UpdateAvailable != nil {
			f20.UpdateAvailable = resp.DomainStatus.ServiceSoftwareOptions.UpdateAvailable
		}
		if resp.DomainStatus.ServiceSoftwareOptions.UpdateStatus != nil {
			f20.UpdateStatus = resp.DomainStatus.ServiceSoftwareOptions.UpdateStatus
		}
		cr.Status.AtProvider.ServiceSoftwareOptions = f20
	} else {
		cr.Status.AtProvider.ServiceSoftwareOptions = nil
	}
	if resp.DomainStatus.SnapshotOptions != nil {
		f21 := &svcapitypes.SnapshotOptions{}
		if resp.DomainStatus.SnapshotOptions.AutomatedSnapshotStartHour != nil {
			f21.AutomatedSnapshotStartHour = resp.DomainStatus.SnapshotOptions.AutomatedSnapshotStartHour
		}
		cr.Spec.ForProvider.SnapshotOptions = f21
	} else {
		cr.Spec.ForProvider.SnapshotOptions = nil
	}
	if resp.DomainStatus.UpgradeProcessing != nil {
		cr.Status.AtProvider.UpgradeProcessing = resp.DomainStatus.UpgradeProcessing
	} else {
		cr.Status.AtProvider.UpgradeProcessing = nil
	}
	if resp.DomainStatus.VPCOptions != nil {
		f23 := &svcapitypes.VPCOptions{}
		if resp.DomainStatus.VPCOptions.SecurityGroupIds != nil {
			f23f1 := []*string{}
			for _, f23f1iter := range resp.DomainStatus.VPCOptions.SecurityGroupIds {
				var f23f1elem string
				f23f1elem = *f23f1iter
				f23f1 = append(f23f1, &f23f1elem)
			}
			f23.SecurityGroupIDs = f23f1
		}
		if resp.DomainStatus.VPCOptions.SubnetIds != nil {
			f23f2 := []*string{}
			for _, f23f2iter := range resp.DomainStatus.VPCOptions.SubnetIds {
				var f23f2elem string
				f23f2elem = *f23f2iter
				f23f2 = append(f23f2, &f23f2elem)
			}
			f23.SubnetIDs = f23f2
		}
		cr.Spec.ForProvider.VPCOptions = f23
	} else {
		cr.Spec.ForProvider.VPCOptions = nil
	}

	return e.postCreate(ctx, cr, resp, managed.ExternalCreation{}, err)
}

func (e *external) Update(ctx context.Context, mg cpresource.Managed) (managed.ExternalUpdate, error) {
	return e.update(ctx, mg)

}

func (e *external) Delete(ctx context.Context, mg cpresource.Managed) error {
	cr, ok := mg.(*svcapitypes.ElasticsearchDomain)
	if !ok {
		return errors.New(errUnexpectedObject)
	}
	cr.Status.SetConditions(xpv1.Deleting())
	input := GenerateDeleteElasticsearchDomainInput(cr)
	ignore, err := e.preDelete(ctx, cr, input)
	if err != nil {
		return errors.Wrap(err, "pre-delete failed")
	}
	if ignore {
		return nil
	}
	resp, err := e.client.DeleteElasticsearchDomainWithContext(ctx, input)
	return e.postDelete(ctx, cr, resp, awsclient.Wrap(cpresource.Ignore(IsNotFound, err), errDelete))
}

type option func(*external)

func newExternal(kube client.Client, client svcsdkapi.ElasticsearchServiceAPI, opts []option) *external {
	e := &external{
		kube:           kube,
		client:         client,
		preObserve:     nopPreObserve,
		postObserve:    nopPostObserve,
		lateInitialize: nopLateInitialize,
		isUpToDate:     alwaysUpToDate,
		preCreate:      nopPreCreate,
		postCreate:     nopPostCreate,
		preDelete:      nopPreDelete,
		postDelete:     nopPostDelete,
		update:         nopUpdate,
	}
	for _, f := range opts {
		f(e)
	}
	return e
}

type external struct {
	kube           client.Client
	client         svcsdkapi.ElasticsearchServiceAPI
	preObserve     func(context.Context, *svcapitypes.ElasticsearchDomain, *svcsdk.DescribeElasticsearchDomainInput) error
	postObserve    func(context.Context, *svcapitypes.ElasticsearchDomain, *svcsdk.DescribeElasticsearchDomainOutput, managed.ExternalObservation, error) (managed.ExternalObservation, error)
	lateInitialize func(*svcapitypes.ElasticsearchDomainParameters, *svcsdk.DescribeElasticsearchDomainOutput) error
	isUpToDate     func(*svcapitypes.ElasticsearchDomain, *svcsdk.DescribeElasticsearchDomainOutput) (bool, error)
	preCreate      func(context.Context, *svcapitypes.ElasticsearchDomain, *svcsdk.CreateElasticsearchDomainInput) error
	postCreate     func(context.Context, *svcapitypes.ElasticsearchDomain, *svcsdk.CreateElasticsearchDomainOutput, managed.ExternalCreation, error) (managed.ExternalCreation, error)
	preDelete      func(context.Context, *svcapitypes.ElasticsearchDomain, *svcsdk.DeleteElasticsearchDomainInput) (bool, error)
	postDelete     func(context.Context, *svcapitypes.ElasticsearchDomain, *svcsdk.DeleteElasticsearchDomainOutput, error) error
	update         func(context.Context, cpresource.Managed) (managed.ExternalUpdate, error)
}

func nopPreObserve(context.Context, *svcapitypes.ElasticsearchDomain, *svcsdk.DescribeElasticsearchDomainInput) error {
	return nil
}

func nopPostObserve(_ context.Context, _ *svcapitypes.ElasticsearchDomain, _ *svcsdk.DescribeElasticsearchDomainOutput, obs managed.ExternalObservation, err error) (managed.ExternalObservation, error) {
	return obs, err
}
func nopLateInitialize(*svcapitypes.ElasticsearchDomainParameters, *svcsdk.DescribeElasticsearchDomainOutput) error {
	return nil
}
func alwaysUpToDate(*svcapitypes.ElasticsearchDomain, *svcsdk.DescribeElasticsearchDomainOutput) (bool, error) {
	return true, nil
}

func nopPreCreate(context.Context, *svcapitypes.ElasticsearchDomain, *svcsdk.CreateElasticsearchDomainInput) error {
	return nil
}
func nopPostCreate(_ context.Context, _ *svcapitypes.ElasticsearchDomain, _ *svcsdk.CreateElasticsearchDomainOutput, cre managed.ExternalCreation, err error) (managed.ExternalCreation, error) {
	return cre, err
}
func nopPreDelete(context.Context, *svcapitypes.ElasticsearchDomain, *svcsdk.DeleteElasticsearchDomainInput) (bool, error) {
	return false, nil
}
func nopPostDelete(_ context.Context, _ *svcapitypes.ElasticsearchDomain, _ *svcsdk.DeleteElasticsearchDomainOutput, err error) error {
	return err
}
func nopUpdate(context.Context, cpresource.Managed) (managed.ExternalUpdate, error) {
	return managed.ExternalUpdate{}, nil
}