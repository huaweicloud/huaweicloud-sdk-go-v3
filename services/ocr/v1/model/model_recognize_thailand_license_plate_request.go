package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type RecognizeThailandLicensePlateRequest struct {
	Body *ThailandLicensePlateRequestBody `json:"body,omitempty" xml:"body"`
}

func (o RecognizeThailandLicensePlateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecognizeThailandLicensePlateRequest struct{}"
	}

	return strings.Join([]string{"RecognizeThailandLicensePlateRequest", string(data)}, " ")
}
