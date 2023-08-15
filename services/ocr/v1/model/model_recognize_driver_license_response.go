package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RecognizeDriverLicenseResponse Response Object
type RecognizeDriverLicenseResponse struct {
	Result         *DriverLicenseResult `json:"result,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o RecognizeDriverLicenseResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecognizeDriverLicenseResponse struct{}"
	}

	return strings.Join([]string{"RecognizeDriverLicenseResponse", string(data)}, " ")
}
