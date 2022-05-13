package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type RecognizeHealthCodeRequest struct {
	Body *HealthCodeRequestBody `json:"body,omitempty"`
}

func (o RecognizeHealthCodeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecognizeHealthCodeRequest struct{}"
	}

	return strings.Join([]string{"RecognizeHealthCodeRequest", string(data)}, " ")
}
