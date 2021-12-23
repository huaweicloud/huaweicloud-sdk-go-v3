package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type RecognizeDriverLicenseRequest struct {
	Body *DriverLicenseRequestBody `json:"body,omitempty"`
}

func (o RecognizeDriverLicenseRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecognizeDriverLicenseRequest struct{}"
	}

	return strings.Join([]string{"RecognizeDriverLicenseRequest", string(data)}, " ")
}
