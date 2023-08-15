package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RecognizeVehicleLicenseResponse Response Object
type RecognizeVehicleLicenseResponse struct {
	Result         *VehicleLicenseResult `json:"result,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o RecognizeVehicleLicenseResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecognizeVehicleLicenseResponse struct{}"
	}

	return strings.Join([]string{"RecognizeVehicleLicenseResponse", string(data)}, " ")
}
