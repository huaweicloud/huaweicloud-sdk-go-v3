package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type RecognizeVehicleLicenseRequest struct {
	Body *VehicleLicenseRequestBody `json:"body,omitempty"`
}

func (o RecognizeVehicleLicenseRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecognizeVehicleLicenseRequest struct{}"
	}

	return strings.Join([]string{"RecognizeVehicleLicenseRequest", string(data)}, " ")
}
