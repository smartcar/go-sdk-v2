// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/smartcar/go-sdk-v2/pkg/models/shared"
	"net/http"
)

type GetCadillacVoltageRequest struct {
	VehicleID string `pathParam:"style=simple,explode=false,name=vehicle_id"`
}

func (o *GetCadillacVoltageRequest) GetVehicleID() string {
	if o == nil {
		return ""
	}
	return o.VehicleID
}

type GetCadillacVoltageResponse struct {
	// returns the voltage of the charger measured by the vehicle.
	ChargeVoltage *shared.ChargeVoltage
	ContentType   string
	StatusCode    int
	RawResponse   *http.Response
}

func (o *GetCadillacVoltageResponse) GetChargeVoltage() *shared.ChargeVoltage {
	if o == nil {
		return nil
	}
	return o.ChargeVoltage
}

func (o *GetCadillacVoltageResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetCadillacVoltageResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetCadillacVoltageResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
