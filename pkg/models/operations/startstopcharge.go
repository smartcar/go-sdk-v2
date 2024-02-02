// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/smartcar/go-sdk-v2/v3/pkg/models/shared"
	"net/http"
)

type StartStopChargeRequest struct {
	ChargeAction *shared.ChargeAction `request:"mediaType=application/json"`
	VehicleID    string               `pathParam:"style=simple,explode=false,name=vehicle_id"`
}

func (o *StartStopChargeRequest) GetChargeAction() *shared.ChargeAction {
	if o == nil {
		return nil
	}
	return o.ChargeAction
}

func (o *StartStopChargeRequest) GetVehicleID() string {
	if o == nil {
		return ""
	}
	return o.VehicleID
}

type StartStopChargeResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// return Success Response
	SuccessResponse *shared.SuccessResponse
}

func (o *StartStopChargeResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *StartStopChargeResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *StartStopChargeResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *StartStopChargeResponse) GetSuccessResponse() *shared.SuccessResponse {
	if o == nil {
		return nil
	}
	return o.SuccessResponse
}
