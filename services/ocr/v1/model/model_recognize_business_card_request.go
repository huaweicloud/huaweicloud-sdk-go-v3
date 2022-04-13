package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type RecognizeBusinessCardRequest struct {
	Body *BusinessCardRequestBody `json:"body,omitempty"`
}

func (o RecognizeBusinessCardRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecognizeBusinessCardRequest struct{}"
	}

	return strings.Join([]string{"RecognizeBusinessCardRequest", string(data)}, " ")
}
