package engine

import (
	"context"

	appv1 "github.com/argoproj/argo-cd/pkg/apis/application/v1alpha1"
	"github.com/argoproj/argo-cd/util/argo"
	"github.com/argoproj/argo-cd/util/db"
	"github.com/argoproj/argo-cd/util/settings"
)

type ReconciliationSettings interface {
	GetAppInstanceLabelKey() (string, error)
	GetResourcesFilter() (*settings.ResourcesFilter, error)
	GetResourceOverrides() (map[string]appv1.ResourceOverride, error)
	Subscribe(subCh chan<- *settings.ArgoCDSettings)
	Unsubscribe(subCh chan<- *settings.ArgoCDSettings)
	GetConfigManagementPlugins() ([]appv1.ConfigManagementPlugin, error)
	GetKustomizeBuildOptions() (string, error)
}

type CredentialsStore interface {
	GetCluster(ctx context.Context, name string) (*appv1.Cluster, error)
	WatchClusters(ctx context.Context, callback func(event *db.ClusterEvent)) error
	ListHelmRepositories(ctx context.Context) ([]*appv1.Repository, error)
	GetRepository(ctx context.Context, url string) (*appv1.Repository, error)
}

type AuditLogger interface {
	LogAppEvent(app *appv1.Application, info argo.EventInfo, message string)
}
