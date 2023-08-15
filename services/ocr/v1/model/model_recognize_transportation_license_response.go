package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RecognizeTransportationLicenseResponse Response Object
type RecognizeTransportationLicenseResponse struct {
	Result         *TransportationLicenseResult `json:"result,omitempty"`
	HttpStatusCode int                          `json:"-"`
}

func (o RecognizeTransportationLicenseResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecognizeTransportationLicenseResponse struct{}"
	}

	return strings.Join([]string{"RecognizeTransportationLicenseResponse", string(data)}, " ")
}
