// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// Temperature - returns the exterior temperature of a Tesla.
type Temperature struct {
	Temperature *float32 `json:"temperature,omitempty"`
}

func (o *Temperature) GetTemperature() *float32 {
	if o == nil {
		return nil
	}
	return o.Temperature
}
