// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/smartcar/go-sdk-v2/pkg/models/shared"
	"net/http"
)

type GetChevroletChargeTimeRequest struct {
	VehicleID string `pathParam:"style=simple,explode=false,name=vehicle_id"`
}

func (o *GetChevroletChargeTimeRequest) GetVehicleID() string {
	if o == nil {
		return ""
	}
	return o.VehicleID
}

type GetChevroletChargeTimeResponse struct {
	// returns the date and time the vehicle expects to "complete" this charging session.
	ChargeTime  *shared.ChargeTime
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}

func (o *GetChevroletChargeTimeResponse) GetChargeTime() *shared.ChargeTime {
	if o == nil {
		return nil
	}
	return o.ChargeTime
}

func (o *GetChevroletChargeTimeResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetChevroletChargeTimeResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetChevroletChargeTimeResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
