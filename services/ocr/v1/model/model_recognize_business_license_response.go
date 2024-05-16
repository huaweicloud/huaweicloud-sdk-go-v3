package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RecognizeBusinessLicenseResponse Response Object
type RecognizeBusinessLicenseResponse struct {
	Result *BusinessLicenseResult `json:"result,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RecognizeBusinessLicenseResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecognizeBusinessLicenseResponse struct{}"
	}

	return strings.Join([]string{"RecognizeBusinessLicenseResponse", string(data)}, " ")
}
