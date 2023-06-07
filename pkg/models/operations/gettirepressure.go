// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Smartcar-API/pkg/models/shared"
	"net/http"
)

type GetTirePressureRequest struct {
	VehicleID string `pathParam:"style=simple,explode=false,name=vehicle_id"`
}

type GetTirePressureResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// return Pressure reading
	TirePressure *shared.TirePressure
}
