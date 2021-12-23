package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type RecognizeBusinessLicenseRequest struct {
	Body *BusinessLicenseRequestBody `json:"body,omitempty"`
}

func (o RecognizeBusinessLicenseRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecognizeBusinessLicenseRequest struct{}"
	}

	return strings.Join([]string{"RecognizeBusinessLicenseRequest", string(data)}, " ")
}
