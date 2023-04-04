package wasmdealer

import (
	"context"

	api "github.com/containerd/containerd/api/services/wasmdealer/v1"
	"github.com/containerd/containerd/plugin"
	"github.com/containerd/containerd/runtime"
	"github.com/containerd/containerd/services"
	"google.golang.org/grpc"
)

func init() {
	plugin.Register(&plugin.Registration{
		Type:     plugin.ServicePlugin,
		ID:       services.WasmdealerService,
		Requires: []plugin.Type{
			plugin.RuntimePluginV2,
		},
		InitFn:   initLocal,
	})
}

func initLocal(ic *plugin.InitContext) (interface{}, error) {
  v2r, err := ic.GetByID(plugin.RuntimePluginV2, "task")
	if err != nil {
		return nil, err
	}
  l := &local {
    runtime: v2r.(runtime.PlatformRuntime),
  }

  return l, nil

}

type local struct {
  runtime runtime.PlatformRuntime
}

func (l *local) Create(ctx context.Context, r *api.CreateTaskRequest, _ ...grpc.CallOption) (*api.CreateTaskResponse, error) {
	// opts := runtime.CreateOpts{
	// 	Spec: container.Spec,
	// 	IO: runtime.IO{
	// 		Stdin:    r.Stdin,
	// 		Stdout:   r.Stdout,
	// 		Stderr:   r.Stderr,
	// 		Terminal: r.Terminal,
	// 	},
	// 	Checkpoint:     checkpointPath,
	// 	Runtime:        container.Runtime.Name,
	// 	RuntimeOptions: container.Runtime.Options,
	// 	TaskOptions:    r.Options,
	// }
  return &api.CreateTaskResponse{
    ContainerId: "youtest",
  }, nil
}

