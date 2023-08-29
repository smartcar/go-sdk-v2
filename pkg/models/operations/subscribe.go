// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/smartcar/go-sdk-v2/pkg/models/shared"
	"net/http"
)

type SubscribeRequest struct {
	WebhookInfo *shared.WebhookInfo `request:"mediaType=application/json"`
	VehicleID   string              `pathParam:"style=simple,explode=false,name=vehicle_id"`
	WebhookID   string              `pathParam:"style=simple,explode=false,name=webhookId"`
}

func (o *SubscribeRequest) GetWebhookInfo() *shared.WebhookInfo {
	if o == nil {
		return nil
	}
	return o.WebhookInfo
}

func (o *SubscribeRequest) GetVehicleID() string {
	if o == nil {
		return ""
	}
	return o.VehicleID
}

func (o *SubscribeRequest) GetWebhookID() string {
	if o == nil {
		return ""
	}
	return o.WebhookID
}

type SubscribeResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// return Compatibility
	SuccessResponse *shared.SuccessResponse
}

func (o *SubscribeResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *SubscribeResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *SubscribeResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *SubscribeResponse) GetSuccessResponse() *shared.SuccessResponse {
	if o == nil {
		return nil
	}
	return o.SuccessResponse
}
