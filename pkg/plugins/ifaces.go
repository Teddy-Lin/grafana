package plugins

import (
	"context"
	"io"

	"github.com/grafana/grafana-plugin-sdk-go/backend"

	"github.com/grafana/grafana/pkg/plugins/backendplugin"
	"github.com/grafana/grafana/pkg/plugins/repository"
)

// Store is the storage for plugins.
type Store interface {
	// Plugin finds a plugin by its ID.
	Plugin(ctx context.Context, pluginID string) (PluginDTO, bool)
	// Plugins returns plugins by their requested type.
	Plugins(ctx context.Context, pluginTypes ...Type) []PluginDTO
	// Add adds a plugin from the repository to the store.
	Add(ctx context.Context, pluginID, version string, repo repository.Service, opts CompatabilityOpts) error
	// Remove removes a plugin from the store.
	Remove(ctx context.Context, pluginID string) error
}

type CompatabilityOpts struct {
	GrafanaVersion string
}

// Loader is responsible for loading plugins from the file system.
type Loader interface {
	// Load will return a list of plugins found in the provided file system paths.
	Load(ctx context.Context, class Class, paths []string, ignore map[string]struct{}) ([]*Plugin, error)
}

// Client is used to communicate with backend plugin implementations.
type Client interface {
	backend.QueryDataHandler
	backend.CheckHealthHandler
	backend.StreamHandler
	backend.CallResourceHandler
	backend.CollectMetricsHandler
}

// BackendFactoryProvider provides a backend factory for a provided plugin.
type BackendFactoryProvider interface {
	BackendFactory(ctx context.Context, p *Plugin) backendplugin.PluginFactoryFunc
}

type RendererManager interface {
	// Renderer returns a renderer plugin.
	Renderer() *Plugin
}

type StaticRouteResolver interface {
	Routes() []*StaticRoute
}

type ErrorResolver interface {
	PluginErrors() []*Error
}

type PluginLoaderAuthorizer interface {
	// CanLoadPlugin confirms if a plugin is authorized to load
	CanLoadPlugin(plugin *Plugin) bool
}

// ListPluginDashboardFilesArgs list plugin dashboard files argument model.
type ListPluginDashboardFilesArgs struct {
	PluginID string
}

// GetPluginDashboardFilesArgs list plugin dashboard files result model.
type ListPluginDashboardFilesResult struct {
	FileReferences []string
}

// GetPluginDashboardFileContentsArgs get plugin dashboard file content argument model.
type GetPluginDashboardFileContentsArgs struct {
	PluginID      string
	FileReference string
}

// GetPluginDashboardFileContentsResult get plugin dashboard file content result model.
type GetPluginDashboardFileContentsResult struct {
	Content io.ReadCloser
}

// DashboardFileStore is the interface for plugin dashboard file storage.
type DashboardFileStore interface {
	// ListPluginDashboardFiles lists plugin dashboard files.
	ListPluginDashboardFiles(ctx context.Context, args *ListPluginDashboardFilesArgs) (*ListPluginDashboardFilesResult, error)

	// GetPluginDashboardFileContents gets the referenced plugin dashboard file content.
	GetPluginDashboardFileContents(ctx context.Context, args *GetPluginDashboardFileContentsArgs) (*GetPluginDashboardFileContentsResult, error)
}
