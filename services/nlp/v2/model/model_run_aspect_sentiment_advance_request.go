package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type RunAspectSentimentAdvanceRequest struct {
	Body *AspectSentimentAdvanceRequest `json:"body,omitempty"`
}

func (o RunAspectSentimentAdvanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunAspectSentimentAdvanceRequest struct{}"
	}

	return strings.Join([]string{"RunAspectSentimentAdvanceRequest", string(data)}, " ")
}
