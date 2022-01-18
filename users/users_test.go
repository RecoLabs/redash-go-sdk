package users

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/go-openapi/strfmt"
	gen_users "github.com/recolabs/reco/redash-client/gen/client/users"
	mock_users "github.com/recolabs/reco/redash-client/mocks/users"

	"github.com/recolabs/reco/redash-client/gen/client"
	"github.com/recolabs/reco/redash-client/gen/models"
)

const (
	testAddress = "localhost:5005"
	userID      = 1
)

var (
	userName                 = "testUser"
	userOrgName              = "Corp"
	userEmail   strfmt.Email = "test@corp.com"
)

func TestNewUser(t *testing.T) {
	type args struct {
		name    string
		email   strfmt.Email
		orgName string
	}
	tests := []struct {
		name    string
		args    args
		want    *models.User
		wantErr bool
	}{
		{
			name: "Add a new data source successfully",
			args: args{
				name: userName, email: userEmail, orgName: userOrgName,
			},
			want: &models.User{
				Name:    &userName,
				Email:   &userEmail,
				OrgName: &userOrgName,
			},
			wantErr: false,
		},
		{
			name: "Fail to build a new User instance since the email is malformed",
			args: args{
				name: userName, email: "", orgName: userOrgName,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewUser(tt.args.name, tt.args.email, tt.args.orgName)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRequestWrapper_Get(t *testing.T) {
	getUserParams := gen_users.NewGetUsersIDParams()
	getUserParams.ID = userID
	user, _ := NewUser(userName, userEmail, userOrgName)
	happyFlowClient := client.NewHTTPClientWithConfig(nil,
		&client.TransportConfig{Host: testAddress, BasePath: "/api"})
	successfullyListUsersMock := mock_users.ClientService{}
	happyFlowClient.Users = &successfullyListUsersMock
	successfullyListUsersMock.On("GetUsersID", getUserParams, nil).
		Return(&gen_users.GetUsersIDOK{Payload: user}, nil)
	serverErrorClient := client.NewHTTPClientWithConfig(nil,
		&client.TransportConfig{Host: testAddress, BasePath: "/api"})
	serverErrUserMock := mock_users.ClientService{}
	serverErrorClient.Users = &serverErrUserMock
	serverErrUserMock.On("GetUsersID", getUserParams, nil).
		Return(nil, fmt.Errorf("Internal Server Error"))

	type fields struct {
		httpClient *client.Redashclient
	}
	type args struct {
		userID int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.User
		wantErr bool
	}{
		{
			name:    "Successfully get user",
			fields:  fields{httpClient: happyFlowClient},
			args:    args{userID: userID},
			want:    user,
			wantErr: false,
		},
		{
			name:    "Failed to get user due to a server error",
			fields:  fields{httpClient: serverErrorClient},
			args:    args{userID: userID},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &RequestWrapper{
				httpClient: tt.fields.httpClient,
			}
			got, err := u.Get(tt.args.userID)
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
