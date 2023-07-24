// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// CompassDirection - The direction the vehicle is traveling.
type CompassDirection string

const (
	CompassDirectionN  CompassDirection = "N"
	CompassDirectionNe CompassDirection = "NE"
	CompassDirectionE  CompassDirection = "E"
	CompassDirectionSe CompassDirection = "SE"
	CompassDirectionS  CompassDirection = "S"
	CompassDirectionSw CompassDirection = "SW"
	CompassDirectionW  CompassDirection = "W"
	CompassDirectionNw CompassDirection = "NW"
)

func (e CompassDirection) ToPointer() *CompassDirection {
	return &e
}

func (e *CompassDirection) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "N":
		fallthrough
	case "NE":
		fallthrough
	case "E":
		fallthrough
	case "SE":
		fallthrough
	case "S":
		fallthrough
	case "SW":
		fallthrough
	case "W":
		fallthrough
	case "NW":
		*e = CompassDirection(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CompassDirection: %v", v)
	}
}

// Compass - returns the compass heading of a Tesla.
type Compass struct {
	// The direction the vehicle is traveling.
	Direction *CompassDirection `json:"direction,omitempty"`
	// The direction the vehicle is traveling (in degrees).
	Heading *float32 `json:"heading,omitempty"`
}

func (o *Compass) GetDirection() *CompassDirection {
	if o == nil {
		return nil
	}
	return o.Direction
}

func (o *Compass) GetHeading() *float32 {
	if o == nil {
		return nil
	}
	return o.Heading
}
