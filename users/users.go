package users

import (
	"github.com/go-openapi/strfmt"
	"github.com/recolabs/reco/redash-client/gen/client"
	"github.com/recolabs/reco/redash-client/gen/client/users"
	"github.com/recolabs/reco/redash-client/gen/models"
)

type RequestWrapper struct {
	httpClient *client.Redashclient
	opts       []users.ClientOption
}

func NewRequestWrapper(httpClient *client.Redashclient, opts ...users.ClientOption) *RequestWrapper {
	return &RequestWrapper{httpClient: httpClient, opts: opts}
}

// NewUser builds a User object. This method offers no validation,
// it exists for compatibility with the other models
func NewUser(name string, email strfmt.Email, orgName string) (*models.User, error) {
	user := &models.User{
		Name:    &name,
		Email:   &email,
		OrgName: &orgName,
	}

	err := user.Validate(nil)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// GetUser gets the details admin's user details
func (requestWrapper *RequestWrapper) Get(userID int64) (*models.User, error) {
	getUserParams := users.NewGetUsersIDParams()
	getUserParams.ID = userID
	response, err := requestWrapper.httpClient.Users.GetUsersID(getUserParams, nil, requestWrapper.opts...)
	if err != nil {
		return nil, err
	}

	return response.GetPayload(), nil
}
