package visualizations

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/recolabs/redash-go-sdk/common/random"
	"github.com/recolabs/redash-go-sdk/gen/client"
	"github.com/recolabs/redash-go-sdk/gen/models"
	"github.com/recolabs/redash-go-sdk/options"
	gen_visualizations "github.com/recolabs/redash-go-sdk/gen/client/visualizations"
	mock_visualizations "github.com/recolabs/redash-go-sdk/mocks/visualizations"
)

const (
	testAddress         = "localhost:5005"
	malformedJSONString = `{
		"legend": {
			"enabled": true,
			"placement": "auto",
			"traceorder": "normal"
		},
		"xAxis": {
			"type": "-",
			"labels": {
				"enabled": true
			}
		},
		"yAxis": [
			{
				"type": "linear"
			},
			{
				"type": "linear",
				"opposite": true
			}
			}`
	visualizationOptions = `{
				"globalSeriesType": "column",
				"sortX": true,
				"legend": {
					"enabled": true,
					"placement": "auto",
					"traceorder": "normal"
				},
				"xAxis": {
					"type": "-",
					"labels": {
						"enabled": true
					}
				},
				"yAxis": [
					{
						"type": "linear"
					},
					{
						"type": "linear",
						"opposite": true
					}
					],
					"alignYAxesAtZero": false,
					"error_y": {
						"type": "data",
						"visible": true
					},
					"series": {
						"stacking": null,
						"error_y": {
							"type": "data",
							"visible": true
						}
					},
					"seriesOptions": {},
					"valuesOptions": {},
					"columnMapping": {
						"user_email": "x",
						"downloaded_files": "y"
					},
					"direction": {
						"type": "counterclockwise"
					},
					"sizemode": "diameter",
					"coefficient": 1,
					"numberFormat": "0,0[.]00000",
					"percentFormat": "0[.]00%",
					"textFormat": "",
					"missingValuesAsZero": true,
					"showDataLabels": false,
					"dateTimeFormat": "DD/MM/YY HH:mm"
					}`
	queryID         = 116101115116 // ord concatenation of "test"
	visualizationID = int64(1)
	description     = "test visualization"
)

var (
	queryID64                = int64(queryID)
	visualizationsType       = models.VisualizationTypeCHART
	visualizationsName       = models.VisualizationTypeCHART
	testToken                = random.WeakRandString(100)
	optionsMap, _            = options.MapFromString(visualizationOptions)
	invalidVisualizationType = "MAPA"
)

