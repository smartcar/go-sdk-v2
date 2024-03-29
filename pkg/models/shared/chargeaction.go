// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type ChargeActionAction string

const (
	ChargeActionActionStart ChargeActionAction = "START"
	ChargeActionActionStop  ChargeActionAction = "STOP"
)

func (e ChargeActionAction) ToPointer() *ChargeActionAction {
	return &e
}

func (e *ChargeActionAction) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "START":
		fallthrough
	case "STOP":
		*e = ChargeActionAction(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ChargeActionAction: %v", v)
	}
}

type ChargeAction struct {
	Action *ChargeActionAction `json:"action,omitempty"`
}

func (o *ChargeAction) GetAction() *ChargeActionAction {
	if o == nil {
		return nil
	}
	return o.Action
}
