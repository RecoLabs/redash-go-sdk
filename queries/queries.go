package queries

import (
	"github.com/recolabs/reco/redash-client/gen/client"
	"github.com/recolabs/reco/redash-client/gen/client/queries"
	"github.com/recolabs/reco/redash-client/gen/models"
	"github.com/recolabs/reco/redash-client/options"
)

type RequestWrapper struct {
	httpClient *client.Redashclient
	opts       []queries.ClientOption
}

var (
	defaultVersion = int64(1)
	notPublish     = false
)

// NewQuery builds a Query object. The only validation this
// methods provides is that options argument is a valid json string
func NewQuery(
	name,
	optionsJSON,
	description,
	query string,
	id,
	dataSourceID,
	repeated,
	version int64,
	isDraft bool) (
	*models.Query,
	error) {
	optionsDict, err := options.MapFromString(optionsJSON)
	if err != nil {
		return nil, err
	}

	queryObject := &models.Query{
		Name:         &name,
		Options:      optionsDict,
		Description:  description,
		Schedule:     &models.Schedule{Interval: repeated},
		Query:        &query,
		DataSourceID: dataSourceID,
		IsDraft:      isDraft,
		Version:      &version,
		ID:           id,
	}

	err = queryObject.Validate(nil)
	if err != nil {
		return nil, err
	}

	return queryObject, nil
}

func NewRequestWrapper(httpClient *client.Redashclient, opts ...queries.ClientOption) *RequestWrapper {
	return &RequestWrapper{httpClient: httpClient, opts: opts}
}

// List all queries
func (requestWrapper *RequestWrapper) List() (*models.QueryList, error) {
	response, err := requestWrapper.httpClient.Queries.GetQueries(nil, nil)
	if err != nil {
		return nil, err
	}

	return response.GetPayload(), nil
}

// Get a specific query with by query ID
func (requestWrapper *RequestWrapper) Get(queryID int64) (*models.Query, error) {
	getQueryParams := queries.NewGetQueriesIDParams()
	getQueryParams.ID = queryID
	response, err := requestWrapper.httpClient.Queries.GetQueriesID(getQueryParams, nil, requestWrapper.opts...)
	if err != nil {
		return nil, err
	}

	return response.GetPayload(), nil
}

// RegenerateQueryAPIKey changes the query's API key, thus invalidating all existing query URLs
func (requestWrapper *RequestWrapper) RegenerateQueryAPIKey(queryID int64) error {
	regenerateParams := queries.NewPostQueriesIDRegenerateAPIKeyParams().WithID(queryID)
	_, err := requestWrapper.httpClient.Queries.PostQueriesIDRegenerateAPIKey(regenerateParams, nil, requestWrapper.opts...)
	if err != nil {
		return err
	}

	return nil
}

// Add adds a new query
func (requestWrapper *RequestWrapper) Add(query *models.Query) (*models.Query, error) {
	err := query.Validate(nil)
	if err != nil {
		return nil, err
	}
	addParams := queries.NewPostQueriesParams().WithBody(queries.PostQueriesBody{
		Name:         *query.Name,
		Query:        *query.Query,
		Options:      query.Options,
		DataSourceID: query.DataSourceID,
		Description:  query.Description,
		Schedule:     query.Schedule,
		IsDraft:      &query.IsDraft,
		Version:      *query.Version,
	})
	response, err := requestWrapper.httpClient.Queries.PostQueries(addParams, nil, requestWrapper.opts...)
	if err != nil {
		return nil, err
	}
	return response.GetPayload(), nil
}

// Archive query by ID
func (requestWrapper *RequestWrapper) Archive(queryID int64) error {
	archiveParams := queries.NewDeleteQueriesIDParams().WithID(queryID)
	_, err := requestWrapper.httpClient.Queries.DeleteQueriesID(archiveParams, nil, requestWrapper.opts...)
	if err != nil {
		return err
	}
	return nil
}

// Publish sets the query's status as published
func (requestWrapper *RequestWrapper) Publish(queryID int64) (*models.Query, error) {
	publishQueryParams := queries.NewPostQueriesIDParams().WithID(queryID).WithBody(
		queries.PostQueriesIDBody{ID: queryID, IsDraft: &notPublish, Version: &defaultVersion})
	response, err := requestWrapper.httpClient.Queries.PostQueriesID(publishQueryParams, nil, requestWrapper.opts...)
	if err != nil {
		return nil, err
	}

	return response.GetPayload(), nil
}

// Run executes the query, but does not check for the return value
func (requestWrapper *RequestWrapper) ExecuteQuery(queryID int64) error {
	runQueryParams := queries.NewPostQueriesIDResultsParams().WithID(queryID)
	_, err := requestWrapper.httpClient.Queries.PostQueriesIDResults(runQueryParams, nil, requestWrapper.opts...)
	return err
}

// Get result of query by ID
func (requestWrapper *RequestWrapper) GetResult(queryID int64) (*models.QueryResult, error) {
	runQueryParams := queries.NewGetQueriesIDResultsParams().WithID(queryID)
	res, err := requestWrapper.httpClient.Queries.GetQueriesIDResults(runQueryParams, nil, requestWrapper.opts...)
	if err != nil {
		return nil, err
	}
	return res.GetPayload(), nil
}