func TestNewVisualization(t *testing.T) {
	type args struct {
		visualizationType string
		name              string
		optionsJSON       string
		description       string
		queryID           int64
	}
	tests := []struct {
		name    string
		args    args
		want    *models.Visualization
		wantErr bool
	}{
		{
			name: "Create a new visualization successfully",
			args: args{visualizationType: visualizationsType,
				name: visualizationsName, optionsJSON: visualizationOptions,
				description: description, queryID: queryID,
			},
			want: &models.Visualization{
				Type:        &visualizationsType,
				Name:        &visualizationsName,
				Options:     optionsMap,
				Description: description,
				QueryID:     &queryID64},
			wantErr: false,
		},
		{
			name: "Create a new visualization with an invalid type",
			args: args{visualizationType: invalidVisualizationType,
				name: visualizationsName, optionsJSON: visualizationOptions,
				description: description, queryID: queryID,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Create a new visualization without a valid json options string",
			args: args{visualizationType: visualizationsType,
				name: visualizationsName, optionsJSON: malformedJSONString,
				description: description, queryID: queryID,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewVisualization(tt.args.visualizationType, tt.args.name, tt.args.optionsJSON, tt.args.description, tt.args.queryID)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewVisualization() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewVisualization() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRequestWrapper_GetURL(t *testing.T) {
	httpClient := client.NewHTTPClientWithConfig(nil,
		&client.TransportConfig{Host: testAddress, BasePath: "/api"})
	type fields struct {
		httpClient *client.Redashclient
		host       string
		scheme     string
	}
	type args struct {
		visualizationID int64
		queryID         int64
		QueryAPIKey     string
		hideParameters  bool
		hideHeader      bool
		hideLink        bool
		hideTimestamp   bool
	}
	tests := []struct {
		name                 string
		fields               fields
		args                 args
		wantVisualizationURL string
	}{
		{
			name:   "GetURL sanity check",
			fields: fields{httpClient: httpClient, host: testAddress, scheme: "http"},
			args: args{visualizationID: visualizationID, queryID: queryID, QueryAPIKey: testToken,
				hideParameters: false, hideHeader: false, hideLink: false, hideTimestamp: false},
			wantVisualizationURL: fmt.Sprintf(
				"http://%s%s?api_key=%s", testAddress, fmt.Sprintf(EmbedQueryURIFormat, queryID, visualizationID), testToken),
		},
		{
			name:   "GetURL hide objects",
			fields: fields{httpClient: httpClient, host: testAddress, scheme: "http"},
			args: args{visualizationID: visualizationID, queryID: queryID, QueryAPIKey: testToken,
				hideParameters: true, hideHeader: true, hideLink: true, hideTimestamp: true},
			wantVisualizationURL: fmt.Sprintf(
				"http://%s%s?api_key=%s&hide_parameters&hide_header&hide_link&hide_timestamp",
				testAddress, fmt.Sprintf(EmbedQueryURIFormat, queryID, visualizationID), testToken),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &RequestWrapper{
				httpClient: tt.fields.httpClient,
				host:       tt.fields.host,
				scheme:     tt.fields.scheme,
				apiKey:     testToken,
			}
			gotVisualizationURL := v.GetURL(tt.args.visualizationID, tt.args.queryID, tt.args.hideParameters,
				tt.args.hideHeader, tt.args.hideLink, tt.args.hideTimestamp)
			if gotVisualizationURL != tt.wantVisualizationURL {
				t.Errorf("RequestWrapper.GetURL() = %v, want %v", gotVisualizationURL, tt.wantVisualizationURL)
			}
		})
	}
}

func TestRequestWrapper_Add(t *testing.T) {
	invalidVis := &models.Visualization{
		Type:        &invalidVisualizationType,
		Name:        &visualizationsName,
		Options:     visualizationOptions,
		Description: description,
		QueryID:     &queryID64}

	vis, _ := NewVisualization(visualizationsType, visualizationsName, visualizationOptions, description, queryID)
	addParams := gen_visualizations.NewPostVisualizationsParams().WithBody(gen_visualizations.PostVisualizationsBody{
		Name:    *vis.Name,
		Type:    *vis.Type,
		Options: vis.Options,
		QueryID: *vis.QueryID,
	})

	happyFlowClient := client.NewHTTPClientWithConfig(nil,
		&client.TransportConfig{Host: testAddress, BasePath: "/api"})
	successfullyAddVisualizationMock := mock_visualizations.ClientService{}
	happyFlowClient.Visualizations = &successfullyAddVisualizationMock
	successfullyAddVisualizationMock.On("PostVisualizations", addParams, nil).Return(
		&gen_visualizations.PostVisualizationsOK{Payload: vis}, nil)
	serverErrorClient := client.NewHTTPClientWithConfig(nil,
		&client.TransportConfig{Host: testAddress, BasePath: "/api"})
	serverErrVisualizationMock := mock_visualizations.ClientService{}
	serverErrorClient.Visualizations = &serverErrVisualizationMock
	serverErrVisualizationMock.On("PostVisualizations", addParams, nil).
		Return(nil, fmt.Errorf("Internal Server Error"))

	type fields struct {
		httpClient *client.Redashclient
		host       string
	}
	type args struct {
		visualization *models.Visualization
	}
	tests := []struct {
		name                    string
		fields                  fields
		args                    args
		wantResultVisualization *models.Visualization
		wantErr                 bool
	}{
		{
			name:                    "Successfully add a new visualization",
			fields:                  fields{httpClient: happyFlowClient, host: testAddress},
			args:                    args{visualization: vis},
			wantResultVisualization: vis,
			wantErr:                 false,
		},
		{
			name:                    "Add a new visualization with an invalid visualization type",
			fields:                  fields{httpClient: happyFlowClient, host: testAddress},
			args:                    args{visualization: invalidVis},
			wantResultVisualization: nil,
			wantErr:                 true,
		},
		{
			name:                    "Add a new visualization fails on a server error",
			fields:                  fields{httpClient: serverErrorClient, host: testAddress},
			args:                    args{visualization: vis},
			wantResultVisualization: nil,
			wantErr:                 true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &RequestWrapper{
				httpClient: tt.fields.httpClient,
				host:       tt.fields.host,
			}
			gotResultVisualization, err := v.Add(tt.args.visualization)
			if (err != nil) != tt.wantErr {
				t.Errorf("RequestWrapper.Add() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResultVisualization, tt.wantResultVisualization) {
				t.Errorf("RequestWrapper.Add() = %v, want %v", gotResultVisualization, tt.wantResultVisualization)
			}
		})
	}
}

func TestRequestWrapper_Delete(t *testing.T) {
	deleteParams := gen_visualizations.NewDeleteVisualizationsIDParams().WithID(visualizationID)
	happyFlowClient := client.NewHTTPClientWithConfig(nil,
		&client.TransportConfig{Host: testAddress, BasePath: "/api"})
	successfullyDeleteVisualizationMock := mock_visualizations.ClientService{}
	happyFlowClient.Visualizations = &successfullyDeleteVisualizationMock
	successfullyDeleteVisualizationMock.On("DeleteVisualizationsID", deleteParams, nil).Return(
		&gen_visualizations.DeleteVisualizationsIDOK{}, nil)
	serverErrorClient := client.NewHTTPClientWithConfig(nil,
		&client.TransportConfig{Host: testAddress, BasePath: "/api"})
	serverErrVisualizationMock := mock_visualizations.ClientService{}
	serverErrorClient.Visualizations = &serverErrVisualizationMock
	serverErrVisualizationMock.On("DeleteVisualizationsID", deleteParams, nil).
		Return(nil, fmt.Errorf("Internal Server Error"))

	type fields struct {
		httpClient *client.Redashclient
		host       string
	}
	type args struct {
		visualizationID int64
		QueryAPIKey     string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "Successfully archive an existing visualization",
			fields:  fields{httpClient: happyFlowClient, host: testAddress},
			args:    args{visualizationID: visualizationID, QueryAPIKey: testToken},
			wantErr: false,
		},
		{
			name:    "Archive visualization fails on a server error",
			fields:  fields{httpClient: serverErrorClient, host: testAddress},
			args:    args{visualizationID: visualizationID, QueryAPIKey: testToken},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &RequestWrapper{
				httpClient: tt.fields.httpClient,
				host:       tt.fields.host,
			}
			if err := v.Delete(tt.args.visualizationID); (err != nil) != tt.wantErr {
				t.Errorf("RequestWrapper.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDataToVisualization(t *testing.T) {
	type input struct {
		visString string
	}

	type expectedOutput struct {
		description string
		ID          int64
		err         error
	}

	tests := []struct {
		name           string
		input          input
		expectedOutput expectedOutput
	}{
		{name: "Test Counter visualization",
			input: input{visString: "{\n  \"id\": 33,\n  \"type\": \"COUNTER\",\n  \"name\": \"Value Created (annual $)\",\n " +
				" \"description\": \"desc\",\n  \"options\": {\n    \"counterLabel\": \"\",\n    \"counterColName\": \"_col0\",\n   " +
				" \"rowNumber\": 1,\n    \"targetRowNumber\": 1,\n    \"stringDecimal\": 0,\n    \"stringDecChar\": \".\",\n  " +
				" \"stringThouSep\": \",\",\n    \"tooltipFormat\": \"0,0.000\"\n  },\n  \"updated_at\": \"2021-12-28T08:54:00.685Z\",\n " +
				" \"created_at\": \"2021-12-28T08:35:33.978Z\"\n}"},
			expectedOutput: expectedOutput{err: nil, description: "desc", ID: 33},
		},
		{name: "Test Table visualization",
			input: input{visString: "{\n  \"id\": 31,\n  \"type\": \"TABLE\",\n " +
				"\"name\": \"Table\",\n  \"description\": \"\",\n  \"options\": {},\n " +
				" \"updated_at\": \"2021-12-28T08:31:03.049Z\",\n  \"created_at\": \"2021-12-28T08:31:03.049Z\"\n}"},
			expectedOutput: expectedOutput{err: nil, description: "", ID: 31},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			vis, err := ToVisualization(test.input.visString)
			if err != nil {
				t.Fatalf("Error while casting to visualization")
			}
			if vis.Description != test.expectedOutput.description {
				t.Errorf("expected description: %v, got: %v", test.expectedOutput.description, vis.Description)
			}
			if vis.ID != test.expectedOutput.ID {
				t.Errorf("expected description: %v, got: %v", test.expectedOutput.ID, vis.ID)
			}
			if vis.Options == nil {
				t.Errorf("options is nil %v", vis.Options)
			}
		})
	}
}
