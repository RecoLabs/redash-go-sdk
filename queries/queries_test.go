package queries

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/recolabs/redash-go-sdk/gen/client"
	"github.com/recolabs/redash-go-sdk/gen/models"
	"github.com/recolabs/redash-go-sdk/options"
	gen_queries "github.com/recolabs/redash-go-sdk/gen/client/queries"
	mock_queries "github.com/recolabs/redash-go-sdk/mocks/queries"
)

const (
	testAddress  = "localhost:5005"
	queryOptions = `{
		"parameters": []
		}`

	malformedQueryOptions = `{
			"parameters": [],
			}`

	dataSourceID       = 1
	id                 = 1
	defaultDescription = "Description"
	defaultRepeated    = int64(1000)
	queryID            = 1
)

var (
	isDraft       = false
	queryName     = "test query"
	queryContent  = "SELECT * FROM db.public.table"
	optionsMap, _ = options.MapFromString(queryOptions)
)

func TestNewQuery(t *testing.T) {
	type args struct {
		name         string
		optionsJSON  string
		query        string
		dataSourceID int64
		id           int64
		description  string
		repeated     int64
	}
	tests := []struct {
		name    string
		args    args
		want    *models.Query
		wantErr bool
	}{
		{
			name: "Create a new data source successfully",
			args: args{
				name: queryName, optionsJSON: queryOptions, query: queryContent,
				id: id, dataSourceID: dataSourceID, description: defaultDescription,
				repeated: defaultRepeated,
			},
			want: &models.Query{
				Name:         &queryName,
				Query:        &queryContent,
				Options:      optionsMap,
				DataSourceID: dataSourceID,
				Description:  defaultDescription,
				ID:           id,
				Version:      &defaultVersion,
				IsDraft:      isDraft,
				Schedule:     &models.Schedule{Interval: defaultRepeated},
			},
			wantErr: false,
		},
		{
			name: "Create a new data source without a name",
			args: args{
				optionsJSON:  malformedQueryOptions,
				query:        queryContent,
				dataSourceID: dataSourceID,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Create a new data source without a valid json options string",
			args: args{
				name:         queryName,
				optionsJSON:  malformedQueryOptions,
				query:        queryContent,
				dataSourceID: dataSourceID,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewQuery(
				tt.args.name,
				tt.args.optionsJSON,
				tt.args.description,
				tt.args.query,
				tt.args.id,
				tt.args.dataSourceID,
				tt.args.repeated,
				defaultVersion,
				isDraft)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewQuery() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRequestWrapper_List(t *testing.T) {
	query, _ := NewQuery(queryName, queryOptions, defaultDescription, queryContent, dataSourceID, id, defaultRepeated, defaultVersion, isDraft)
	happyFlowClient := client.NewHTTPClientWithConfig(nil,
		&client.TransportConfig{Host: testAddress, BasePath: "/api"})
	successfullyListQueriesMock := mock_queries.ClientService{}
	happyFlowClient.Queries = &successfullyListQueriesMock
	successfullyListQueriesMock.On(
		"GetQueries", (*gen_queries.GetQueriesParams)(nil), nil).
		Return(
			&gen_queries.GetQueriesOK{Payload: &models.QueryList{Results: []*models.Query{query}}}, nil)
	serverErrorClient := client.NewHTTPClientWithConfig(nil,
		&client.TransportConfig{Host: testAddress, BasePath: "/api"})
	serverErrQueryMock := mock_queries.ClientService{}
	serverErrorClient.Queries = &serverErrQueryMock
	serverErrQueryMock.On("GetQueries", (*gen_queries.GetQueriesParams)(nil), nil).
		Return(nil, fmt.Errorf("Internal Server Error"))

	type fields struct {
		httpClient *client.Redashclient
	}
	tests := []struct {
		name    string
		fields  fields
		want    *models.QueryList
		wantErr bool
	}{
		{
			name:    "Successfully list queries",
			fields:  fields{httpClient: happyFlowClient},
			want:    &models.QueryList{Results: []*models.Query{query}},
			wantErr: false,
		},
		{
			name:    "Listing data sources fails due toserver error",
			fields:  fields{httpClient: serverErrorClient},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &RequestWrapper{
				httpClient: tt.fields.httpClient,
			}
			got, err := q.List()
			if (err != nil) != tt.wantErr {
				t.Errorf("RequestWrapper.List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RequestWrapper.List() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRequestWrapper_Get(t *testing.T) {
	getQueryParams := gen_queries.NewGetQueriesIDParams()
	getQueryParams.ID = queryID
	query, _ := NewQuery(queryName, queryOptions, defaultDescription, queryContent, dataSourceID, id, defaultRepeated, defaultVersion, isDraft)
	happyFlowClient := client.NewHTTPClientWithConfig(nil,
		&client.TransportConfig{Host: testAddress, BasePath: "/api"})
	successfullyListQueriesMock := mock_queries.ClientService{}
	happyFlowClient.Queries = &successfullyListQueriesMock
	successfullyListQueriesMock.On(
		"GetQueriesID", getQueryParams, nil).Return(
		&gen_queries.GetQueriesIDOK{Payload: query}, nil)
	serverErrorClient := client.NewHTTPClientWithConfig(nil,
		&client.TransportConfig{Host: testAddress, BasePath: "/api"})
	serverErrQueryMock := mock_queries.ClientService{}
	serverErrorClient.Queries = &serverErrQueryMock
	serverErrQueryMock.On("GetQueriesID", getQueryParams, nil).
		Return(nil, fmt.Errorf("Internal Server Error"))

	type fields struct {
		httpClient *client.Redashclient
	}
	type args struct {
		queryID int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.Query
		wantErr bool
	}{
		{
			name:    "Successfully get query",
			fields:  fields{httpClient: happyFlowClient},
			args:    args{queryID: queryID},
			want:    query,
			wantErr: false,
		},
		{
			name:    "Failed to get query due to a server error",
			fields:  fields{httpClient: serverErrorClient},
			args:    args{queryID: queryID},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &RequestWrapper{
				httpClient: tt.fields.httpClient,
			}
			got, err := q.Get(tt.args.queryID)
			if (err != nil) != tt.wantErr {
				t.Errorf("RequestWrapper.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RequestWrapper.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRequestWrapper_RegenerateQueryAPIKey(t *testing.T) { //nolint:dupl // tests should be explicit
	regenerateParams := gen_queries.NewPostQueriesIDRegenerateAPIKeyParams().WithID(queryID)
	happyFlowClient := client.NewHTTPClientWithConfig(nil,
		&client.TransportConfig{Host: testAddress, BasePath: "/api"})
	successfullyListQueriesMock := mock_queries.ClientService{}
	happyFlowClient.Queries = &successfullyListQueriesMock
	successfullyListQueriesMock.On("PostQueriesIDRegenerateAPIKey", regenerateParams, nil).
		Return(nil, nil)
	serverErrorClient := client.NewHTTPClientWithConfig(nil,
		&client.TransportConfig{Host: testAddress, BasePath: "/api"})
	serverErrQueryMock := mock_queries.ClientService{}
	serverErrorClient.Queries = &serverErrQueryMock
	serverErrQueryMock.On("PostQueriesIDRegenerateAPIKey", regenerateParams, nil).
		Return(nil, fmt.Errorf("Internal Server Error"))

	type fields struct {
		httpClient *client.Redashclient
	}
	type args struct {
		queryID int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "Successfully regenerate the query's API key",
			fields:  fields{httpClient: happyFlowClient},
			args:    args{queryID: queryID},
			wantErr: false,
		},
		{
			name:    "API key regeneration failed  due to a server error",
			fields:  fields{httpClient: serverErrorClient},
			args:    args{queryID: queryID},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &RequestWrapper{
				httpClient: tt.fields.httpClient,
			}
			if err := q.RegenerateQueryAPIKey(tt.args.queryID); (err != nil) != tt.wantErr {
				t.Errorf("RequestWrapper.RegenerateQueryAPIKey() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRequestWrapper_Add(t *testing.T) {
	addParams := gen_queries.NewPostQueriesParams().WithBody(gen_queries.PostQueriesBody{
		Name:         queryName,
		Description:  defaultDescription,
		Query:        queryContent,
		Options:      optionsMap,
		DataSourceID: dataSourceID,
		Version:      defaultVersion,
		IsDraft:      &isDraft,
		Schedule:     &models.Schedule{Interval: defaultRepeated},
	})

	invalidQuery := models.Query{
		Options:      queryOptions,
		Query:        &queryContent,
		DataSourceID: dataSourceID,
		Description:  defaultDescription,
	}
	query, _ := NewQuery(queryName, queryOptions, defaultDescription, queryContent, id, dataSourceID, defaultRepeated, 1, isDraft)
	happyFlowClient := client.NewHTTPClientWithConfig(nil,
		&client.TransportConfig{Host: testAddress, BasePath: "/api"})
	successfullyListQueriesMock := mock_queries.ClientService{}
	happyFlowClient.Queries = &successfullyListQueriesMock
	successfullyListQueriesMock.On("PostQueries", addParams, nil).Return(
		&gen_queries.PostQueriesOK{Payload: query}, nil)
	serverErrorClient := client.NewHTTPClientWithConfig(nil,
		&client.TransportConfig{Host: testAddress, BasePath: "/api"})
	serverErrQueryMock := mock_queries.ClientService{}
	serverErrorClient.Queries = &serverErrQueryMock
	serverErrQueryMock.On("PostQueries", addParams, nil).
		Return(nil, fmt.Errorf("Internal Server Error"))

	type fields struct {
		httpClient *client.Redashclient
	}
	type args struct {
		query *models.Query
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.Query
		wantErr bool
	}{
		{
			name:    "Successfully add a new query",
			fields:  fields{httpClient: happyFlowClient},
			args:    args{query: query},
			want:    query,
			wantErr: false,
		},
		{
			name:    "Add a new query fails on validation since query has no name",
			fields:  fields{httpClient: happyFlowClient},
			args:    args{query: &invalidQuery},
			want:    nil,
			wantErr: true,
		},

		{
			name:    "Add a new data source fails due to a server error",
			fields:  fields{httpClient: serverErrorClient},
			args:    args{query: query},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &RequestWrapper{
				httpClient: tt.fields.httpClient,
			}
			got, err := q.Add(tt.args.query)
			if (err != nil) != tt.wantErr {
				t.Errorf("RequestWrapper.Add() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RequestWrapper.Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRequestWrapper_Archive(t *testing.T) { //nolint:dupl // tests should be explicit
	archiveParams := gen_queries.NewDeleteQueriesIDParams().WithID(queryID)
	happyFlowClient := client.NewHTTPClientWithConfig(nil,
		&client.TransportConfig{Host: testAddress, BasePath: "/api"})
	successfullyListQueriesMock := mock_queries.ClientService{}
	happyFlowClient.Queries = &successfullyListQueriesMock
	successfullyListQueriesMock.On("DeleteQueriesID", archiveParams, nil).
		Return(nil, nil)
	serverErrorClient := client.NewHTTPClientWithConfig(nil,
		&client.TransportConfig{Host: testAddress, BasePath: "/api"})
	serverErrQueryMock := mock_queries.ClientService{}
	serverErrorClient.Queries = &serverErrQueryMock
	serverErrQueryMock.On("DeleteQueriesID", archiveParams, nil).
		Return(nil, fmt.Errorf("Internal Server Error"))

	type fields struct {
		httpClient *client.Redashclient
	}
	type args struct {
		queryID int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "Successfully archive an existing query",
			fields:  fields{httpClient: happyFlowClient},
			args:    args{queryID: queryID},
			wantErr: false,
		},
		{
			name:    "Archive query fails due to a server error",
			fields:  fields{httpClient: serverErrorClient},
			args:    args{queryID: queryID},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &RequestWrapper{
				httpClient: tt.fields.httpClient,
			}
			if err := q.Archive(tt.args.queryID); (err != nil) != tt.wantErr {
				t.Errorf("RequestWrapper.Archive() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRequestWrapper_Publish(t *testing.T) {
	query, _ := NewQuery(queryName, queryOptions, defaultDescription, queryContent, dataSourceID, id, defaultRepeated, defaultVersion, isDraft)
	publishedQuery := query
	publishedQuery.IsDraft = false
	publishParam := gen_queries.NewPostQueriesIDParams().WithID(queryID).WithBody(
		gen_queries.PostQueriesIDBody{
			ID:      id,
			IsDraft: &publishedQuery.IsDraft,
			Version: &defaultVersion,
		},
	)
	happyFlowClient := client.NewHTTPClientWithConfig(nil,
		&client.TransportConfig{Host: testAddress, BasePath: "/api"})
	successfulQueriesMock := mock_queries.ClientService{}
	happyFlowClient.Queries = &successfulQueriesMock
	successfulQueriesMock.On("PostQueriesID", publishParam, nil).
		Return(&gen_queries.PostQueriesIDOK{Payload: query}, nil)
	serverErrorClient := client.NewHTTPClientWithConfig(nil,
		&client.TransportConfig{Host: testAddress, BasePath: "/api"})
	serverErrQueryMock := mock_queries.ClientService{}
	serverErrorClient.Queries = &serverErrQueryMock
	serverErrQueryMock.On("PostQueriesID", publishParam, nil).
		Return(nil, fmt.Errorf("Internal Server Error"))

	type fields struct {
		httpClient *client.Redashclient
	}
	type args struct {
		queryID int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.Query
		wantErr bool
	}{
		{
			name:    "Successfully publish an existing query",
			fields:  fields{httpClient: happyFlowClient},
			args:    args{queryID: queryID},
			want:    publishedQuery,
			wantErr: false,
		},
		{
			name:    "Publish query fails due to a server error",
			fields:  fields{httpClient: serverErrorClient},
			args:    args{queryID: queryID},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &RequestWrapper{
				httpClient: tt.fields.httpClient,
			}
			got, err := q.Publish(tt.args.queryID)
			if (err != nil) != tt.wantErr {
				t.Errorf("RequestWrapper.Publish() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RequestWrapper.Publish() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRequestWrapper_Run(t *testing.T) {
	runQueryParams := gen_queries.NewPostQueriesIDResultsParams().WithID(queryID).WithBody(
		gen_queries.PostQueriesIDResultsBody{ID: nil, MaxAge: nil, ApplyAutoLimit: nil})

	happyFlowClient := client.NewHTTPClientWithConfig(nil,
		&client.TransportConfig{Host: testAddress, BasePath: "/api"})
	successfullyRunQueriesMock := mock_queries.ClientService{}
	happyFlowClient.Queries = &successfullyRunQueriesMock
	successfullyRunQueriesMock.On("PostQueriesIDResults", runQueryParams, nil).
		Return(nil, nil)
	serverErrorClient := client.NewHTTPClientWithConfig(nil,
		&client.TransportConfig{Host: testAddress, BasePath: "/api"})
	serverErrQueryMock := mock_queries.ClientService{}
	serverErrorClient.Queries = &serverErrQueryMock
	serverErrQueryMock.On("PostQueriesIDResults", runQueryParams, nil).
		Return(nil, fmt.Errorf("Internal Server Error"))

	type fields struct {
		httpClient *client.Redashclient
	}
	type args struct {
		queryID int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "Successfully run an existing query",
			fields:  fields{httpClient: happyFlowClient},
			args:    args{queryID: queryID},
			wantErr: false,
		},
		{
			name:    "Run query fails due to a server error",
			fields:  fields{httpClient: serverErrorClient},
			args:    args{queryID: queryID},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &RequestWrapper{
				httpClient: tt.fields.httpClient,
			}
			if err := q.ExecuteQuery(tt.args.queryID); (err != nil) != tt.wantErr {
				t.Errorf("RequestWrapper.Archive() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
