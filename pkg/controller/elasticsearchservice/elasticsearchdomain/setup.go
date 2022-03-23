package elasticsearchdomain

import (
	//"context"
    //svcsdk "github.com/aws/aws-sdk-go/service/elasticsearchservice"
    svcapitypes "github.com/crossplane/provider-aws/apis/elasticsearchservice/v1alpha1"
	ctrl "sigs.k8s.io/controller-runtime"
	//"github.com/aws/aws-sdk-go/aws"

	"github.com/crossplane/crossplane-runtime/pkg/controller"
	"github.com/crossplane/crossplane-runtime/pkg/event"
	//"github.com/crossplane/crossplane-runtime/pkg/meta"
	"github.com/crossplane/crossplane-runtime/pkg/reconciler/managed"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
)

// SetupElasticsearchDomain adds a controller that reconciles ElasticsearchDomain.
func SetupElasticsearchDomain(mgr ctrl.Manager, o controller.Options) error {
	name := managed.ControllerName(svcapitypes.ElasticsearchDomainGroupKind)
	return ctrl.NewControllerManagedBy(mgr).
		Named(name).
		WithOptions(o.ForControllerRuntime()).
		For(&svcapitypes.ElasticsearchDomain{}).
		Complete(managed.NewReconciler(mgr,
			resource.ManagedKind(svcapitypes.ElasticsearchDomainGroupVersionKind),
			//managed.WithInitializers(),
			managed.WithExternalConnecter(&connector{kube: mgr.GetClient()}),
			managed.WithPollInterval(o.PollInterval),
			managed.WithLogger(o.Logger.WithValues("controller", name)),
			managed.WithRecorder(event.NewAPIRecorder(mgr.GetEventRecorderFor(name)))))
}
