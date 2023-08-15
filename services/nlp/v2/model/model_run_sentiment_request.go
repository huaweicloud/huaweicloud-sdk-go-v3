package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RunSentimentRequest Request Object
type RunSentimentRequest struct {
	Body *HwCloudSentimentReq `json:"body,omitempty"`
}

func (o RunSentimentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunSentimentRequest struct{}"
	}

	return strings.Join([]string{"RunSentimentRequest", string(data)}, " ")
}
