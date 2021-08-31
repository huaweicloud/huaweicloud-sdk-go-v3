package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type RecognizeVehicleLicenseRequest struct {
	Body *VehicleLicenseRequestBody `json:"body,omitempty"`
}

func (o RecognizeVehicleLicenseRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "RecognizeVehicleLicenseRequest struct{}"
	}

	return strings.Join([]string{"RecognizeVehicleLicenseRequest", string(data)}, " ")
}
