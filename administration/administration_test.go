package administration

import (
	"fmt"
	"testing"

	"github.com/recolabs/reco/redash-client/gen/client"
	admin "github.com/recolabs/reco/redash-client/gen/client/administration"
	mock_administration "github.com/recolabs/reco/redash-client/mocks/administration"
)

const (
	testAddress = "localhost:5005"
)

func TestPingFunction(t *testing.T) {
	pingParams := admin.NewGetPingParams()
	happyFlowClient := client.NewHTTPClientWithConfig(nil,
		&client.TransportConfig{Host: testAddress, BasePath: "/api"})

	successfullyAdministrationMock := mock_administration.ClientService{}
	happyFlowClient.Administration = &successfullyAdministrationMock
	successfullyAdministrationMock.On("GetPing", pingParams, nil).Return(admin.NewGetPingOK(), nil)

	serverErrorClient := client.NewHTTPClientWithConfig(nil,
		&client.TransportConfig{Host: testAddress, BasePath: "/api"})
	serverErrAdministrationMock := mock_administration.ClientService{}
	serverErrorClient.Administration = &serverErrAdministrationMock
	serverErrAdministrationMock.On("GetPing", pingParams, nil).
		Return(admin.NewGetPingOK(), fmt.Errorf("Internal Server Error"))

	type fields struct {
		httpClient *client.Redashclient
		host       string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name:    "Successfully ping",
			fields:  fields{httpClient: happyFlowClient, host: testAddress},
			wantErr: false,
		},
		{
			name:    "Ping fails on a server error",
			fields:  fields{httpClient: serverErrorClient, host: testAddress},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			administration := &RequestWrapper{
				httpClient: tt.fields.httpClient,
				host:       tt.fields.host,
			}
			err := administration.Ping()
			if (err != nil) != tt.wantErr {
				t.Errorf("RequestWrapper.Add() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
