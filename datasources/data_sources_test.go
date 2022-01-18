// Package datasources provides methods of adding, removing and listing Data Sources
package datasources

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/recolabs/redash-go-sdk/gen/client"
	"github.com/recolabs/redash-go-sdk/gen/models"
	"github.com/recolabs/redash-go-sdk/options"
	gen_data_sources "github.com/recolabs/redash-go-sdk/gen/client/data_sources"
	mock_data_sources "github.com/recolabs/redash-go-sdk/mocks/data_sources"
)

const (
	testAddress       = "localhost:5005"
	postgresqlOptions = `{
		"dbname": "aa",
		"host": "1.1.1.1",
		"port": 5432
	}`
	malformedPostgresqlOptions = `{
		"dbname": "aa",
		"host": "1.1.1.1",
		"port": 5432,
	}`

	dataSourceID = int64(1)
)

var (
	dataSourceType        = models.DataSourceTypeGoogleSpreadsheets
	invalidDataSourceType = "Test Test"
	dataSourceName        = "Google Sheets"
	optionsMap, _         = options.MapFromString(postgresqlOptions)
)

func TestNewDataSource(t *testing.T) {
	type args struct {
		dataSourceType string
		name           string
		optionsJSON    string
	}
	tests := []struct {
		name    string
		args    args
		want    *models.DataSource
		wantErr bool
	}{
		{
			name: "Create a new data source successfully",
			args: args{dataSourceType: dataSourceType,
				name: dataSourceName, optionsJSON: postgresqlOptions,
			},
			want: &models.DataSource{
				Type:    &dataSourceType,
				Name:    &dataSourceName,
				Options: optionsMap,
			},
			wantErr: false,
		},
		{
			name: "Create a new data source without with an invalid type string",
			args: args{
				dataSourceType: invalidDataSourceType,
				name:           dataSourceName,
				optionsJSON:    malformedPostgresqlOptions,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Create a new data source without a valid json options string",
			args: args{
				dataSourceType: dataSourceType,
				name:           dataSourceName,
				optionsJSON:    malformedPostgresqlOptions,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewDataSource(tt.args.dataSourceType, tt.args.name, tt.args.optionsJSON)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewDataSource() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDataSource() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRequestWrapper_Add(t *testing.T) {
	invalidDs := &models.DataSource{
		Type:    &invalidDataSourceType,
		Name:    &dataSourceName,
		Options: postgresqlOptions,
	}
	ds, _ := NewDataSource(dataSourceType, dataSourceName, postgresqlOptions)
	addParams := gen_data_sources.NewPostDataSourcesParams().WithBody(gen_data_sources.PostDataSourcesBody{
		Name:    *ds.Name,
		Type:    *ds.Type,
		Options: ds.Options,
	})

	happyFlowClient := client.NewHTTPClientWithConfig(nil,
		&client.TransportConfig{Host: testAddress, BasePath: "/api"})
	successfullyAddDataSourceMock := mock_data_sources.ClientService{}
	happyFlowClient.DataSources = &successfullyAddDataSourceMock
	successfullyAddDataSourceMock.On("PostDataSources", addParams, nil).Return(
		&gen_data_sources.PostDataSourcesOK{Payload: ds}, nil)
	serverErrorClient := client.NewHTTPClientWithConfig(nil,
		&client.TransportConfig{Host: testAddress, BasePath: "/api"})
	serverErrDataSourceMock := mock_data_sources.ClientService{}
	serverErrorClient.DataSources = &serverErrDataSourceMock
	serverErrDataSourceMock.On("PostDataSources", addParams, nil).
		Return(nil, fmt.Errorf("Internal Server Error"))

	type fields struct {
		httpClient *client.Redashclient
	}
	type args struct {
		dataSource *models.DataSource
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.DataSource
		wantErr bool
	}{
		{
			name:    "Successfully add a new data source",
			fields:  fields{httpClient: happyFlowClient},
			args:    args{dataSource: ds},
			want:    ds,
			wantErr: false,
		},
		{
			name: `Add a new data source fails due to a validation error since data
					    source object containing an invalid type string`,
			fields:  fields{httpClient: serverErrorClient},
			args:    args{dataSource: invalidDs},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Add a new data source fails due to a server error",
			fields:  fields{httpClient: serverErrorClient},
			args:    args{dataSource: ds},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &RequestWrapper{
				httpClient: tt.fields.httpClient,
			}
			got, err := d.Add(tt.args.dataSource)
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

func TestRequestWrapper_List(t *testing.T) {
	ds, _ := NewDataSource(dataSourceType, dataSourceName, postgresqlOptions)
	happyFlowClient := client.NewHTTPClientWithConfig(nil,
		&client.TransportConfig{Host: testAddress, BasePath: "/api"})
	successfullyListDataSourceMock := mock_data_sources.ClientService{}
	happyFlowClient.DataSources = &successfullyListDataSourceMock
	successfullyListDataSourceMock.On("List", (*gen_data_sources.ListParams)(nil), nil).Return(
		&gen_data_sources.ListOK{Payload: []*models.DataSource{ds}}, nil)
	serverErrorClient := client.NewHTTPClientWithConfig(nil,
		&client.TransportConfig{Host: testAddress, BasePath: "/api"})
	serverErrDataSourceMock := mock_data_sources.ClientService{}
	serverErrorClient.DataSources = &serverErrDataSourceMock
	serverErrDataSourceMock.On("List", (*gen_data_sources.ListParams)(nil), nil).
		Return(nil, fmt.Errorf("Internal Server Error"))

	type fields struct {
		httpClient *client.Redashclient
	}
	tests := []struct {
		name                 string
		fields               fields
		wantResultDataSource []*models.DataSource
		wantErr              bool
	}{
		{
			name:                 "Successfully list data sources",
			fields:               fields{httpClient: happyFlowClient},
			wantResultDataSource: []*models.DataSource{ds},
			wantErr:              false,
		},
		{
			name:                 "Listing data sources fails due to a server error",
			fields:               fields{httpClient: serverErrorClient},
			wantResultDataSource: nil,
			wantErr:              true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &RequestWrapper{
				httpClient: tt.fields.httpClient,
			}
			gotResultDataSource, err := d.List()
			if (err != nil) != tt.wantErr {
				t.Errorf("RequestWrapper.List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResultDataSource, tt.wantResultDataSource) {
				t.Errorf("RequestWrapper.List() = %v, want %v", gotResultDataSource, tt.wantResultDataSource)
			}
		})
	}
}

func TestRequestWrapper_Delete(t *testing.T) {
	deleteParams := gen_data_sources.NewDeleteDataSourcesIDParams().WithID(dataSourceID)

	happyFlowClient := client.NewHTTPClientWithConfig(nil,
		&client.TransportConfig{Host: testAddress, BasePath: "/api"})
	successfullyDeleteDataSourceMock := mock_data_sources.ClientService{}
	happyFlowClient.DataSources = &successfullyDeleteDataSourceMock
	successfullyDeleteDataSourceMock.On("DeleteDataSourcesID", deleteParams, nil).Return(
		&gen_data_sources.DeleteDataSourcesIDNoContent{}, nil)
	serverErrorClient := client.NewHTTPClientWithConfig(nil,
		&client.TransportConfig{Host: testAddress, BasePath: "/api"})
	serverErrDataSourceMock := mock_data_sources.ClientService{}
	serverErrorClient.DataSources = &serverErrDataSourceMock
	serverErrDataSourceMock.On("DeleteDataSourcesID", deleteParams, nil).
		Return(nil, fmt.Errorf("Internal Server Error"))

	type fields struct {
		httpClient *client.Redashclient
	}
	type args struct {
		dataSourceID int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "Successfully removed an existing data source",
			fields:  fields{httpClient: happyFlowClient},
			args:    args{dataSourceID: dataSourceID},
			wantErr: false,
		},
		{
			name:    "Data source deletion fails due to a server error",
			fields:  fields{httpClient: serverErrorClient},
			args:    args{dataSourceID: dataSourceID},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &RequestWrapper{
				httpClient: tt.fields.httpClient,
			}
			if err := d.Delete(tt.args.dataSourceID); (err != nil) != tt.wantErr {
				t.Errorf("RequestWrapper.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
