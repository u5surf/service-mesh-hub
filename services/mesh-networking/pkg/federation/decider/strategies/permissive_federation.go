package strategies

import (
	"context"

	zephyr_core_types "github.com/solo-io/service-mesh-hub/pkg/api/core.zephyr.solo.io/v1alpha1/types"
	zephyr_discovery "github.com/solo-io/service-mesh-hub/pkg/api/discovery.zephyr.solo.io/v1alpha1"
	zephyr_discovery_types "github.com/solo-io/service-mesh-hub/pkg/api/discovery.zephyr.solo.io/v1alpha1/types"
	zephyr_networking "github.com/solo-io/service-mesh-hub/pkg/api/networking.zephyr.solo.io/v1alpha1"
	"github.com/solo-io/service-mesh-hub/services/mesh-networking/pkg/federation/dns"
)

func NewPermissiveFederation(meshServiceClient zephyr_discovery.MeshServiceClient) FederationStrategy {
	return &permissiveFederation{
		meshServiceClient: meshServiceClient,
	}
}

type permissiveFederation struct {
	meshServiceClient zephyr_discovery.MeshServiceClient
}

func (p *permissiveFederation) WriteFederationToServices(
	ctx context.Context,
	vm *zephyr_networking.VirtualMesh,
	meshNameToMetadata MeshNameToMetadata,
) error {
	for _, serverMeshRef := range vm.Spec.Meshes {
		serverMeshMetadata, ok := meshNameToMetadata[serverMeshRef.GetName()]
		if !ok {
			return MeshMetadataMissing(serverMeshRef.GetName())
		}

		servicesInMesh := serverMeshMetadata.MeshServices

		var federatedToWorkloads []*zephyr_core_types.ResourceRef
		for _, clientMesh := range vm.Spec.Meshes {
			// skip `serverMeshRef` - we don't want to federate a service to the same mesh that it's in
			if clientMesh.GetName() == serverMeshRef.GetName() && clientMesh.GetNamespace() == serverMeshRef.GetNamespace() {
				continue
			}

			clientMeshMetadata, ok := meshNameToMetadata[clientMesh.GetName()]
			if !ok {
				return MeshMetadataMissing(clientMesh.GetName())
			}

			// get the workloads belonging to this mesh (the mesh that the clients are in)
			for _, workload := range clientMeshMetadata.MeshWorkloads {
				federatedToWorkloads = append(federatedToWorkloads, &zephyr_core_types.ResourceRef{
					Name:      workload.GetName(),
					Namespace: workload.GetNamespace(),
				})
			}
		}

		for _, service := range servicesInMesh {
			serviceClusterName := serverMeshMetadata.ClusterName

			service.Spec.Federation = &zephyr_discovery_types.MeshServiceSpec_Federation{
				MulticlusterDnsName:  dns.BuildMulticlusterDnsName(service.Spec.GetKubeService().GetRef(), serviceClusterName),
				FederatedToWorkloads: federatedToWorkloads,
			}
		}

		err := updateServices(ctx, servicesInMesh, p.meshServiceClient)
		if err != nil {
			return err
		}
	}

	return nil
}
