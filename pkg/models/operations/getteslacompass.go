// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/smartcar/go-sdk-v2/v3/pkg/models/shared"
	"net/http"
)

type GetTeslaCompassRequest struct {
	VehicleID string `pathParam:"style=simple,explode=false,name=vehicle_id"`
}

func (o *GetTeslaCompassRequest) GetVehicleID() string {
	if o == nil {
		return ""
	}
	return o.VehicleID
}

type GetTeslaCompassResponse struct {
	// returns the compass heading of a Tesla.
	Compass *shared.Compass
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetTeslaCompassResponse) GetCompass() *shared.Compass {
	if o == nil {
		return nil
	}
	return o.Compass
}

func (o *GetTeslaCompassResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetTeslaCompassResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetTeslaCompassResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
