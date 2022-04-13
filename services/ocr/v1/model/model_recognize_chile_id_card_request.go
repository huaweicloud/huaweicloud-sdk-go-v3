package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type RecognizeChileIdCardRequest struct {
	Body *ChileIdCardRequestBody `json:"body,omitempty"`
}

func (o RecognizeChileIdCardRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecognizeChileIdCardRequest struct{}"
	}

	return strings.Join([]string{"RecognizeChileIdCardRequest", string(data)}, " ")
}
