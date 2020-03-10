package snapshot

import (
	"context"
	"sync"
	"time"

	"github.com/solo-io/go-utils/contextutils"
	discovery_v1alpha1 "github.com/solo-io/mesh-projects/pkg/api/discovery.zephyr.solo.io/v1alpha1"
	discovery_controllers "github.com/solo-io/mesh-projects/pkg/api/discovery.zephyr.solo.io/v1alpha1/controller"
	networking_v1alpha1 "github.com/solo-io/mesh-projects/pkg/api/networking.zephyr.solo.io/v1alpha1"
	networking_controllers "github.com/solo-io/mesh-projects/pkg/api/networking.zephyr.solo.io/v1alpha1/controller"
	"go.uber.org/zap"
)

type MeshNetworkingSnapshot struct {
	MeshServices  []*discovery_v1alpha1.MeshService
	MeshGroups    []*networking_v1alpha1.MeshGroup
	MeshWorkloads []*discovery_v1alpha1.MeshWorkload
}

type UpdatedMeshService struct {
	Old *discovery_v1alpha1.MeshService
	New *discovery_v1alpha1.MeshService
}

type UpdatedMeshGroup struct {
	Old *networking_v1alpha1.MeshGroup
	New *networking_v1alpha1.MeshGroup
}

type UpdatedMeshWorkload struct {
	Old *discovery_v1alpha1.MeshWorkload
	New *discovery_v1alpha1.MeshWorkload
}

type UpdatedResources struct {
	MeshServices  []UpdatedMeshService
	MeshGroups    []UpdatedMeshGroup
	MeshWorkloads []UpdatedMeshWorkload
}

