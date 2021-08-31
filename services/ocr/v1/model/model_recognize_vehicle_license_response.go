package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type RecognizeVehicleLicenseResponse struct {
	Result         *VehicleLicenseResult `json:"result,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o RecognizeVehicleLicenseResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "RecognizeVehicleLicenseResponse struct{}"
	}

	return strings.Join([]string{"RecognizeVehicleLicenseResponse", string(data)}, " ")
}
