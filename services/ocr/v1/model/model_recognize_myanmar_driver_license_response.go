package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type RecognizeMyanmarDriverLicenseResponse struct {
	Result         *MyanmarDriverLicenseResult `json:"result,omitempty"`
	HttpStatusCode int                         `json:"-"`
}

func (o RecognizeMyanmarDriverLicenseResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecognizeMyanmarDriverLicenseResponse struct{}"
	}

	return strings.Join([]string{"RecognizeMyanmarDriverLicenseResponse", string(data)}, " ")
}
