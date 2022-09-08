package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type RecognizeMyanmarDriverLicenseRequest struct {
	Body *MyanmarDriverLicenseRequestBody `json:"body,omitempty"`
}

func (o RecognizeMyanmarDriverLicenseRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecognizeMyanmarDriverLicenseRequest struct{}"
	}

	return strings.Join([]string{"RecognizeMyanmarDriverLicenseRequest", string(data)}, " ")
}
