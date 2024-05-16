package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RecognizeVehicleCertificateResponse Response Object
type RecognizeVehicleCertificateResponse struct {
	Result *VehicleCertificateResult `json:"result,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RecognizeVehicleCertificateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecognizeVehicleCertificateResponse struct{}"
	}

	return strings.Join([]string{"RecognizeVehicleCertificateResponse", string(data)}, " ")
}
