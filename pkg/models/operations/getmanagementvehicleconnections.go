// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/smartcar/go-sdk-v2/pkg/models/shared"
	"net/http"
)

var GetManagementVehicleConnectionsServerList = []string{
	"https://management.smartcar.com/v2.0",
}

type GetManagementVehicleConnectionsSecurity struct {
	Password string `security:"scheme,type=http,subtype=basic,name=password"`
	Username string `security:"scheme,type=http,subtype=basic,name=username"`
}

func (o *GetManagementVehicleConnectionsSecurity) GetPassword() string {
	if o == nil {
		return ""
	}
	return o.Password
}

func (o *GetManagementVehicleConnectionsSecurity) GetUsername() string {
	if o == nil {
		return ""
	}
	return o.Username
}

type GetManagementVehicleConnectionsRequest struct {
	// Used for accessing pages other than the first page. Each page returned has a cursor value that can be passed here to fetch the “next” page.
	Cursor *int64 `queryParam:"style=form,explode=true,name=cursor"`
	// Number of connections to return (default: 10, maximum: 100).
	Limit *int64 `queryParam:"style=form,explode=true,name=limit"`
	// Filter for connections created by the provider user ID.
	UserID *string `queryParam:"style=form,explode=true,name=user_id"`
	// Filter for connections to the provided vehicle ID.
	VehicleID *string `queryParam:"style=form,explode=true,name=vehicle_id"`
}

func (o *GetManagementVehicleConnectionsRequest) GetCursor() *int64 {
	if o == nil {
		return nil
	}
	return o.Cursor
}

func (o *GetManagementVehicleConnectionsRequest) GetLimit() *int64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *GetManagementVehicleConnectionsRequest) GetUserID() *string {
	if o == nil {
		return nil
	}
	return o.UserID
}

func (o *GetManagementVehicleConnectionsRequest) GetVehicleID() *string {
	if o == nil {
		return nil
	}
	return o.VehicleID
}

type GetManagementVehicleConnectionsResponse struct {
	// returns vehicle connections
	ConnectionsResponse *shared.ConnectionsResponse
	ContentType         string
	StatusCode          int
	RawResponse         *http.Response
}

func (o *GetManagementVehicleConnectionsResponse) GetConnectionsResponse() *shared.ConnectionsResponse {
	if o == nil {
		return nil
	}
	return o.ConnectionsResponse
}

func (o *GetManagementVehicleConnectionsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetManagementVehicleConnectionsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetManagementVehicleConnectionsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
