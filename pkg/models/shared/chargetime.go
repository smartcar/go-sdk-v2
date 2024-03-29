// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/smartcar/go-sdk-v2/pkg/utils"
	"time"
)

type ChargeTime struct {
	// The date and time the vehicle expects to complete this charging session.
	Time *time.Time `json:"time,omitempty"`
}

func (c ChargeTime) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *ChargeTime) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *ChargeTime) GetTime() *time.Time {
	if o == nil {
		return nil
	}
	return o.Time
}
