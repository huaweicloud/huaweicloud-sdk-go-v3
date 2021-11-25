package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type RunAspectSentimentRequest struct {
	Body *AspectSentimentRequest `json:"body,omitempty"`
}

func (o RunAspectSentimentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunAspectSentimentRequest struct{}"
	}

	return strings.Join([]string{"RunAspectSentimentRequest", string(data)}, " ")
}
