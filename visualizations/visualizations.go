package visualizations

import (
	"encoding/json"
	"fmt"

	"github.com/recolabs/redash-go-sdk/gen/client"
	"github.com/recolabs/redash-go-sdk/gen/models"
	"github.com/recolabs/redash-go-sdk/options"
	vis "github.com/recolabs/redash-go-sdk/gen/client/visualizations"
)

const (
	EmbedQueryURIFormat = "/embed/query/%d/visualization/%d"
)

type RequestWrapper struct {
	httpClient *client.Redashclient
	host       string
	scheme     string
	opts       []vis.ClientOption
	apiKey     string
}

// NewVisualization builds a Visualization object. The only validation this
// methods provides is that options argument is a valid json string
// Note - valid type and name can be found in `gen/models/visualizations.go` - VisualizationType*, VisualizationName*
func NewVisualization(visualizationType, name, optionsJSON, description string, queryID int64) (*models.Visualization, error) {
	optionsMap, err := options.MapFromString(optionsJSON)
	if err != nil {
		return nil, err
	}

	visualization := &models.Visualization{
		Type:        &visualizationType,
		Name:        &name,
		QueryID:     &queryID,
		Options:     optionsMap,
		Description: description,
	}

	err = visualization.Validate(nil)
	if err != nil {
		return nil, err
	}

	return visualization, nil
}

func NewRequestWrapper(httpClient *client.Redashclient, host, scheme, apiKey string, opts ...vis.ClientOption) *RequestWrapper {
	return &RequestWrapper{httpClient: httpClient, host: host, apiKey: apiKey, scheme: scheme, opts: opts}
}

// GetURL gets a visualization URL of visualization number {visualizationID} of the {queryID} query using the query_api key
func (requestWrapper *RequestWrapper) GetURL(visualizationID, queryID int64,
	hideParameters, hideHeader, hideLink, hideTimestamp bool) (visualizationURL string) {
	uri := fmt.Sprintf(EmbedQueryURIFormat, queryID, visualizationID)
	url := fmt.Sprintf("%s://%s%s?api_key=%s", requestWrapper.scheme, requestWrapper.host, uri, requestWrapper.apiKey)
	if hideParameters {
		url = fmt.Sprintf("%s&hide_parameters", url)
	}
	if hideHeader {
		url = fmt.Sprintf("%s&hide_header", url)
	}
	if hideLink {
		url = fmt.Sprintf("%s&hide_link", url)
	}
	if hideTimestamp {
		url = fmt.Sprintf("%s&hide_timestamp", url)
	}
	return url
}

// Add a new visualization to the query (the query is defined is visualization.QueryID
func (requestWrapper *RequestWrapper) Add(visualization *models.Visualization) (resultVisualization *models.Visualization, err error) {
	err = visualization.Validate(nil)
	if err != nil {
		return nil, err
	}
	addParams := vis.NewPostVisualizationsParams().WithBody(vis.PostVisualizationsBody{
		Name:    *visualization.Name,
		Type:    *visualization.Type,
		Options: visualization.Options,
		QueryID: *visualization.QueryID,
	})

	response, err := requestWrapper.httpClient.Visualizations.PostVisualizations(addParams, nil, requestWrapper.opts...)
	if err != nil {
		return nil, err
	}
	return response.GetPayload(), nil
}

// Remove an existing visualization to the query(the query is defined is visualization.QueryID)
func (requestWrapper *RequestWrapper) Delete(visualizationID int64) error {
	deleteParams := vis.NewDeleteVisualizationsIDParams().WithID(visualizationID)
	_, err := requestWrapper.httpClient.Visualizations.DeleteVisualizationsID(deleteParams, nil, requestWrapper.opts...)
	if err != nil {
		return err
	}
	return nil
}

func ToVisualization(visualizationData string) (*models.Visualization, error) {
	var visObj models.Visualization
	err := json.Unmarshal([]byte(visualizationData), &visObj)
	return &visObj, err
}
