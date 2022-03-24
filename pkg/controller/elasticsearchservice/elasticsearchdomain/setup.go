package elasticsearchdomain

import (
	"context"
	svcsdk "github.com/aws/aws-sdk-go/service/elasticsearchservice"
	svcapitypes "github.com/crossplane/provider-aws/apis/elasticsearchservice/v1alpha1"
	ctrl "sigs.k8s.io/controller-runtime"

	//"github.com/aws/aws-sdk-go/aws"

	"github.com/crossplane/crossplane-runtime/pkg/controller"
	"github.com/crossplane/crossplane-runtime/pkg/event"

	"github.com/crossplane/crossplane-runtime/pkg/meta"
	"github.com/crossplane/crossplane-runtime/pkg/reconciler/managed"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
    awsclients "github.com/crossplane/provider-aws/pkg/clients"
)

// SetupElasticsearchDomain adds a controller that reconciles ElasticsearchDomain.
func SetupElasticsearchDomain(mgr ctrl.Manager, o controller.Options) error {
	name := managed.ControllerName(svcapitypes.ElasticsearchDomainGroupKind)
	opts := []option{
		func(e *external) {
			e.preObserve = preObserve
			//e.postObserve = postObserve
			//e.lateInitialize = lateInitialize
			//e.isUpToDate = isUpToDate
			e.preCreate = preCreate
			//e.postCreate = postCreate
			e.preDelete = preDelete
			//e.postDelete = postDelete
		},
	}
	return ctrl.NewControllerManagedBy(mgr).
		Named(name).
		WithOptions(o.ForControllerRuntime()).
		For(&svcapitypes.ElasticsearchDomain{}).
		Complete(managed.NewReconciler(mgr,
			resource.ManagedKind(svcapitypes.ElasticsearchDomainGroupVersionKind),
			managed.WithExternalConnecter(&connector{kube: mgr.GetClient(), opts: opts}),
			managed.WithPollInterval(o.PollInterval),
			managed.WithLogger(o.Logger.WithValues("controller", name)),
			managed.WithRecorder(event.NewAPIRecorder(mgr.GetEventRecorderFor(name)))))
}


func postObserve(_ context.Context, cr *svcapitypes.ElasticsearchDomain, obj *svcsdk.DescribeElasticsearchDomainInput) error {
	obj.DomainName = awsclients.String(meta.GetExternalName(cr))
	return nil
}

func preDelete(_ context.Context, cr *svcapitypes.ElasticsearchDomain, obj *svcsdk.DeleteElasticsearchDomainInput) (bool, error) {
	obj.DomainName = awsclients.String(meta.GetExternalName(cr))
	return false, nil
}

func preCreate(_ context.Context, cr *svcapitypes.ElasticsearchDomain, obj *svcsdk.CreateElasticsearchDomainInput) error {
	obj.DomainName = awsclients.String(meta.GetExternalName(cr))
	return nil
}

