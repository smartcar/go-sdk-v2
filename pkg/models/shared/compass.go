// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// Direction - The direction the vehicle is traveling.
type Direction string

const (
	DirectionN  Direction = "N"
	DirectionNe Direction = "NE"
	DirectionE  Direction = "E"
	DirectionSe Direction = "SE"
	DirectionS  Direction = "S"
	DirectionSw Direction = "SW"
	DirectionW  Direction = "W"
	DirectionNw Direction = "NW"
)

func (e Direction) ToPointer() *Direction {
	return &e
}

func (e *Direction) UnmarshalJSON(data []byte) error {
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
		*e = Direction(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Direction: %v", v)
	}
}

type Compass struct {
	// The direction the vehicle is traveling.
	Direction *Direction `json:"direction,omitempty"`
	// The direction the vehicle is traveling (in degrees).
	Heading *float32 `json:"heading,omitempty"`
}

func (o *Compass) GetDirection() *Direction {
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
