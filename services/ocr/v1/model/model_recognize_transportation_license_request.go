package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type RecognizeTransportationLicenseRequest struct {
	Body *TransportationLicenseRequestBody `json:"body,omitempty" xml:"body"`
}

func (o RecognizeTransportationLicenseRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecognizeTransportationLicenseRequest struct{}"
	}

	return strings.Join([]string{"RecognizeTransportationLicenseRequest", string(data)}, " ")
}
