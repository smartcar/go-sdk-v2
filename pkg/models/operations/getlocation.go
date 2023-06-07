// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Smartcar-API/pkg/models/shared"
	"net/http"
)

type GetLocationRequest struct {
	VehicleID string `pathParam:"style=simple,explode=false,name=vehicle_id"`
}

type GetLocationResponse struct {
	ContentType string
	// A list of vehicles
	Location    *shared.Location
	StatusCode  int
	RawResponse *http.Response
}
