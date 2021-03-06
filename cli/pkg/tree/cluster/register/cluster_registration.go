package register

import (
	"context"
	"fmt"
	"strings"

	"github.com/rotisserie/eris"
	"github.com/solo-io/service-mesh-hub/cli/pkg/common"
	"github.com/solo-io/service-mesh-hub/cli/pkg/options"
	cluster_internal "github.com/solo-io/service-mesh-hub/cli/pkg/tree/cluster/internal"
	zephyr_core_types "github.com/solo-io/service-mesh-hub/pkg/api/core.zephyr.solo.io/v1alpha1/types"
	cluster_registration "github.com/solo-io/service-mesh-hub/pkg/clients/cluster-registration"
	"github.com/solo-io/service-mesh-hub/pkg/kubeconfig"
	"github.com/spf13/afero"
	"helm.sh/helm/v3/pkg/chartutil"
	"sigs.k8s.io/yaml"
)

const (
	MeshctlDiscoverySource = "meshctl"
)

var (
	FailedLoadingRemoteConfig = func(err error) error {
		return eris.Wrap(err, "Failed to load the kube config for the remote cluster")
	}
	FailedToCreateAuthToken = func(saRef *zephyr_core_types.ResourceRef, remoteKubeConfig, remoteContext string) string {
		return fmt.Sprintf("Failed to create an auth token for service account %s.%s in cluster "+
			"pointed to by kube config %s with context %s. This operation is not atomic, so the service account may "+
			"have been created and left in the cluster while a later step failed. \n",
			saRef.GetNamespace(), saRef.GetName(), remoteKubeConfig, remoteContext)
	}
	FailedToWriteKubeCluster = func(err error) error {
		return eris.Wrap(err, "Could not write KubernetesCluster resource to master cluster")
	}
	FailedToReadValuesFiles = func(err error, valuesFiles []string) error {
		return eris.Wrapf(err, "Failed to read values file: [%s]", strings.Join(valuesFiles, ", "))
	}
)

// write a new kube config secret to the master cluster containing creds for talking to the remote cluster as the given service account
func RegisterCluster(
	ctx context.Context,
	kubeClientsFactory common.KubeClientsFactory,
	clientsFactory common.ClientsFactory,
	opts *options.Options,
	kubeLoader kubeconfig.KubeLoader,
	fs afero.Fs,
) error {
	if err := cluster_internal.VerifyRemoteContextFlags(opts); err != nil {
		return err
	}
	if err := cluster_internal.VerifyMasterCluster(clientsFactory, opts); err != nil {
		return err
	}
	registerOpts := opts.Cluster.Register
	// default the remote kube config/context to the root settings
	remoteConfigPath, remoteContext := opts.Root.KubeConfig, opts.Root.KubeContext
	if registerOpts.RemoteKubeConfig != "" {
		// if we specified a kube config for the remote cluster, use that instead
		remoteConfigPath = registerOpts.RemoteKubeConfig
	}
	// if we didn't have a context from the root, or if we had an override for the
	// remote context, use the remote context instead
	if remoteContext == "" || registerOpts.RemoteContext != "" {
		remoteContext = registerOpts.RemoteContext
	}
	remoteConfig, err := kubeLoader.GetConfigWithContext("", remoteConfigPath, remoteContext)
	if err != nil {
		return err
	}
	masterCfg, err := kubeLoader.GetRestConfigForContext(opts.Root.KubeConfig, opts.Root.KubeContext)
	if err != nil {
		return err
	}
	kubeClients, err := kubeClientsFactory(masterCfg, opts.Root.WriteNamespace)
	if err != nil {
		return err
	}
	helmOverrideValues, err := readValuesFiles(fs, opts.Cluster.Register.CsrAgentHelmChartValueFileNames)
	if err != nil {
		return FailedToReadValuesFiles(err, opts.Cluster.Register.CsrAgentHelmChartValueFileNames)
	}
	return kubeClients.ClusterRegistrationClient.Register(
		ctx,
		remoteConfig,
		registerOpts.RemoteClusterName,
		registerOpts.RemoteWriteNamespace,
		remoteContext,
		MeshctlDiscoverySource,
		cluster_registration.ClusterRegisterOpts{
			Overwrite:                        registerOpts.Overwrite,
			UseDevCsrAgentChart:              registerOpts.UseDevCsrAgentChart,
			LocalClusterDomainOverride:       registerOpts.LocalClusterDomainOverride,
			CsrAgentHelmChartValuesFileNames: helmOverrideValues,
		},
	)
}

func readValuesFiles(fs afero.Fs, valuesFiles []string) (map[string]interface{}, error) {
	if len(valuesFiles) == 0 {
		return nil, nil
	}
	mergedMap := map[string]interface{}{}
	for _, valuesFile := range valuesFiles {
		bytes, err := afero.ReadFile(fs, valuesFile)
		if err != nil {
			return nil, err
		}
		currentMap := map[string]interface{}{}
		if err := yaml.Unmarshal(bytes, &currentMap); err != nil {
			return nil, err
		}
		mergedMap = chartutil.CoalesceTables(mergedMap, currentMap)
	}
	return mergedMap, nil
}
