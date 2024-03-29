// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type EngineOil struct {
	// The engine oil’s remaining life span (as a percentage). Oil life is based on the current quality of the oil. (in percent).
	LifeRemaining *float32 `json:"lifeRemaining,omitempty"`
}

func (o *EngineOil) GetLifeRemaining() *float32 {
	if o == nil {
		return nil
	}
	return o.LifeRemaining
}
