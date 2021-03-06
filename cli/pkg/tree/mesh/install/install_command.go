package mesh_install

import (
	"io"
	"strings"

	"github.com/google/wire"
	"github.com/rotisserie/eris"
	"github.com/solo-io/service-mesh-hub/cli/pkg/cliconstants"
	"github.com/solo-io/service-mesh-hub/cli/pkg/common"
	"github.com/solo-io/service-mesh-hub/cli/pkg/common/files"
	"github.com/solo-io/service-mesh-hub/cli/pkg/options"
	install_istio "github.com/solo-io/service-mesh-hub/cli/pkg/tree/mesh/install/istio"
	zephyr_core_types "github.com/solo-io/service-mesh-hub/pkg/api/core.zephyr.solo.io/v1alpha1/types"
	"github.com/solo-io/service-mesh-hub/pkg/common/docker"
	"github.com/solo-io/service-mesh-hub/pkg/kubeconfig"
	"github.com/spf13/cobra"
)

type MeshInstallCommand *cobra.Command

var (
	MeshInstallProviderSet = wire.NewSet(
		MeshInstallRootCmd,
	)
	validMeshTypes = []string{
		strings.ToLower(zephyr_core_types.MeshType_ISTIO.String()),
	}
	UnsupportedMeshTypeError = func(meshType string) error {
		return eris.Errorf(
			"Mesh Type: (%s) is not one of the supported Mesh types [%s]",
			meshType,
			strings.Join(validMeshTypes, "|"),
		)
	}
)

func MeshInstallRootCmd(
	clientsFactory common.ClientsFactory,
	opts *options.Options,
	out io.Writer,
	in io.Reader,
	kubeLoader kubeconfig.KubeLoader,
	imageNameParser docker.ImageNameParser,
	fileReader files.FileReader,
) MeshInstallCommand {
	installCommand := cliconstants.MeshInstallCommand(validMeshTypes)
	cmd := &cobra.Command{
		Use:     installCommand.Use,
		Short:   installCommand.Short,
		Aliases: installCommand.Aliases,
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			meshType := args[0]

			switch strings.ToUpper(meshType) {
			case zephyr_core_types.MeshType_ISTIO.String():
				istioInstaller, err := install_istio.NewIstioInstaller(
					out,
					in,
					clientsFactory,
					opts,
					opts.Root.KubeConfig,
					opts.Root.KubeContext,
					kubeLoader,
					imageNameParser,
					fileReader,
				)
				if err != nil {
					return err
				}
				return istioInstaller.Install()
			default:
				return UnsupportedMeshTypeError(meshType)
			}
		},
	}

	options.AddMeshInstallFlags(cmd, opts)

	return cmd
}