// an implementation of `MeshNetworkingSnapshotGenerator` that is guaranteed to only ever push
// snapshots that are considered valid by the `MeshNetworkingSnapshotValidator` to its listeners
func NewMeshNetworkingSnapshotGenerator(
	ctx context.Context,
	snapshotValidator MeshNetworkingSnapshotValidator,
	meshServiceController discovery_controllers.MeshServiceController,
	meshGroupController networking_controllers.MeshGroupController,
	meshWorkloadController discovery_controllers.MeshWorkloadController,
) (MeshNetworkingSnapshotGenerator, error) {
	generator := &networkingSnapshotGenerator{
		snapshotValidator: snapshotValidator,
		snapshot:          MeshNetworkingSnapshot{},
	}

	err := meshServiceController.AddEventHandler(ctx, &discovery_controllers.MeshServiceEventHandlerFuncs{
		OnCreate: func(obj *discovery_v1alpha1.MeshService) error {
			generator.snapshotMutex.Lock()
			defer generator.snapshotMutex.Unlock()

			updatedMeshServices := append([]*discovery_v1alpha1.MeshService{}, generator.snapshot.MeshServices...)
			updatedMeshServices = append(updatedMeshServices, obj)

			updatedSnapshot := generator.snapshot
			updatedSnapshot.MeshServices = updatedMeshServices
			if generator.snapshotValidator.ValidateMeshServiceUpsert(ctx, obj, &updatedSnapshot) {
				generator.isSnapshotPushNeeded = true
				generator.snapshot = updatedSnapshot
			}

			return nil
		},
		OnUpdate: func(old, new *discovery_v1alpha1.MeshService) error {
			generator.snapshotMutex.Lock()
			defer generator.snapshotMutex.Unlock()

			var updatedMeshServices []*discovery_v1alpha1.MeshService
			for _, existingMeshService := range generator.snapshot.MeshServices {
				if existingMeshService.GetName() == old.GetName() && existingMeshService.GetNamespace() == old.GetNamespace() {
					updatedMeshServices = append(updatedMeshServices, new)
				} else {
					updatedMeshServices = append(updatedMeshServices, existingMeshService)
				}
			}

			updatedSnapshot := generator.snapshot
			updatedSnapshot.MeshServices = updatedMeshServices

			if generator.snapshotValidator.ValidateMeshServiceUpsert(ctx, new, &updatedSnapshot) {
				generator.isSnapshotPushNeeded = true
				generator.snapshot = updatedSnapshot
			}

			return nil
		},
		OnDelete: func(obj *discovery_v1alpha1.MeshService) error {
			generator.snapshotMutex.Lock()
			defer generator.snapshotMutex.Unlock()

			var updatedMeshServices []*discovery_v1alpha1.MeshService
			for _, meshService := range generator.snapshot.MeshServices {
				if meshService.GetName() == obj.GetName() && meshService.GetNamespace() == obj.GetNamespace() {
					continue
				}

				updatedMeshServices = append(updatedMeshServices, meshService)
			}

			updatedSnapshot := generator.snapshot
			updatedSnapshot.MeshServices = updatedMeshServices

			if generator.snapshotValidator.ValidateMeshServiceDelete(ctx, obj, &updatedSnapshot) {
				generator.isSnapshotPushNeeded = true
				generator.snapshot = updatedSnapshot
			}

			return nil
		},
		OnGeneric: func(obj *discovery_v1alpha1.MeshService) error {
			return nil
		},
	})
	if err != nil {
		return nil, err
	}

	err = meshGroupController.AddEventHandler(ctx, &networking_controllers.MeshGroupEventHandlerFuncs{
		OnCreate: func(obj *networking_v1alpha1.MeshGroup) error {
			generator.snapshotMutex.Lock()
			defer generator.snapshotMutex.Unlock()

			updatedMeshGroups := append([]*networking_v1alpha1.MeshGroup{}, generator.snapshot.MeshGroups...)
			updatedMeshGroups = append(updatedMeshGroups, obj)

			updatedSnapshot := generator.snapshot
			updatedSnapshot.MeshGroups = updatedMeshGroups

			if generator.snapshotValidator.ValidateMeshGroupUpsert(ctx, obj, &updatedSnapshot) {
				generator.isSnapshotPushNeeded = true
				generator.snapshot = updatedSnapshot
			}

			return nil
		},
		OnUpdate: func(old, new *networking_v1alpha1.MeshGroup) error {
			generator.snapshotMutex.Lock()
			defer generator.snapshotMutex.Unlock()

			var updatedMeshGroups []*networking_v1alpha1.MeshGroup
			for _, existingMeshGroup := range generator.snapshot.MeshGroups {
				if existingMeshGroup.GetName() == old.GetName() && existingMeshGroup.GetNamespace() == old.GetNamespace() {
					updatedMeshGroups = append(updatedMeshGroups, new)
				} else {
					updatedMeshGroups = append(updatedMeshGroups, existingMeshGroup)
				}
			}

			updatedSnapshot := generator.snapshot
			updatedSnapshot.MeshGroups = updatedMeshGroups

			if generator.snapshotValidator.ValidateMeshGroupUpsert(ctx, new, &updatedSnapshot) {
				generator.isSnapshotPushNeeded = true
				generator.snapshot = updatedSnapshot
			}

			return nil
		},
		OnDelete: func(obj *networking_v1alpha1.MeshGroup) error {
			generator.snapshotMutex.Lock()
			defer generator.snapshotMutex.Unlock()

			var updatedMeshGroups []*networking_v1alpha1.MeshGroup
			for _, meshGroup := range generator.snapshot.MeshGroups {
				if meshGroup.GetName() == obj.GetName() && meshGroup.GetNamespace() == obj.GetNamespace() {
					continue
				}

				updatedMeshGroups = append(updatedMeshGroups, meshGroup)
			}

			updatedSnapshot := generator.snapshot
			updatedSnapshot.MeshGroups = updatedMeshGroups

			if generator.snapshotValidator.ValidateMeshGroupDelete(ctx, obj, &updatedSnapshot) {
				generator.isSnapshotPushNeeded = true
				generator.snapshot = updatedSnapshot
			}

			return nil
		},
		OnGeneric: func(obj *networking_v1alpha1.MeshGroup) error {
			return nil
		},
	})
	if err != nil {
		return nil, err
	}

	err = meshWorkloadController.AddEventHandler(ctx, &discovery_controllers.MeshWorkloadEventHandlerFuncs{
		OnCreate: func(obj *discovery_v1alpha1.MeshWorkload) error {
			generator.snapshotMutex.Lock()
			defer generator.snapshotMutex.Unlock()

			updatedMeshWorkloads := append([]*discovery_v1alpha1.MeshWorkload{}, generator.snapshot.MeshWorkloads...)
			updatedMeshWorkloads = append(updatedMeshWorkloads, obj)

			updatedSnapshot := generator.snapshot
			updatedSnapshot.MeshWorkloads = updatedMeshWorkloads

			if generator.snapshotValidator.ValidateMeshWorkloadUpsert(ctx, obj, &updatedSnapshot) {
				generator.isSnapshotPushNeeded = true
				generator.snapshot = updatedSnapshot
			}

			return nil
		},
		OnUpdate: func(old, new *discovery_v1alpha1.MeshWorkload) error {
			generator.snapshotMutex.Lock()
			defer generator.snapshotMutex.Unlock()

			var updatedMeshWorkloads []*discovery_v1alpha1.MeshWorkload
			for _, existingMeshWorkload := range generator.snapshot.MeshWorkloads {
				if existingMeshWorkload.GetName() == old.GetName() && existingMeshWorkload.GetNamespace() == old.GetNamespace() {
					updatedMeshWorkloads = append(updatedMeshWorkloads, new)
				} else {
					updatedMeshWorkloads = append(updatedMeshWorkloads, existingMeshWorkload)
				}
			}

			updatedSnapshot := generator.snapshot
			updatedSnapshot.MeshWorkloads = updatedMeshWorkloads

			if generator.snapshotValidator.ValidateMeshWorkloadUpsert(ctx, new, &updatedSnapshot) {
				generator.isSnapshotPushNeeded = true
				generator.snapshot = updatedSnapshot
			}

			return nil
		},
		OnDelete: func(obj *discovery_v1alpha1.MeshWorkload) error {
			generator.snapshotMutex.Lock()
			defer generator.snapshotMutex.Unlock()

			var updatedMeshWorkloads []*discovery_v1alpha1.MeshWorkload
			for _, meshWorkload := range generator.snapshot.MeshWorkloads {
				if meshWorkload.GetName() == obj.GetName() && meshWorkload.GetNamespace() == obj.GetNamespace() {
					continue
				}

				updatedMeshWorkloads = append(updatedMeshWorkloads, meshWorkload)
			}

			updatedSnapshot := generator.snapshot
			updatedSnapshot.MeshWorkloads = updatedMeshWorkloads

			if generator.snapshotValidator.ValidateMeshWorkloadDelete(ctx, obj, &updatedSnapshot) {
				generator.isSnapshotPushNeeded = true
				generator.snapshot = updatedSnapshot
			}

			return nil
		},
		OnGeneric: func(obj *discovery_v1alpha1.MeshWorkload) error {
			return nil
		},
	})
	if err != nil {
		return nil, err
	}

	return generator, nil
}

