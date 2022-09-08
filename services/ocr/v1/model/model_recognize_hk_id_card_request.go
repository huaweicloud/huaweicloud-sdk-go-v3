package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type RecognizeHkIdCardRequest struct {
	Body *HkIdCardRequestBody `json:"body,omitempty"`
}

func (o RecognizeHkIdCardRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecognizeHkIdCardRequest struct{}"
	}

	return strings.Join([]string{"RecognizeHkIdCardRequest", string(data)}, " ")
}
