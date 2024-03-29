// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type Location struct {
	Latitude  float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
}

func (o *Location) GetLatitude() float32 {
	if o == nil {
		return 0.0
	}
	return o.Latitude
}

func (o *Location) GetLongitude() float32 {
	if o == nil {
		return 0.0
	}
	return o.Longitude
}
