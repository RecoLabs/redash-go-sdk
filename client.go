// A simple go client for Redash, wraps some Redash's REST API in go methods.
// Client Initialization
//
// import (
//    redashclient "github.com/recolabs/reco/redash-client/redashClient"
//    "github.com/recolabs/reco/redash-client/gen/client"
// )
// redashClient := redashclient.NewClient(
// 	 "{API_KEY}",
//	 &client.TransportConfig{
//		 Host: "{HOST_ADDRESS}",
//	 })
//
// Usage
//
// Calls in the client are of the form:
//  client.<Queries/Visualizations/Users/DataSources>.<Method>(...)
//
// For example:
// List the current queries in Redash
//  redashClient.Queries.List()
//
//
// Add a query to Redash
//  queryString := "SELECT email, count(*) public.table GROUP BY name"
//  redashClient.Queries.Add(1, "Example Query", queryString, []queries.Parameters{}))

package redashclient

import (
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/recolabs/reco/redash-client/administration"
	"github.com/recolabs/reco/redash-client/datasources"
	"github.com/recolabs/reco/redash-client/gen/client"
	"github.com/recolabs/reco/redash-client/queries"
	"github.com/recolabs/reco/redash-client/users"
	"github.com/recolabs/reco/redash-client/visualizations"
)

// Client wrapping ReDash's rest API
type Client struct {
	Queries        *queries.RequestWrapper
	Visualizations *visualizations.RequestWrapper
	Users          *users.RequestWrapper
	DataSources    *datasources.RequestWrapper
	Administration *administration.RequestWrapper
	apiKey         string
}

// Build A new client object with every RequestWrapper using the provided configuration
func NewClient(apiKey string, cfg *client.TransportConfig, opts ...ClientOption) Client {
	if cfg == nil {
		cfg = client.DefaultTransportConfig()
	}
	if cfg.Host == "" {
		cfg.Host = client.DefaultHost
	}
	if cfg.BasePath == "" {
		cfg.BasePath = client.DefaultBasePath
	}
	if cfg.Schemes == nil {
		cfg.Schemes = client.DefaultSchemes
	}

	// create transport and client
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	scheme := selectScheme(cfg.Schemes)

	return NewClientWithTransport(apiKey, transport, scheme, opts...)
}

// Build A new client object with every RequestWrapper using the provided transport
func NewClientWithTransport(apiKey string, transport *httptransport.Runtime, scheme string, opts ...ClientOption) Client {
	// auto authorize every api call
	transport.DefaultAuthentication = httptransport.APIKeyAuth("Authorization", "header", apiKey)

	var optsArray ClientOptions = opts
	dsClientOption := optsArray.toDataSourcesClientOptions()
	usersClientOption := optsArray.toUsersClientOptions()
	vizClientOption := optsArray.toVizClientOptions()
	queriesClientOption := optsArray.toQueryClientOptions()
	administrationClientOption := optsArray.toAdministrationClientOptions()

	httpClient := client.New(transport, nil)
	return Client{
		Queries:        queries.NewRequestWrapper(httpClient, queriesClientOption...),
		Visualizations: visualizations.NewRequestWrapper(httpClient, transport.Host, scheme, apiKey, vizClientOption...),
		Users:          users.NewRequestWrapper(httpClient, usersClientOption...),
		DataSources:    datasources.NewRequestWrapper(httpClient, dsClientOption...),
		Administration: administration.NewRequestWrapper(httpClient, administrationClientOption...),
		apiKey:         apiKey,
	}
}

// GetAPIKey return the apiKey that used to initialize the client
func GetAPIKey(clientRef *Client) string {
	return clientRef.apiKey
}

const https = "https"

// Select the scheme to use based on the provided schemes - this function is a duplicate of the "runtime" package
// due to the fact that the "runtime" package function is not exported
func selectScheme(schemes []string) string {
	schLen := len(schemes)
	if schLen == 0 {
		return ""
	}

	scheme := schemes[0]
	// prefer https, but skip when not possible
	if scheme != https && schLen > 1 {
		for _, sch := range schemes {
			if sch == https {
				scheme = sch
				break
			}
		}
	}
	return scheme
}
