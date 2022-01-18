package redashclient

import (
	"github.com/go-openapi/runtime"
	gen_admin "github.com/recolabs/reco/redash-client/gen/client/administration"
	gen_ds "github.com/recolabs/reco/redash-client/gen/client/data_sources"
	gen_queries "github.com/recolabs/reco/redash-client/gen/client/queries"
	gen_users "github.com/recolabs/reco/redash-client/gen/client/users"
	gen_viz "github.com/recolabs/reco/redash-client/gen/client/visualizations"
)

// ClientOption is the option for Client methods.
// It is used to configure the client just before execution and thus overwrite any default behavior
type ClientOption func(*runtime.ClientOperation)
type ClientOptions []ClientOption

func (opts ClientOptions) toDataSourcesClientOptions() []gen_ds.ClientOption {
	if opts == nil {
		return nil
	}

	result := make([]gen_ds.ClientOption, len(opts))
	for i, opt := range opts {
		result[i] = func(co *runtime.ClientOperation) {
			opt(co)
		}
	}
	return result
}

func (opts ClientOptions) toUsersClientOptions() []gen_users.ClientOption {
	if opts == nil {
		return nil
	}

	result := make([]gen_users.ClientOption, len(opts))
	for i, opt := range opts {
		result[i] = func(co *runtime.ClientOperation) {
			opt(co)
		}
	}
	return result
}

func (opts ClientOptions) toVizClientOptions() []gen_viz.ClientOption {
	if opts == nil {
		return nil
	}

	result := make([]gen_viz.ClientOption, len(opts))
	for i, opt := range opts {
		result[i] = func(co *runtime.ClientOperation) {
			opt(co)
		}
	}
	return result
}

func (opts ClientOptions) toQueryClientOptions() []gen_queries.ClientOption {
	if opts == nil {
		return nil
	}

	result := make([]gen_queries.ClientOption, len(opts))
	for i, opt := range opts {
		result[i] = func(co *runtime.ClientOperation) {
			opt(co)
		}
	}
	return result
}

func (opts ClientOptions) toAdministrationClientOptions() []gen_admin.ClientOption {
	if opts == nil {
		return nil
	}

	result := make([]gen_admin.ClientOption, len(opts))
	for i, opt := range opts {
		result[i] = func(co *runtime.ClientOperation) {
			opt(co)
		}
	}
	return result
}
