// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/smartcar/go-sdk-v2/v3/pkg/models/shared"
	"net/http"
)

type GetPermissionsRequest struct {
	// Number of vehicles to return
	Limit *int64 `queryParam:"style=form,explode=true,name=limit"`
	// Index to start vehicle list at
	Offset    *int64 `queryParam:"style=form,explode=true,name=offset"`
	VehicleID string `pathParam:"style=simple,explode=false,name=vehicle_id"`
}

func (o *GetPermissionsRequest) GetLimit() *int64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *GetPermissionsRequest) GetOffset() *int64 {
	if o == nil {
		return nil
	}
	return o.Offset
}

func (o *GetPermissionsRequest) GetVehicleID() string {
	if o == nil {
		return ""
	}
	return o.VehicleID
}

type GetPermissionsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// The applications permissions
	Permission *shared.Permission
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetPermissionsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetPermissionsResponse) GetPermission() *shared.Permission {
	if o == nil {
		return nil
	}
	return o.Permission
}

func (o *GetPermissionsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetPermissionsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
