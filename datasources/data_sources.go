// Package datasources provides methods of adding, removing and listing Data Sources
package datasources

import (
	"github.com/recolabs/reco/redash-client/gen/client"
	ds "github.com/recolabs/reco/redash-client/gen/client/data_sources"
	"github.com/recolabs/reco/redash-client/gen/models"
	"github.com/recolabs/reco/redash-client/options"
)

// RequestWrapper is the struct holding the configuration for the DataSource manipulation code
type RequestWrapper struct {
	httpClient *client.Redashclient
	opts       []ds.ClientOption
}

// NewDataSourceRequest builds a DataSourceRequest object. The only validation this
// methods provides is that options argument is a valid json string
// This method builds a new data source to be added to redash, thus the ID and Syntax members will be uninitialized.
// Note - valid type and name can be found in `gen/models/data_source.go` - DataSourceName*, DataSourceType*
func NewDataSource(dataSourceType, name, optionsJSON string) (*models.DataSource, error) {
	optionsMap, err := options.MapFromString(optionsJSON)
	if err != nil {
		return nil, err
	}

	dataSource := &models.DataSource{
		Type:    &dataSourceType,
		Name:    &name,
		Options: optionsMap,
	}
	err = dataSource.Validate(nil)
	if err != nil {
		return nil, err
	}

	return dataSource, nil
}

// NewRequestWrapper returns a simple RequestWrapper object with the provided configuration
func NewRequestWrapper(httpClient *client.Redashclient, opts ...ds.ClientOption) *RequestWrapper {
	return &RequestWrapper{httpClient: httpClient, opts: opts}
}

// Add a the provided DataSource to Redash,
// only the DataSource's name, type and options string are used other struct members hjy be ignored
// Its best to use the NewDataSource method to build the data source from string variables
func (requestWrapper *RequestWrapper) Add(dataSource *models.DataSource) (*models.DataSource, error) {
	err := dataSource.Validate(nil)
	if err != nil {
		return nil, err
	}

	addParams := ds.NewPostDataSourcesParams().WithBody(ds.PostDataSourcesBody{
		Name:    *dataSource.Name,
		Options: dataSource.Options,
		Type:    *dataSource.Type,
	})

	response, err := requestWrapper.httpClient.DataSources.PostDataSources(addParams, nil, requestWrapper.opts...)
	if err != nil {
		return nil, err
	}
	return response.GetPayload(), nil
}

// List existing data sources, if the user api key is has admin permission
// all data sources are listed. Otherwise, only data souces that belong to the user group the user whose api
// key the client uses will be displayed
//
// see relvant redash implemntation:
// https://github.com/getredash/redash/blob/965db26cabfc0de83f65c91439092ed94db4de4a/redash/handlers/data_sources.py#L118
func (requestWrapper *RequestWrapper) List() (resultDataSource []*models.DataSource, err error) {
	response, err := requestWrapper.httpClient.DataSources.List(nil, nil, requestWrapper.opts...)
	if err != nil {
		return nil, err
	}

	return response.GetPayload(), nil
}

// Remove deletes an existing data sources from Redash, it does not return the data source,
// only an error if the deletion has failed
func (requestWrapper *RequestWrapper) Delete(dataSourceID int64) error {
	deleteParams := ds.NewDeleteDataSourcesIDParams().WithID(dataSourceID)
	_, err := requestWrapper.httpClient.DataSources.DeleteDataSourcesID(deleteParams, nil, requestWrapper.opts...)
	if err != nil {
		return err
	}
	return nil
}
