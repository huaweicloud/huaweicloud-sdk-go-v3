package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type RecognizeLicensePlateRequest struct {
	Body *LicensePlateRequestBody `json:"body,omitempty" xml:"body"`
}

func (o RecognizeLicensePlateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecognizeLicensePlateRequest struct{}"
	}

	return strings.Join([]string{"RecognizeLicensePlateRequest", string(data)}, " ")
}
