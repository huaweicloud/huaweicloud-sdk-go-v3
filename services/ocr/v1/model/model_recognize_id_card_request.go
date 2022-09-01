package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type RecognizeIdCardRequest struct {
	Body *IdCardRequestBody `json:"body,omitempty" xml:"body"`
}

func (o RecognizeIdCardRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecognizeIdCardRequest struct{}"
	}

	return strings.Join([]string{"RecognizeIdCardRequest", string(data)}, " ")
}
