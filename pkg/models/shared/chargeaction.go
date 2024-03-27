// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type Action string

const (
	ActionStart Action = "START"
	ActionStop  Action = "STOP"
)

func (e Action) ToPointer() *Action {
	return &e
}

func (e *Action) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "START":
		fallthrough
	case "STOP":
		*e = Action(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Action: %v", v)
	}
}

type ChargeAction struct {
	Action *Action `json:"action,omitempty"`
}

func (o *ChargeAction) GetAction() *Action {
	if o == nil {
		return nil
	}
	return o.Action
}
