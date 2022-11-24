package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type RunEntitySentimentRequest struct {
	Body *EntitySentimentReq `json:"body,omitempty"`
}

func (o RunEntitySentimentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunEntitySentimentRequest struct{}"
	}

	return strings.Join([]string{"RunEntitySentimentRequest", string(data)}, " ")
}
