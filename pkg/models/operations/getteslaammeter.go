// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/smartcar/go-sdk-v2/v3/pkg/models/shared"
	"net/http"
)

type GetTeslaAmmeterRequest struct {
	VehicleID string `pathParam:"style=simple,explode=false,name=vehicle_id"`
}

func (o *GetTeslaAmmeterRequest) GetVehicleID() string {
	if o == nil {
		return ""
	}
	return o.VehicleID
}

type GetTeslaAmmeterResponse struct {
	// returns the amperage of the charger measured by the vehicle.
	ChargeAmmeter *shared.ChargeAmmeter
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetTeslaAmmeterResponse) GetChargeAmmeter() *shared.ChargeAmmeter {
	if o == nil {
		return nil
	}
	return o.ChargeAmmeter
}

func (o *GetTeslaAmmeterResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetTeslaAmmeterResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetTeslaAmmeterResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
