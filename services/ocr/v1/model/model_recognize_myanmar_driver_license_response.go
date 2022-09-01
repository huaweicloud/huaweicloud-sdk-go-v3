package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type RecognizeMyanmarDriverLicenseResponse struct {
	Result         *MyanmarDriverLicenseResult `json:"result,omitempty" xml:"result"`
	HttpStatusCode int                         `json:"-"`
}

func (o RecognizeMyanmarDriverLicenseResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecognizeMyanmarDriverLicenseResponse struct{}"
	}

	return strings.Join([]string{"RecognizeMyanmarDriverLicenseResponse", string(data)}, " ")
}
