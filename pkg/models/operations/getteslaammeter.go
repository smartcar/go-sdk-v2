// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/smartcar/go-sdk-v2/pkg/models/shared"
	"net/http"
)

type GetTeslaAmmeterRequest struct {
	VehicleID string `pathParam:"style=simple,explode=false,name=vehicle_id"`
}

type GetTeslaAmmeterResponse struct {
	// returns the amperage of the charger measured by the vehicle.
	ChargeAmmeter *shared.ChargeAmmeter
	ContentType   string
	StatusCode    int
	RawResponse   *http.Response
}