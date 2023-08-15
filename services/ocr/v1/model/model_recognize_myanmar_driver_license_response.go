package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RecognizeMyanmarDriverLicenseResponse Response Object
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
