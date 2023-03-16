package wasmmodules

import (
	"context"
	"github.com/opencontainers/go-digest"
	"time"
)

// WasmModule - metadata
type WasmModule struct {
	// Name of the wasm module
	// Required: true
	Name string

	// Target describes the root content for this wasm module
	Target WasmModuleSpec

	CreatedAt, UpdatedAt time.Time
}

type WasmModuleSpec struct {
	// Basic info >>>
	Author string
	// <<< Basic info

	// Descriptor >>>
	// Digest is the digest of the wasm module
	Digest digest.Digest

	// Size is the size of the wasm module
	Size int64

	// Name of the wasm runtime running the module
	Runtime string

	// URL specifies the URL from which this wasm module MAY be downloaded
	URL string

	// Annotations contains arbitrary metadata relating to the targeted wasm module.
	Annotations map[string]string
	// <<< Descriptor

	// Config >>>
	// Cmd defines the default command for the wasm module / wasm instance
	Cmd []string
	// <<< Config
}

// Store and interact with wasm modules
type Store interface {
	Get(ctx context.Context, name string) (WasmModule, error)
	List(ctx context.Context, filters ...string) ([]WasmModule, error)
	Create(ctx context.Context, wasmModule WasmModule) (WasmModule, error)

	// Update will replace the data in the store with the provided wasm. If
	// one or more fieldpaths are provided, only those fields will be updated.
	Update(ctx context.Context, wasmModule WasmModule, fieldpaths ...string) (WasmModule, error)
	Delete(ctx context.Context, name string) error
}