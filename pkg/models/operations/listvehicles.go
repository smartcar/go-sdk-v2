// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/smartcar/go-sdk-v2/pkg/models/shared"
	"net/http"
)

type ListVehiclesRequest struct {
	// Number of vehicles to return
	Limit *int64 `queryParam:"style=form,explode=true,name=limit"`
	// Index to start vehicle list at
	Offset *int64 `queryParam:"style=form,explode=true,name=offset"`
}

type ListVehiclesResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// A list of vehicles
	VehiclesResponse *shared.VehiclesResponse
}
