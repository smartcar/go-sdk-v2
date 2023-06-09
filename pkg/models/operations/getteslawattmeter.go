// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/smartcar/go-sdk-v2/pkg/models/shared"
	"net/http"
)

type GetTeslaWattmeterRequest struct {
	VehicleID string `pathParam:"style=simple,explode=false,name=vehicle_id"`
}

type GetTeslaWattmeterResponse struct {
	// returns the wattage of the charger measured by the vehicle.
	ChargeWattage *shared.ChargeWattage
	ContentType   string
	StatusCode    int
	RawResponse   *http.Response
}