type networkingSnapshotGenerator struct {
	snapshotValidator MeshNetworkingSnapshotValidator

	listeners     []MeshNetworkingSnapshotListener
	listenerMutex sync.Mutex

	// important that snapshot is NOT a reference- we depend on being able to copy it
	// and change fields without mutating the real thing
	// accesses to `isSnapshotPushNeeded` should be gated on the `snapshotMutex`
	snapshot MeshNetworkingSnapshot
	// version of the snapshot being sent, will appear in the logger context values
	version              uint
	isSnapshotPushNeeded bool
	snapshotMutex        sync.Mutex
}

func (f *networkingSnapshotGenerator) RegisterListener(listener MeshNetworkingSnapshotListener) {
	f.listenerMutex.Lock()
	defer f.listenerMutex.Unlock()

	f.listeners = append(f.listeners, listener)
}

func (f *networkingSnapshotGenerator) StartPushingSnapshots(ctx context.Context, snapshotFrequency time.Duration) {
	for {
		select {
		case <-ctx.Done():
			return
		case <-time.After(snapshotFrequency):
			f.snapshotMutex.Lock()
			f.listenerMutex.Lock()

			if f.isSnapshotPushNeeded {
				f.version++
				snapshotContext := contextutils.WithLoggerValues(ctx,
					zap.Uint("snapshot_version", f.version),
					zap.Int("num_mesh_services", len(f.snapshot.MeshServices)),
					zap.Int("num_mesh_workloads", len(f.snapshot.MeshWorkloads)),
					zap.Int("num_mesh_groups", len(f.snapshot.MeshGroups)),
				)
				for _, listener := range f.listeners {
					listener.Sync(snapshotContext, &f.snapshot)
				}

				f.isSnapshotPushNeeded = false
			}

			// important to unlock the mutexes in the same order as they were locked here
			// it's a runtime error to attempt to unlock an already unlocked mutex
			// if the order is changed here, a race condition could cause a repeated unlock
			f.snapshotMutex.Unlock()
			f.listenerMutex.Unlock()
		}
	}
}
